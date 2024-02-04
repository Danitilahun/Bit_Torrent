package torrentmodels

// FileMetadata represents metadata information about a file within a torrent.
type FileMetadata struct {
	// FilePath is the path to the file within the torrent.
	// It can be the full path or a relative path from the root of the torrent.
	FilePath string

	// FileName is the name of the file.
	// It provides a convenient way to access just the name of the file without the path.
	FileName string

	// FileSize indicates the size of the file in bytes.
	FileSize int64

	// FileOffset denotes the offset of the file within the torrent.
	// It is the position within the torrent where the file's data begins.
	FileOffset int64
}
