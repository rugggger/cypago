package app

import (
	"database/sql"
	"github.com/cypago/pkg/db"
	"github.com/cypago/pkg/logger"
	"github.com/cypago/pkg/scans_resources"
	_ "github.com/lib/pq" // Import the pq package
)

type App struct {
	sqlConnection *sql.DB
	sr            *scans_resources.ScansResources
}

func New() *App {
	logger.InitLogger()
	conn := db.GetDBConnection()
	sr := scans_resources.New()

	return &App{
		sqlConnection: conn,
		sr:            sr,
	}
}
