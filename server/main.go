package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tornvallalexander/at-dashboard/server/db"
	"net/http"
)

var (
	listenAddr = "localhost:8000"
)

func main() {
	rdb, err := db.ConnectDB()
	CheckErrorFatal(err)

	router := initRouter(rdb)
	err = router.Run(listenAddr)
	CheckErrorFatal(err)
}

func initRouter(rdb *db.Database) *gin.Engine {
	r := gin.Default()

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
