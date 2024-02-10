package main

import (
	"crypto/rand"
	"fmt"

	"github.com/Danitilahun/Bit_Torrent/bitfield"
	"github.com/Danitilahun/Bit_Torrent/common"
	"github.com/Danitilahun/Bit_Torrent/download"
	"github.com/Danitilahun/Bit_Torrent/fileUtils"
	messageutils "github.com/Danitilahun/Bit_Torrent/messageUtils"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/peerinteraction"
	"github.com/Danitilahun/Bit_Torrent/piecehandler"
	"github.com/Danitilahun/Bit_Torrent/seed"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
	"github.com/Danitilahun/Bit_Torrent/tracker"
	"log"
	mrand "math/rand"
	"net"
	"time"
)

func main() {
	manifest := fileUtils.ReadManifestFromFile("debian-11.6.0-amd64-netinst.iso.torrent")

	// Create files
	blobFile := fileUtils.LoadOrCreateDownloadBlob(&manifest)
	memCopy := make([]byte, manifest.TotalLength)
	blobFile.ReadAt(memCopy, 0)

	// Load progress from persistent storage
	currentBitField, bitfieldFile := bitfield.LoadOrCreateBitFieldFromFile(&manifest)

	totalDownloaded := countDownloadedPieces(currentBitField)

	fmt.Println("Total downloaded", totalDownloaded)

	// Get peers list
	id := [20]byte{}
	rand.Read(id[:])

	peerAddresses, err := tracker.GetPeersList(manifest, id, common.Port)
	if err != nil {
		fmt.Println("Can't get peers", err)
		panic(err)
	}
	// Create an empty peerAddresses slice
	// peerAddresses := []peer.PeerAddress{}

	// Append the new address
	// newAddress := peer.PeerAddress{IP: net.ParseIP("192.168.137.13"), Port: 6881}
	// peerAddresses = append(peerAddresses, newAddress)
	// Print the slice containing only the new address
	fmt.Println(peerAddresses)

	// channels
	workChannel := make(chan download.PieceJob, len(manifest.PieceHashes))
	pieceJobResultChannel := make(chan *download.PieceJobResult)
	seedRequestChannel := make(chan *seed.SeedRequest)

	// Call the function to create work for each piece
	createWorkForPieces(&manifest, currentBitField, &workChannel)

	// create common structure for leecher and seeder
	peers := make([]*peer.Peer, len(peerAddresses))

	for i, peerAddress := range peerAddresses {
		go peerinteraction.StartPeerWorker(peers, i, peerAddress, id, manifest, common.Port, &workChannel, currentBitField, &pieceJobResultChannel, &seedRequestChannel, nil)
	}

	// Listen for seeding requests
	go func() {
		for {
			seedRequest := <-seedRequestChannel
			// handle seeding request
			go seed.HandleSeedingRequest(seedRequest, memCopy, currentBitField, &manifest)
		}
	}()

	// Start the seeding server in a separate goroutine
	go startSeedingServer(&peers, id, manifest, common.Port, &workChannel, currentBitField, &pieceJobResultChannel, &seedRequestChannel)

	// Optimistic Unchoking
	go optimisticUnchoking(peers)

	// process results
	for {
		pieceJobResult := <-pieceJobResultChannel
		if pieceJobResult == nil {
			continue
		}

		copy(memCopy[pieceJobResult.PieceIndex*int(manifest.PieceLength):], pieceJobResult.PieceData)
		// write piece to file
		piecehandler.WritePieceToFile(&manifest, pieceJobResult, blobFile)
		// update bitfield
		currentBitField.MarkPiece(pieceJobResult.PieceIndex)
		currentBitField.WriteToFile(&manifest, bitfieldFile)

		// update progress
		totalDownloaded++
		fmt.Printf("Downloaded %v/%v pieces\n", totalDownloaded, len(manifest.PieceHashes))

		// send have message to all peers
		for _, peer := range peers {
			if peer != nil {
				messageutils.SendHaveMessage(peer, pieceJobResult.PieceIndex)
			}
		}

		// check if download is finished
		if totalDownloaded == len(manifest.PieceHashes) {
			fmt.Println("Download finished")
			fileUtils.WriteBlobToFiles(&manifest)
		}
	}
}

func optimisticUnchoking(peers []*peer.Peer) {
	for {
		if len(peers) != 0 {
			// unchoke random peer
			peerIndex := mrand.Intn(len(peers))
			if peers[peerIndex] != nil {
				if peers[peerIndex].IsChoked {
					peers[peerIndex].IsChoked = false
					go messageutils.SendUnchokeMessage(peers[peerIndex])
				}
			}
		}
		time.Sleep(31 * time.Second)
	}
}

func startSeedingServer(peers *[]*peer.Peer, id [20]byte, manifest torrentmodels.TorrentManifest, port int, workChannel *chan download.PieceJob, currentBitField *bitfield.Bitfield, pieceJobResultChannel *chan *download.PieceJobResult, seedRequestChannel *chan *seed.SeedRequest) {
	ListenAddr := ":" + fmt.Sprint(port)
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Listening on %s...\n", ListenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		peerInstance := peer.Peer{
			Address: peer.PeerAddress{
				IP:   conn.RemoteAddr().(*net.TCPAddr).IP,
				Port: uint16(port),
			},
			Conn:       conn,
			Interested: false,
			IsChoked:   true,
			IsChoking:  false,
			BitField:   make([]byte, len(manifest.PieceHashes)),
		}

		*peers = append(*peers, &peerInstance)

		go peerinteraction.StartPeerWorker(*peers, len(*peers)-1, peerInstance.Address, id, manifest, port, workChannel, currentBitField, pieceJobResultChannel, seedRequestChannel, &conn)
	}
}

// createWorkForPieces generates download jobs for each piece in the manifest that hasn't been downloaded yet.
func createWorkForPieces(manifest *torrentmodels.TorrentManifest, currentBitField *bitfield.Bitfield, workChannel *chan download.PieceJob) {
	for index, hash := range manifest.PieceHashes {
		// Ignore already downloaded pieces
		if !currentBitField.HasPiece(index) {
			*workChannel <- download.PieceJob{
				PieceIndex:  index,
				PieceHash:   hash,
				PieceLength: piecehandler.GetPieceLength(index, int(manifest.PieceLength), int(manifest.TotalLength)),
			}
		}
	}
}

func countDownloadedPieces(bitField *bitfield.Bitfield) int {
	totalDownloaded := 0

	for _, piece := range *bitField {
		for i := 0; i < 8; i++ {
			if piece&(1<<uint(i)) != 0 {
				totalDownloaded++
			}
		}
	}

	return totalDownloaded
}
