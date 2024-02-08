package download

// DownloadJobResult represents the result of a successfully downloaded piece.
type PieceJobResult struct {
	// PieceIndex is the index of the downloaded piece.
	PieceIndex int
	// PieceData is the actual data of the downloaded piece.
	PieceData []byte
}
