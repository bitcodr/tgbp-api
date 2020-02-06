package config

import (
	"database/sql"
	"github.com/spf13/viper"
	"time"
	"log"
	"os"
)

type AppInterface interface {
	Environment()
	DB() (*sql.DB, error)
}

type App struct {
	BotUsername, DBName, DBUserName, DBPass, CurrentTime string
}

var (
	AppConfig  *viper.Viper
	LangConfig *viper.Viper
)

func (app *App) Environment() {
	app.appConfig()
	app.langConfig()
	app.setAppConfig()
}

func (app *App) appConfig() {
	AppConfig = viper.New()
	AppConfig.SetConfigType("yaml")
	AppConfig.SetConfigName("config")
	AppConfig.AddConfigPath(os.Getenv("BOT_API_ROOT_DIR"))
	err := AppConfig.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func (app *App) langConfig() {
	LangConfig = viper.New()
	LangConfig.SetConfigType("yaml")
	LangConfig.SetConfigName("lang")
	LangConfig.AddConfigPath(os.Getenv("BOT_API_ROOT_DIR") + "/resource/lang")
	err := LangConfig.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func (app *App) setAppConfig() {
	app.DBName = AppConfig.GetString("DATABASES.MYSQL.DATABASE")
	app.DBUserName = AppConfig.GetString("DATABASES.MYSQL.USERNAME")
	app.DBPass = AppConfig.GetString("DATABASES.MYSQL.PASSWORD")
	app.CurrentTime = time.Now().UTC().Format("2006-01-02 03:04:05")
}
