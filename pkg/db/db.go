package db

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bitbucket.org/liamstask/goose/lib/goose"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *gorm.DB

var initOnce sync.Once

func Init() {
	initOnce.Do(func() {
		dbUserName := viper.GetString("database.username")
		dbPassword := viper.GetString("database.password")
		dbUrl := viper.GetString("database.url")
		dbName := viper.GetString("database.name")
		dbSchema := viper.GetString("database.schema")
		dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable search_path=%s password=%s", dbUrl, dbUserName, dbName, dbSchema, dbPassword)
		maxIdleConnection := viper.GetInt("postgress.maxIdleConnections")
		maxOpenConnection := viper.GetInt("postgres.maxOpenConnection")
		connectionMaxLifeTime := viper.GetInt("postgres.connMaxxLifetimeInHours")

		db, err := gorm.Open("postgres", dbURI)
		if err != nil {
			log.Error("Failed to connect to DB.", err)
			os.Exit(1)
		}

		db.DB().SetMaxIdleConns(maxIdleConnection)
		db.DB().SetMaxOpenConns(maxOpenConnection)
		db.DB().SetConnMaxLifetime(time.Hour * time.Duration(connectionMaxLifeTime))

		workingDir, err := os.Getwd()
		if err != nil {
			log.Error("Not able to fetch working directory")
			os.Exit(1)
		}
		migrationDir := workingDir + "/internal/migrations"

		migrateConf := &goose.DBConf{
			MigrationsDir: migrationDir,
			Driver: goose.DBDriver{
				Name:    "postgres",
				OpenStr: dbURI,
				Import:  "githhub.com/lib/pq",
				Dialect: &goose.PostgresDialect{},
			},
		}

		latest, _ := goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
		log.Info("Most recent db version ", latest)
		err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, db.DB())
		if err != nil {
			log.Error("Error in migration: ", err.Error())
			os.Exit(1)
		}
	})
}

func GetDB() *gorm.DB {
	return db
}
