package main

import (
	"flag"
	"http-echo-demo/db"
	"http-echo-demo/filedb"
	"http-echo-demo/handlers"
	"log"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string
var dsn string

func init() {
	port = os.Getenv("PORT")
	dsn = os.Getenv("DB_CON")
	if port == "" {
		flag.StringVar(&port, "port", "58089", "--port=58090 | -port=58089")
	}

	if dsn == "" {
		flag.StringVar(&dsn, "db", `host=localhost user=admin password=admin123 dbname=usersdb port=45432 sslmode=disable TimeZone=Asia/Shanghai`, `--db=host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai`)
	}

	flag.Parse()
	log.SetFlags(log.Ldate | log.Lshortfile)
	//log.SetFlags()
	//f, err := os.OpenFile("mylogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//log.SetOutput(os.Stdout)
}

func main() {

	dbConn, err := db.GetConnection(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = dbConn

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", handlers.Root)
	e.GET("/ping", handlers.Ping)
	e.GET("/health", handlers.Health)

	//userHandler := handlers.NewUser(db.NewUserDB(dbConn))
	userHandler := handlers.NewUser(filedb.NewUserFileDB("users.db"))

	e.POST("/user", userHandler.Create)

	if err := e.Start(":" + port); err != nil {
		log.Println(err.Error())
		runtime.Goexit()
	}
}
