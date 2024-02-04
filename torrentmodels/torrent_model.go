package torrentmodels

// TorrentManifest represents the metadata information of a torrent.
type TorrentManifest struct {
	PieceHashes   [][20]byte     // An array of SHA-1 hashes for each piece of the torrent file.
	Announce      string         // The URL of the tracker server.
	AnnounceList  []string       // A list of backup tracker URLs.
	InfoHash      [20]byte       // The SHA-1 hash of the 'info' section of the torrent.
	PieceLength   int64          // The size (in bytes) of each piece in the torrent.
	TotalLength   int64          // The total size (in bytes) of all files in the torrent.
	TorrentName   string         // The name of the torrent.
	Comment       string         // Any additional comments or information.
	CreatedBy     string         // The name of the software or tool that created the torrent.
	FilesMetadata []FileMetadata // Information about individual files in the torrent.
}
