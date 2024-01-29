package services

import (
	"thor/internal/config"
	"thor/internal/services/download"
)

type Services struct {
	Download download.IDownload

	Configs *config.ServerConfigs
}

func NewServices(cfg *config.ServerConfigs) (*Services, error) {
	torrentSvc, err := download.Factory(cfg)

	if err != nil {
		return nil, err
	}

	return &Services{torrentSvc, cfg}, nil
}
