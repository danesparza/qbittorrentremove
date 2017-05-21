package main

// Torrent describes a torrent.
type Torrent struct {
	Hash          string  `json:"hash"`
	Name          string  `json:"name"`
	Priority      int     `json:"priority"`
	Ratio         float64 `json:"ratio"`
	Size          int64   `json:"size"`
	Seeds         int     `json:"num_seeds"`
	Leechs        int     `json:"num_leechs"`
	DownloadSpeed int     `json:"dlspeed"`
	UploadSpeed   int     `json:"upspeed"`

	// RawState is the state reported by qBittorrent.
	//
	// The possible states are:
	//
	// * pausedDL
	// * pausedUP
	// * stalledDL
	// * stalledUP
	// * queuedDL
	// * queuedUP
	// * downloading
	// * uploading
	RawState string `json:"state"`
}
