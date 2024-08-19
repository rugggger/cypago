package scans_resources

import (
	"database/sql"
	"time"
)

type ScansResources struct {
	sql *sql.DB
}

type Scan struct {
	ID    string    `json:"id"`
	Start time.Time `json:"start"`
}

type Resource struct {
	rType string
	name  string
}

func New(sql *sql.DB) *ScansResources {
	return &ScansResources{
		sql: sql,
	}
}

func (sr *ScansResources) GetScans() ([]*Scan, error) {

	var scans []*Scan

	query := `
SELECT  s.id, s.start
FROM scans s
`

	rows, err := sr.sql.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ID, &scan.Start); err != nil {
			return nil, err
		}
		scans = append(scans, &scan)
	}

	return scans, nil
}
