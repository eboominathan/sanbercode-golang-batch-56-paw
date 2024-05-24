package main

import (
	"database/sql"
	"fmt"
	"os"
	"practice/controllers"
	"practice/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {

	err = godotenv.Load("config/.env")

	host := os.Getenv("DB_HOST")
	strPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, _ := strconv.Atoi(strPort)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {

		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(db)

	defer db.Close()

	router := gin.Default()
	router.GET("/persons", controllers.GetAllPerson)
	router.POST("/persons", controllers.InsertPerson)
	router.PUT("/persons/:id", controllers.UpdatePerson)
	router.DELETE("/persons/:id", controllers.DeletePerson)

	  var appPort = envPortOr("3000")

  	log.Fatal(http.ListenAndServe(appPort, handler))
	
	// router.Run("0.0.0.0:" + PORT)
}

func envPortOr(port string) string {
  // If `PORT` variable in environment exists, return it
  if envPort := os.Getenv("PORT"); envPort != "" {
    return ":" + envPort
  }
  // Otherwise, return the value of `port` variable from function argument
  return ":" + port
}
