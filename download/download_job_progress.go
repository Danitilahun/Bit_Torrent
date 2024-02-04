package download

// DownloadJobProgress represents the progress of downloading a piece.
type PieceJobProgress struct {
	// PieceIndex is the index of the piece being downloaded.
	PieceIndex int
	// Buffer is the portion of the piece data received so far.
	Buffer []byte
	// TotalDownloaded is the total number of bytes downloaded for the piece.
	TotalDownloaded int
	// PieceLength is the total length of the piece in bytes.
	PieceLength int
}
