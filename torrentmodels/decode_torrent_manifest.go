// File Name: decode_torrent_manifest.go
// Package Name: torrentmodels

package torrentmodels

import (
	"crypto/sha1"
	"fmt"
	"path"

	"github.com/IncSW/go-bencode"
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
	// Marshaling (Serialization): Marshaling is the process of converting in-memory
	// data structures or objects into a format (such as JSON, XML, or binary) that
	// can be easily stored or transmitted. It is commonly used when you want to save or send data.

	// Bencode is a simple binary encoding format used primarily in the context of BitTorrent, a
	// peer-to-peer file-sharing protocol. It is designed to represent dictionaries (hash tables),
	// lists, integers, and byte strings in a compact binary format. Bencode is used to encode the
	// metadata of torrent files, describing the structure and content of the files being shared in a BitTorrent network.

	infoBytes, _ := bencode.Marshal(info)

	// The SHA-1 hash is commonly used in various cryptographic and data integrity verification scenarios.
	// A checksum is a value derived from the data in a way that is relatively easy to calculate and that
	// ideally changes significantly when the data changes. Checksums are commonly used in various fields
	//  for error checking, data integrity verification, and detecting changes in data.

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
		fmt.Println("Files exist")
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
