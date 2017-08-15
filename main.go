package main

import (
	_ "github.com/TheBeege/mentor-me-api/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"fmt"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func init() {
	dbHost := os.Getenv("dbhost")
	dbUser := os.Getenv("dbuser")
	dbPass := os.Getenv("dbpass")
	dbDatabase := os.Getenv("dbdatabase")
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbDatabase)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", conString)
	orm.SetMaxIdleConns("default", 30)
	// Set to UTC time
	orm.DefaultTimeLoc = time.UTC
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
