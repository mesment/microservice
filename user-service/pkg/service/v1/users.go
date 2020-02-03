package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mesment/microservice/user-service/pkg/api/users/v1"
	"github.com/mesment/microservice/user-service/pkg/logger"
	"github.com/mesment/microservice/user-service/pkg/models"
	"github.com/mesment/microservice/user-service/pkg/util"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)
var log  = logger.NewSugardLogger()


type Repository interface {
	Login(ctx context.Context, in *v1.LoginRequest) (*v1.AuthResponse, error)
	Signup(ctx context.Context, in *v1.SignupRequest) (*v1.AuthResponse, error)
	UserNameUsed(ctx context.Context, in *v1.UserNameUsedRequest) (*v1.UserNameUsedResponse, error)
	EmailUsed(ctx context.Context, in *v1.EmailUsedRequest) (*v1.EmailUsedResponse, error)
	AuthUser(ctx context.Context, in *v1.AuthUserRequest) (*v1.AuthUserResponse, error)
}

type UserService struct{
	db *gorm.DB
	tokenService Authable
}


func NewUserService(cfg *util.Config) (*UserService, error) {

	db, err := util.GetDBConnection(cfg)
	if err != nil {
		log.Infof("connect db error:", err)
		return nil, err
	}
	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(models.User{})
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)
	svc := &UserService{db:db, tokenService:&TokenService{}}
	return svc, nil
}

func (u *UserService)Login(ctx context.Context, in *v1.LoginRequest) (*v1.AuthResponse, error)  {
	login, password := in.GetLogin(), in.GetPassword()

	user, err := u.ByNameOrEmail(login)
	if err != nil {
		log.Infof("login by name or email error:", err)
		return nil, err
	}
	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Infof("password error:", err)
		return &v1.AuthResponse{}, errors.New("password is wrong")
	}
	token, err := u.tokenService.Encode(user)
	if err != nil {
		return &v1.AuthResponse{}, err
	}
	return &v1.AuthResponse{Token:token}, nil
}

func (u *UserService)Signup(ctx context.Context, in *v1.SignupRequest) (*v1.AuthResponse, error) {
	userName, email, password := in.GetUserName(), in.GetEmail(), in.GetPassword()
	match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)
	if len(userName) < 4 || len(userName) > 20 || len(email) < 7 || len(email) > 35 || len(password) < 6 || len(password) > 128 || !match {
		return &v1.AuthResponse{}, errors.New("Validation failed")
	}
	res , err := u.UserNameUsed(ctx, &v1.UserNameUsedRequest{UserName: userName})
	if err != nil {
		log.Infof("Error returned from UserNameUsed: ", err.Error())
		return &v1.AuthResponse{}, errors.New("Something went wrong")
	}
	if res.GetUsed() {
		return &v1.AuthResponse{}, errors.New("UserName is already exist")
	}

	resp, err := u.EmailUsed(ctx, &v1.EmailUsedRequest{Email: email})
	if err != nil {
		log.Infof("Error returned from EmailUsed: ", err.Error())
		return &v1.AuthResponse{}, errors.New("Something went wrong")
	}
	if resp.GetUsed() {
		return &v1.AuthResponse{}, errors.New("Email is already exist")
	}

	// 对明文密码加密保存
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return &v1.AuthResponse{},errors.New(fmt.Sprintf("error hashing password: %v", err))
	}

	newUser := models.User{ID: uuid.NewV4().String(), UserName: userName, Email: email, Password: string(hashedPassword)}

	token, err := u.tokenService.Encode(&newUser)
	if err != nil {
		return &v1.AuthResponse{}, err
	}
	u.db.Create(&newUser)
	return &v1.AuthResponse{Token:token}, nil
}

func (u *UserService)UserNameUsed(ctx context.Context, in *v1.UserNameUsedRequest) (*v1.UserNameUsedResponse, error) {
	userName := in.GetUserName()
	result := &models.User{}

	u.db.Where(&models.User{UserName:userName}).First(result)
	return &v1.UserNameUsedResponse{Used:*result != models.NilUser}, nil
}

func (u *UserService)EmailUsed(ctx context.Context, in *v1.EmailUsedRequest) (*v1.EmailUsedResponse, error) {
	email := in.GetEmail()
	result := &models.User{}

	u.db.Where(&models.User{Email:email}).First(result)
	return &v1.EmailUsedResponse{Used:*result != models.NilUser}, nil
}


func (u *UserService)AuthUser(ctx context.Context, in *v1.AuthUserRequest) (*v1.AuthUserResponse, error) {
	token := in.GetToken()
	claim, err := u.tokenService.Decode(token)
	if err != nil {
		log.Infof("UserFromToken error:", err)
		return &v1.AuthUserResponse{}, err
	}
	user := claim.User
	return  &v1.AuthUserResponse{ID:user.ID, UserName:user.UserName, Email:user.Email},nil
}


func (u *UserService) ByName(name string) (*models.User, error) {
	user := &models.User{}

	err := u.db.Where("user_name = ? ", name).First(user).Error
	switch err {
	case nil :
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, err
	default:
		return nil, errors.New("unexpect error")
	}
}

func (u *UserService) ByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Where("email = ? ", email).First(user).Error
	switch err {
	case nil :
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, err
	default:
		return nil, errors.New("unexpect error")
	}
}


func (u *UserService) ByNameOrEmail(login string) (*models.User, error) {
	user := &models.User{}
	err := u.db.Where("email = ? ", login).Or("user_name = ?", login).First(user).Error
	switch err {
	case nil :
		return user, nil
	case gorm.ErrRecordNotFound:
		return nil, err
	default:
		return nil, errors.New("unexpect error")
	}
}
