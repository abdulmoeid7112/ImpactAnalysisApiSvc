package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	JWTSecretKey = "jwt.secret.key"

	ServerHost = "server.host"
	ServerPort = "server.port"

	DatasetApiURL = "dataset.api.url"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	_ = viper.BindEnv(JWTSecretKey, "JWT_SECRET_KEY")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	_ = viper.BindEnv(DatasetApiURL, "DATASET_API_URL")

	// defaults
	viper.SetDefault(DbName, "user")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	viper.SetDefault(DbUser, "admin")
	viper.SetDefault(DbPass, "TestPass1234")

	viper.SetDefault(JWTSecretKey,  []byte("323123123123123"))

	viper.SetDefault(ServerHost, "127.0.0.1")
	viper.SetDefault(ServerPort, "8080")

	viper.SetDefault(DatasetApiURL, "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv")
}