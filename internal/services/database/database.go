package database

import (
	"database/sql"

	"github.com/felipereyel/thor/internal/config"
	_ "modernc.org/sqlite"
)

type database struct {
	conn *sql.DB
}

func Factory(cfg *config.ServerConfigs) (*database, error) {
	conn, err := sql.Open("sqlite", cfg.DatabaseFile)
	if err != nil {
		return nil, err
	}

	return &database{conn}, nil
}

func (db *database) Close() error {
	return db.conn.Close()
}

func (db *database) CreateTrack(hash string) error {
	query := `INSERT INTO tracks (hash, status) VALUES (?, 'created')`
	_, err := db.conn.Exec(query, hash)
	return err
}

func (db *database) ListTracks() ([]string, error) {
	query := `SELECT hash FROM tracks`
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var hashes []string
	for rows.Next() {
		var hash string
		err := rows.Scan(&hash)
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, hash)
	}

	return hashes, nil
}

func (db *database) UpdateTrackStatus(hash string, status string) error {
	query := `UPDATE tracks SET status = ? WHERE hash = ?`
	_, err := db.conn.Exec(query, status, hash)
	return err
}
