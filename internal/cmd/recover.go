package cmd

import (
	"thor/internal/routine"
	"thor/internal/services"
)

func recoverTorrents(svcs *services.Services) error {
	hashes, err := svcs.Database.ListNonDeletedTracks()
	if err != nil {
		return err
	}

	for _, hash := range hashes {
		torrent, err := svcs.Download.AddDownload(hash)
		if err != nil {
			return err
		}

		go routine.HandleTorrent(svcs, torrent)
	}

	return nil
}
