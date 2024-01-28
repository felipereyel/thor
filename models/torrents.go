package models

type Torrent struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var EmptyTorrent = Torrent{}
