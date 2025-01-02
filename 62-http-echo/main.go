package main

import (
	"errors"
	"flag"
	"fmt"
	"http-echo-demo/db"
	"http-echo-demo/handlers"
	"http-echo-demo/messages"
	"log"
	"os"
	"runtime"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
	_ "net/http/pprof"
)

var port string
var dsn string
var kafka_servers string
var secretKey = "somesecretkey"

func init() {
	port = os.Getenv("PORT")
	dsn = os.Getenv("DB_CON")
	if port == "" {
		flag.StringVar(&port, "port", "58089", "--port=58090 | -port=58089")
	}

	if dsn == "" {
		flag.StringVar(&dsn, "db", `host=localhost user=admin password=admin123 dbname=usersdb port=45432 sslmode=disable TimeZone=Asia/Shanghai`, `--db=host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai`)
	}

	if kafka_servers == "" {
		kafka_servers = "localhost:9092"
	}

	flag.Parse()
	log.SetFlags(log.Ldate | log.Lshortfile)
	//log.SetFlags()
	//f, err := os.OpenFile("mylogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//log.SetOutput(os.Stdout)
}

func main() {

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	dbConn, err := db.GetConnection(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	_ = dbConn

	userMessage := messages.NewUserMessage(kafka_servers, "user_create")
	go userMessage.Init()
	userHandler := handlers.NewUser(db.NewUserDB(dbConn), userMessage)
	//userHandler := handlers.NewUser(filedb.NewUserFileDB("users.db"))
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	public := e.Group("/public")

	public.GET("/", handlers.Root)
	public.GET("/ping", handlers.Ping)
	public.GET("/health", handlers.Health)
	public.POST("/login", userHandler.Login(secretKey))
	public.POST("/users", userHandler.Create)

	private := e.Group("/private")
	private.Use(JWTMiddleare)

	private.GET("/users/:id", userHandler.GetByID)

	if err := e.Start(":" + port); err != nil {
		log.Println(err.Error())
		runtime.Goexit()
	}
}

func JWTMiddleare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization: Bearer")
		if tokenString == "" {
			c.JSON(401, map[string]any{"error": "Authorization token required"})
			log.Println("error:" + "Authorization token required")
			return errors.New("some error")
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			fmt.Println("Token error", err.Error())
			c.JSON(401, map[string]any{"error": "some error in the token"})
			log.Println("error:" + err.Error())
			return err
		}
		if !token.Valid {
			c.JSON(401, "token is not valid")
			log.Println("token is not valid")
			return errors.New("token is not valid")
		}
		//claims, err := token.Claims.(jwt.MapClaims)
		if token.Valid {
			return next(c)
		}
		return nil
	}
}
