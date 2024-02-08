package tracker

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/Danitilahun/Bit_Torrent/peer"
	"github.com/Danitilahun/Bit_Torrent/torrentmodels"
	"github.com/IncSW/go-bencode"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Trackers are servers that help coordinate the distribution of files amongst peers in the BitTorrent network.
// The URL includes various query parameters that provide the tracker with information about the peer making the
//  request and the torrent they are participating in.

func GetPeersList(manifest torrentmodels.TorrentManifest, peerId [20]byte, port int) (peers []peer.PeerAddress, err error) {
	fmt.Printf("Getting peers list from trackers\n")
	announcer := []string{manifest.Announce}
	announcer = append(announcer, manifest.AnnounceList...)
	var trackerResponse interface{} = nil

	for _, announce := range announcer {
		announceUrl, err := getTrackerRequestUrl(manifest, announce, peerId, port)
		if err != nil {
			continue
		}
		response, err := getTrackerResponse(announceUrl)

		if err != nil {
			continue
		}
		fmt.Printf("Got peers list from tracker %v\n", announce)
		trackerResponse = response
		break
	}

	if trackerResponse == nil {
		return nil, errors.New("can't get peers from any tracker")
	}
	peers, err = getPeersFromTrackerResponse(trackerResponse)
	return
}

func getPeersFromTrackerResponse(trackerResponse interface{}) (peers []peer.PeerAddress, err error) {
	receivedPeers := trackerResponse.(map[string]interface{})["peers"].([]byte)

	if len(receivedPeers)%6 != 0 {
		return nil, errors.New("invalid peers list")
	}

	for i := 0; i < len(receivedPeers); i += 6 {
		peers = append(peers, peer.PeerAddress{
			IP:   receivedPeers[i : i+4],
			Port: binary.BigEndian.Uint16(receivedPeers[i+4 : i+6]),
		})
	}

	return
}

func getTrackerResponse(announceUrl string) (trackerResp interface{}, err error) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(announceUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := make([]byte, 2048)
	resp.Body.Read(data)

	trackerResp, err = bencode.Unmarshal(data)
	return
}

func getTrackerRequestUrl(manifest torrentmodels.TorrentManifest, announce string, peerId [20]byte, port int) (string, error) {
	baseUrl, err := url.Parse(announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash":  []string{string(manifest.InfoHash[:])},
		"peer_id":    []string{string(peerId[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(int(manifest.TotalLength))},
	}

	baseUrl.RawQuery = params.Encode()
	return baseUrl.String(), nil
}
