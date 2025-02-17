package routine

import (
	"fmt"

	"thor/internal/services"

	"github.com/anacrolix/torrent"
)

func HandleTorrent(svcs *services.Services, torrent *torrent.Torrent) {
	hash := torrent.InfoHash().HexString()
	upsert := func(status string) {
		fmt.Printf("[HandleTorrent][%s] %s\n", hash, status)
		if err := svcs.Database.UpsertTrack(hash, status); err != nil {
			fmt.Printf("[HandleTorrent][%s] Failed to upsert %s: %s\n", hash, status, err.Error())
		}
	}

	upsert("loading")
	<-torrent.GotInfo()

	upsert("downloading")
	torrent.DownloadAll()

	<-torrent.Complete.On()
	upsert("downloaded")
}
