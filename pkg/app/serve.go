package app

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	Port = ":8090"
)

func (a *App) Serve() {
	logrus.Infof("started server")
	http.HandleFunc("/scans", a.scans)
	err := http.ListenAndServe(Port, nil)
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
