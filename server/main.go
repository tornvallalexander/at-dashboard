package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tornvallalexander/at-dashboard/server/db"
	"net/http"
	"os"
	"time"
)

var (
	listenAddr = ":" + os.Getenv("PORT")
)

func main() {
	rdb, err := db.ConnectDB()
	CheckFatalError(err)
	router := initRouter(rdb)
	err = router.Run(listenAddr)
	CheckFatalError(err)
}

func initRouter(rdb *db.Database) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/top-words", func(c *gin.Context) {
		words, err := rdb.GetAllTopWords()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"words": words,
		})
	})

	r.GET("/random-top-words", func(c *gin.Context) {
		randWords, err := rdb.GetRandomTopWords()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"words": randWords,
		})
	})

	return r
}
