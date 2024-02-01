package database

type IDatabase interface {
	Close() error
	CreateTrack(hash string) error
	ListTracks() ([]string, error)
	UpdateTrackStatus(hash string, status string) error
}
