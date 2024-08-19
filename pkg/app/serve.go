package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	Port = ":8090"
)

func (a *App) Serve() {
	logrus.Infof("started server")
	r := gin.Default()

	r.GET("/scans", func(c *gin.Context) {
		// Handle GET request
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
		// Handle GET request
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

	err := r.Run(Port)

	if err != nil {
		logrus.Fatalf("couldn't start server: %v", err)
	}

}

func (a *App) sendJson(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		logrus.Errorf("error marshalling %v", o)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		logrus.Errorf("error writing response")
	}
}

func (a *App) scans(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var res interface{}
	var err error
	switch req.Method {
	case http.MethodGet:
		res, err = a.sr.GetScans()
	case http.MethodPost:
		res, err = a.sr.CreateScan()

	default:
		// Method not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if err != nil {
		logrus.Errorf("error %v", err)
	}
	a.sendJson(w, res)
}

func (a *App) resources(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var res interface{}
	var err error
	switch req.Method {
	case http.MethodGet:
		res, err = a.sr.GetScans()
	case http.MethodPost:
		res, err = a.sr.CreateScan()

	default:
		// Method not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if err != nil {
		logrus.Errorf("error %v", err)
	}
	a.sendJson(w, res)
}
