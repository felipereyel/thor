package download

import (
	"errors"
	"fmt"
	"thor/internal/config"
	"thor/internal/models"
	"thor/internal/utils"

	"github.com/anacrolix/torrent"
)

type downloadSvc struct {
	client *torrent.Client
}

func Factory(cfg *config.ServerConfigs) (*downloadSvc, error) {
	defaultConfig := torrent.NewDefaultClientConfig()
	defaultConfig.DataDir = cfg.DataDir
	client, err := torrent.NewClient(defaultConfig)

	if err != nil {
		return nil, err
	}

	return &downloadSvc{client}, nil
}

func (svc *downloadSvc) Close() error {
	errs := svc.client.Close()

	errorsMsg := ""
	for _, err := range errs {
		errorsMsg = fmt.Sprintf("%s\n%s", errorsMsg, err.Error())
	}

	if errorsMsg != "" {
		return errors.New(errorsMsg)
	}

	return nil
}

func (svc *downloadSvc) ListDownloads() []models.Download {
	torrs := svc.client.Torrents()

	downloads := make([]models.Download, len(torrs))
	for i, torr := range torrs {
		downloads[i] = models.DownloadFromTorrent(torr)
	}

	return downloads
}

func (svc *downloadSvc) AddDownload(hashString string) error {
	infoHash, err := utils.MetafromHex(hashString)
	if err != nil {
		return err
	}

	torrent, new := svc.client.AddTorrentInfoHash(infoHash)
	if !new {
		fmt.Println("torrent already exists")
		return errors.New("torrent already exists")
	}

	go handleTorrent(torrent)
	return nil
}

func handleTorrent(torrent *torrent.Torrent) {
	<-torrent.GotInfo()
	fmt.Printf("Added torrent: %s\n", torrent.Name())

	torrent.DownloadAll()
	<-torrent.Complete.On()
	fmt.Printf("Downloaded torrent: %s\n", torrent.Name())
}
