package scans_resources

import "database/sql"

type ScansResources struct {
	sql *sql.DB
}

type Scan struct {
	ID string `json:"id"`
}

func New(sql *sql.DB) *ScansResources {
	return &ScansResources{
		sql: sql,
	}
}

func (sr *ScansResources) GetScans() ([]*Scan, error) {

	var scans []*Scan

	// Define the query to get scans
	query := "SELECT id, name FROM scans"

	// Execute the query
	rows, err := sr.sql.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set
	for rows.Next() {
		var scan Scan
		if err := rows.Scan(&scan.ID); err != nil {
			return nil, err
		}
		scans = append(scans, &scan)
	}

	scans = []*Scan{
		{ID: "1"},
		{ID: "2"},
		{ID: "3"},
	}
	return scans, nil
}
