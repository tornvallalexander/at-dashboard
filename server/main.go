package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tornvallalexander/at-dashboard/server/db"
	"net/http"
	"time"
)

var (
	listenAddr = "localhost:8000"
)

func main() {
	rdb, err := db.ConnectDB()
	CheckFatalError(err)
	router := initRouter(rdb)
	err = router.Run(listenAddr)
	CheckFatalError(err)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func initRouter(rdb *db.Database) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/topWords", func(c *gin.Context) {
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

	r.GET("/randomTopWords", func(c *gin.Context) {
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
