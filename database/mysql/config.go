package mysql

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"online-ticketing/app/config"
	"time"
)

type DbConnection struct {
	Db *gorm.DB
}

func NewConnection() *DbConnection {
	var i int

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Db.User, config.Db.Pass, config.Db.Host, config.Db.Port, config.Db.DatabaseName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"db_name":    config.Db.DatabaseName,
			"driver":     config.Db.DriverName,
			"connection": config.Db.ConnectionString,
			"error":      err.Error(),
		}).Warningln("Failed to connect database. Trying Reconnecting...")

		for {
			i++
			db, err := gorm.Open(mysql.Open(config.Db.ConnectionString), &gorm.Config{})
			if err != nil {
				logrus.Warningf("Reconnect(%d) %v : %v...\n", i, err, config.Db.ConnectionString)
				time.Sleep(3 * time.Second)
				if i == 20 {
					logrus.WithFields(logrus.Fields{
						"db_name":    config.Db.DatabaseName,
						"driver":     config.Db.DriverName,
						"connection": config.Db.ConnectionString,
						"error":      err.Error(),
					}).Fatalln("Database Connection Failed!")
				} else {
					continue
				}
			}

			return &DbConnection{db}
		}
	} else {
		logrus.Infoln("Connected to Database Successfully!")
	}

	if config.App.AppEnv != "production" {
		db.Debug()
	}

	// Mendapatkan koneksi SQL dari GORM
	sqlDB, err := db.DB()
	if err != nil {
		// Handle error
	}

	sqlDB.SetMaxIdleConns(config.Db.MaxConnectionIdle)
	sqlDB.SetMaxOpenConns(config.Db.MaxConnectionOpen)

	if config.App.AppDebug {
		db = db.Debug()
	}

	return &DbConnection{db}
}
