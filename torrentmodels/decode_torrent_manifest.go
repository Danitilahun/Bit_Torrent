// File Name: decode_torrent_manifest.go
// Package Name: torrentmodels

package torrentmodels

import (
	"crypto/sha1"
	"github.com/IncSW/go-bencode"
	"path"
)

// DecodeTorrentManifest decodes the given data into a TorrentManifest struct.
func DecodeTorrentManifest(data interface{}) TorrentManifest {
	// Cast the data to a map[string]interface{} as it represents a JSON-like structure.
	torrentMap := data.(map[string]interface{})

	// Extract metadata fields from the top-level of the torrentMap.
	announce := string(torrentMap["announce"].([]byte))
	announceList := []string{}
	if torrentMap["announce-list"] != nil {
		// Extract and convert the announce-list to a list of strings.
		for _, item := range torrentMap["announce-list"].([]interface{}) {
			announceList = append(announceList, string(item.([]interface{})[0].([]byte)))
		}
	}
	comment := ""
	if torrentMap["comment"] != nil {
		// Extract the comment field.
		comment = string(torrentMap["comment"].([]byte))
	}
	createdBy := ""
	if torrentMap["created by"] != nil {
		// Extract the created by field.
		createdBy = string(torrentMap["created by"].([]byte))
	}

	// Extract the 'info' field, which contains detailed information about the torrent.
	info := torrentMap["info"].(map[string]interface{})

	// Extract various fields from the 'info' section.
	pieceLength := info["piece length"].(int64)
	torrentName := string(info["name"].([]byte))

	// Calculate the SHA-1 hash of the 'info' section to get the InfoHash.
	infoBytes, _ := bencode.Marshal(info)
	infoHash := sha1.Sum(infoBytes)

	// Extract the raw byte string representing the concatenated SHA-1 hashes of all pieces.
	pieces := info["pieces"].([]byte)

	// Parse the piece hashes into a slice of [20]byte.
	pieceHashes := [][20]byte{}
	for i := 0; i < len(pieces); i += 20 {
		var currentHash [20]byte
		copy(currentHash[:], pieces[i:i+20])
		pieceHashes = append(pieceHashes, currentHash)
	}

	// Extract information about individual files if available.
	filesMetadata := []FileMetadata{}
	var offset int64

	files := []interface{}{info}

	if info["files"] != nil {
		files = info["files"].([]interface{})
	}

	// Loop through files, extract metadata, and build FileMetadata objects.
	for _, file := range files {
		file := file.(map[string]interface{})

		filePathParts := []string{torrentName}

		if file["path"] != nil {
			// If the 'path' field exists, extract and convert it to a list of strings.
			for _, part := range file["path"].([]interface{}) {
				filePathParts = append(filePathParts, string(part.([]byte)))
			}
		} else {
			// If 'path' is not available, use the torrentName as the file path.
			filePathParts = append(filePathParts, torrentName)
		}

		// Extract the file size and create a FileMetadata object.
		fileSize := file["length"].(int64)
		filesMetadata = append(filesMetadata, FileMetadata{
			FilePath:   path.Join(filePathParts...),
			FileName:   filePathParts[len(filePathParts)-1],
			FileSize:   fileSize,
			FileOffset: offset,
		})

		// Update the offset for the next file.
		offset += fileSize
	}

	// Construct and return a TorrentManifest object with all extracted information.
	return TorrentManifest{
		Announce:      announce,
		AnnounceList:  announceList,
		InfoHash:      infoHash,
		PieceHashes:   pieceHashes,
		PieceLength:   pieceLength,
		TotalLength:   offset,
		TorrentName:   torrentName,
		Comment:       comment,
		CreatedBy:     createdBy,
		FilesMetadata: filesMetadata,
	}
}
