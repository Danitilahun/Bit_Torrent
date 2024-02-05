# Bit-Torrent Client Using Golang

## Setup

1. **Initialize Go Modules:**

   - Run the following command to download and install dependencies:
     ```
     go mod tidy
     ```

2. **Add Torrent File in `main.go`:**

   - Open the `main.go` file and include a valid `.torrent` file. For example, you can use a Debian ISO file.

3. **Run the Main Program:**
   - Execute the following command to start the GoTorrent client:
     ```
     go run main.go
     ```

These steps ensure that the necessary dependencies are fetched, a valid torrent file is added to the `main.go` file, and the main program is executed to run the GoTorrent client. Adjustments might be needed based on the specifics of your project structure and configuration.

# Torrent Client Project Features

## Overview

This project entails the creation of a command-line torrent client utilizing Go. The client enables seamless file downloads and uploads via the BitTorrent protocol, harnessing Go's built-in concurrency and networking capabilities for optimal performance.

## Key Components

### Efficient Download/Upload

- The client adeptly handles extensive file transfers, leveraging Go's concurrency features for simultaneous downloads and uploads.

### Parser Accuracy

- Torrent files are parsed with precision, extracting vital information such as tracker URL, file name, size, and pieces.

### Robust Tracker Communication

- Establishes reliable communication with the tracker, ensuring up-to-date information on available peers.

### Piece-Level Management

- Effective piece management guarantees a systematic download and upload process, ensuring file completeness.

### Concurrency for Scalability

- Utilizes Go's concurrency features to seamlessly handle multiple downloads and uploads concurrently.

### Error Resilience

- Robust error handling mechanisms ensure the client gracefully manages unexpected scenarios, maintaining stability.

### Fault Tolerance and Recovery

- The client exhibits resilience, recovering from errors and persisting bitfield information to ensure continuity in processes.

## Limitations

The torrent client has the following limitations:

1. **File Type Support:**

   - Only supports `.torrent` files.
   - Does not support magnet links.

2. **Tracker Protocol:**

   - Only supports HTTP trackers.
   - Does not support other tracker protocols.

3. **Torrent Structure:**
   - Does not support multi-file torrents.
   - Currently designed for single-file torrents.

These limitations outline the current scope of the torrent client and represent areas where future improvements or additional features could be considered.

In summary, the project delivers a functional and efficient torrent client, demonstrating prowess in efficient file transfers, accurate parsing, robust communication, and resilience in the face of errors.

## Group Members

| Name           | ID          | Section |
| -------------- | ----------- | ------- |
| Bilen Mehalek  | UGR/0252/13 | 1       |
| Daniel Tilahun | UGR/2557/13 | 1       |
| Dawit Minale   | UGR/7990/13 | 2       |
| Rihana Ersanu  | UGR/8031/13 | 2       |
| Robel Tesfaye  | UGR/8429/13 | 1       |

[GitHub Repo](https://github.com/Danitilahun/Bit_Torrent)
