package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
)

const (
	Port = ":8090"
)

func (a *App) Serve() {
	logrus.Infof("started server")
	r := gin.Default()

	r.GET("/scans", func(c *gin.Context) {
		res, err := a.sr.GetScans()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"scans": res,
		})
	})

	r.POST("/scans", func(c *gin.Context) {
		res, err := a.sr.CreateScan()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"scans": res,
		})
	})

	r.POST("/scans/:scan_id/resources", func(c *gin.Context) {
		idStr := c.Param("scan_id")
		id, err := strconv.Atoi(idStr)

		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		name, ok := data["name"].(string)
		if !ok {
			log.Fatalf("Name is not a string")
		}
		resourceType, ok := data["resourceType"].(string)
		if !ok {
			log.Fatalf("resourceType is not a string")
		}

		res, err := a.sr.CreateResource(name, resourceType, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": res,
			"id":      id,
			"body":    data,
		})
	})

	err := r.Run(Port)

	if err != nil {
		logrus.Fatalf("couldn't start server: %v", err)
	}

}
