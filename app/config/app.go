package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type AppConfig struct {
	AppEnv   string
	AppName  string
	AppPort  string
	AppDebug bool
}

type Database struct {
	DatabaseName      string
	DriverName        string
	ConnectionString  string
	Host              string
	User              string
	Pass              string
	Port              string
	MaxConnectionOpen int
	MaxConnectionIdle int
}

var (
	App = AppConfig{}
	Db  = Database{}
)

func Init() {

	ex, _ := os.Executable()
	if os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == "" {
		exPath := filepath.Dir(ex)
		if err := godotenv.Load(exPath + "/.env"); err != nil {
			logrus.WithFields(logrus.Fields{
				"executable":  ex,
				"filepath":    exPath,
				"environment": os.Getenv("APP_ENV"),
				"error":       err.Error(),
			}).Fatalln(".env is not loaded properly")
			os.Exit(1)
		}
	}

	// Set Viper to use environment variables
	viper.SetEnvPrefix("APP") // Prefix for environment variables (optional)
	viper.AutomaticEnv()

	App.AppEnv = viper.GetString("ENV")
	App.AppName = viper.GetString("NAME")
	App.AppPort = viper.GetString("PORT")
	App.AppDebug = viper.GetBool("DEBUG")

	// Set Viper to use environment variables with the "DB" prefix
	viper.SetEnvPrefix("DB")
	viper.AutomaticEnv()

	Db.DatabaseName = viper.GetString("NAME")
	Db.DriverName = viper.GetString("DRIVER_NAME")
	//Db.ConnectionString = viper.GetString("DB_CONNECTION_STRING")
	Db.Host = viper.GetString("HOST")
	Db.User = viper.GetString("USER")
	Db.Pass = viper.GetString("PASS")
	Db.Port = viper.GetString("PORT")
	Db.MaxConnectionOpen = viper.GetInt("MAX_CONNECTION_OPEN")
	Db.MaxConnectionIdle = viper.GetInt("MAX_CONNECTION_IDLE")
}
