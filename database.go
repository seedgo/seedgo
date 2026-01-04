package seedgo

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBMap = make(map[string]*gorm.DB)

var DatabaseDriverNotSupportErr error

func DB(dbNames ...string) *gorm.DB {
	dbname := "default"
	if len(dbNames) > 0 {
		dbname = dbNames[0]
	}

	db, ok := DBMap[dbname]
	if !ok {
		Logger.Errorf("datasource %s not exist", dbname)
		return nil
	}

	return db
}

func ParseDatabaseConfig() {
	datasources := viper.GetStringMap("datasource")
	if len(datasources) == 0 {
		return
	}

	for dsName := range datasources {
		dialect, err := parseDatasource(dsName)
		if err != nil {
			Logger.Errorf("datasource: [%s] parse err: %s", dsName, err.Error())
			continue
		}

		db, err := gorm.Open(dialect, &gorm.Config{})
		if err != nil {
			Logger.Errorf("database: [%s] open err: %s", dsName, err.Error())
			continue
		}

		if ServerConfig.Debug {
			db = db.Debug()
		}

		DBMap[dsName] = db
	}
}

func parseDatasource(dsName string) (gorm.Dialector, error) {
	driver := viper.GetString("datasource." + dsName + ".driver")
	if driver == "postgres" {
		return parsePostgres(dsName), nil
	} else if driver == "mysql" {
		return parseMysql(dsName), nil
	}

	return nil,
		fmt.Errorf("%w", DatabaseDriverNotSupportErr)
}

func parseMysql(dsName string) gorm.Dialector {
	confKey := "datasource." + dsName
	conf := viper.Sub(confKey)
	host := conf.GetString("host")
	port := conf.GetString("port")
	database := conf.GetString("database")
	username := conf.GetString("username")
	password := conf.GetString("password")
	charset := conf.GetString("charset")
	loc := conf.GetString("loc")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	return mysql.Open(dsn)
}

func parsePostgres(dsName string) gorm.Dialector {
	confKey := "datasource." + dsName
	conf := viper.Sub(confKey)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		conf.GetString("host"),
		conf.GetString("username"),
		conf.GetString("password"),
		conf.GetString("database"),
		conf.GetInt("port"),
		conf.GetString("sslmode"),
		conf.GetString("timezone"),
	)

	return postgres.Open(dsn)
}
