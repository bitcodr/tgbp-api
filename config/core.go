package config

import (
	"database/sql"
	"time"
)

type AppInterface interface {
	Environment()
	DB() (*sql.DB, error)
	SetAppConfig() *App
}

type App struct {
	ProjectDir, BotUsername, DBName, DBUserName, DBPass, CurrentTime string
}

func (app *App) SetAppConfig() *App {
	app.DBName = AppConfig.GetString("DATABASES.MYSQL.DATABASE")
	app.DBUserName = AppConfig.GetString("DATABASES.MYSQL.USERNAME")
	app.DBPass = AppConfig.GetString("DATABASES.MYSQL.PASSWORD")
	app.CurrentTime = time.Now().UTC().Format("2006-01-02 03:04:05")
	return app
}
