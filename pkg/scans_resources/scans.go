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
	ScanID string
	rType  string
	name   string
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

func (sr *ScansResources) CreateScan() (*Scan, error) {
	var scan *Scan
	query := `
INSERT INTO scans
DEFAULT VALUES
RETURNING id;
`
	rows, err := sr.sql.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ID); err != nil {
			return nil, err
		}
	}
	return scan, nil
}

func (sr *ScansResources) CreateResource(r Resource) (string, error) {
	return "sdf", nil

}
