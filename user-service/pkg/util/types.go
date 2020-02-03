package util

type Config struct {
	Title string   `toml:"title "`
	Server Server
	Mysql    mysql
	Log      *Log
}

type Server struct {
	GrpcPort string   `toml:"grpcPort"`
	RestPort string   `toml:"restPort"`
}

type mysql struct {
	MysqlIp             string `toml:"mysqlIp"`
	MysqlPort           int    `toml:"mysqlPort"`
	MysqlUser           string `toml:"mysqlUsr"`
	MysqlPwd            string `toml:"mysqlPwd"`
	MysqlEncPwd            string `toml:"mysqlEncPwd"`
	DbName              string `toml:"dbName"`
}

type Log struct {
	LogLevel        string `toml:"loglevel"`
	LogFile  string `toml:"logFile"`
	LogTimeFormat  string   `toml:"logTimeFormat"`
}
