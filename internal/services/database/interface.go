package database

type IDatabase interface {
	Close() error
	ListNonDeletedTracks() ([]string, error)
	UpsertTrack(hash string, status string) error
}
