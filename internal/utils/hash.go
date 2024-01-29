package utils

import (
	"fmt"

	"github.com/anacrolix/torrent/metainfo"
)

func MetafromHex(infohash string) (h metainfo.Hash, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error parsing string to InfoHash")
		}
	}()

	h = metainfo.NewHashFromHex(infohash)

	return h, nil
}
