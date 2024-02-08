package download

// DownloadJob represents a job to download a piece of a file in a torrent.
type PieceJob struct {
	// PieceIndex is the index of the piece in the torrent.
	PieceIndex int
	// PieceHash is the SHA-1 hash of the piece.
	PieceHash [20]byte
	// PieceLength is the length of the piece in bytes.
	PieceLength int
}
