package fileUtils

import (
	"fmt"
	"io"
	"os"

	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
	"github.com/IncSW/go-bencode"
)

func LoadOrCreateDownloadBlob(manifest *torrentmodels.TorrentManifest) *os.File {
	os.MkdirAll(manifest.TorrentName, 0700)
	blobFilePath := manifest.TorrentName + "/" + manifest.TorrentName + ".blob"

	if _, err := os.Stat(blobFilePath); os.IsNotExist(err) {
		_, err = os.Create(blobFilePath)
		if err != nil {
			panic(err)
		}
	}

	blobFile, err := os.OpenFile(blobFilePath, os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Can't create file", manifest.TorrentName, err)
		panic(err)
	}

	return blobFile
}

func WriteBlobToFiles(manifest *torrentmodels.TorrentManifest) {
	blobFilePath := manifest.TorrentName + "/" + manifest.TorrentName + ".blob"

	if _, err := os.Stat(blobFilePath); os.IsNotExist(err) {
		_, err = os.Create(blobFilePath)
		if err != nil {
			panic(err)
		}
	}

	blobFile, err := os.OpenFile(blobFilePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Can't open blob file", manifest.TorrentName, err)
		panic(err)
	}

	defer blobFile.Close()

	for _, file := range manifest.FilesMetadata {
		if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
			_, err = os.Create(file.FilePath)
			if err != nil {
				panic(err)
			}
		}

		f, err := os.OpenFile(file.FilePath, os.O_RDWR, 0644)

		if err != nil {
			fmt.Println("Can't create file", file.FilePath, err)
			panic(err)
		}

		err = f.Truncate(file.FileSize)
		if err != nil {
			fmt.Println("Can't truncate file", file.FilePath, err)
			panic(err)
		}
		blobFile.Seek(file.FileOffset, 0)

		_, err = io.CopyN(f, blobFile, file.FileSize)

		if err != nil {
			fmt.Println("Can't copy file", file.FilePath, err)
			panic(err)
		}
		f.Close()
	}
}

func ReadManifestFromFile(filePath string) torrentmodels.TorrentManifest {
	content, err := os.ReadFile(filePath)
	if err != nil {
		println("Can't open torrent file", err)
		panic(err)
	}
	data, err := bencode.Unmarshal(content)
	if err != nil {
		panic(err)
	}
	return torrentmodels.DecodeTorrentManifest(data)
}
