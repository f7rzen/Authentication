package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	connStr := "user=postgres password=760411 dbname=users sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/html/*")
	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", nil)
	})
	r.GET("/register", RegisterPage)
	r.GET("/login", LoginPage)
	r.POST("/register", RegisterUser)
	r.POST("/login", LoginUser)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
