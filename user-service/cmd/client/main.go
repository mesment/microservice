package main

import (
	"context"
	"flag"
	"github.com/mesment/microservice/user-service/pkg/logger"
	"time"
	"google.golang.org/grpc"
	"github.com/mesment/microservice/user-service/pkg/api/users/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// get configuration
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()
	logger.Init("info","")
	log:=logger.NewSugardLogger()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Infof("did not connect: %v", err)
	}
	defer conn.Close()

	client := v1.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Signup
	req1 := v1.SignupRequest{
		UserName:"mesment",
		Password:"12345678",
		Email: "vanley2012@163.com",
	}
	res1, err := client.Signup(ctx, &req1)
	if err != nil {
		log.Errorf("Signup error: %v", err)
	}
	log.Infof("Signup result: <%+v>\n", res1)

	loginRequest := v1.LoginRequest{
		Login:"mesment",
		Password:"123456789",
	}
	loginResponse, err := client.Login(ctx, &loginRequest)
	if err != nil {
		log.Errorf("Login error: %v", err)
	}
	log.Infof("Login result: <%+v>\n", loginResponse)

	// authuser
	authUserRequest := v1.AuthUserRequest{
		Token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklEIjoiOWM5YjYyYTQtYjk2Ny00MDBlLThjMjUtNDMxNDZiMDRmNTViIiwiVXNlck5hbWUiOiJtZXNtZW50IiwiUGFzc3dvcmQiOiIkMmEkMTAkUDZyQ1VSUGxRei9tbVdXSVJBTlh2LktxbWxDMHdOU05hOTRPWXFrYlcwbm83bk5qRy9aVkMiLCJFbWFpbCI6InZhbmxleTIwMTJAMTYzLmNvbSIsImNyZWF0ZV9hdCI6IjIwMjAtMDItMDJUMDc6MDE6NTQrMDg6MDAiLCJ1cGRhdGVfYXQiOiIyMDIwLTAyLTAyVDA3OjAxOjU0KzA4OjAwIiwiRGVsZXRlZEF0IjpudWxsfSwiZXhwIjoxNTgwODU4ODE1LCJpc3MiOiJ1c2VyLXNlcnZpY2UudXNlciJ9.sXUnujm2ndgoblKobFqYzaoV57R-mQOw0dIxI2-SCCY",
	}
	authUserResp, err := client.AuthUser(ctx, &authUserRequest)
	if err != nil {
		log.Fatalf("AuthUser error: %v", err)
	}
	log.Infof("AuthUser result: <%+v>\n", authUserResp)

}
