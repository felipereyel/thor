package database

type IDatabase interface {
	Close() error
	ListTracks() ([]string, error)
	UpsertTrack(hash string, status string) error
}
