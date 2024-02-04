# Torrent Client Using Golang

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

## Project Strengths

- **Efficient Download/Upload**: Seamlessly handles large file transfers with optimal concurrency.
- **Parser Accuracy**: Extracts vital information accurately for precise functionality.

- **Robust Tracker Communication**: Ensures reliable communication with the tracker, keeping peer information up-to-date.

- **Piece-Level Management**: Systematic handling of file pieces for a smooth download and upload process.

- **Concurrency for Scalability**: Leverages Go's concurrency features for efficient handling of concurrent downloads and uploads.

- **Error Resilience**: Gracefully handles unexpected scenarios, ensuring a stable client.

- **Fault Tolerance and Recovery**: Recovers from errors and persists critical information, ensuring uninterrupted processes.

In summary, the project delivers a functional and efficient torrent client, demonstrating prowess in efficient file transfers, accurate parsing, robust communication, and resilience in the face of errors.

## Group Members

| Name           | ID          | Section |
| -------------- | ----------- | ------- |
| Bilen Mehalek  | UGR/0252/13 | 1       |
| Daniel Tilahun | UGR/2557/13 | 1       |
| Dawit Minale   | UGR/7990/13 | 2       |
| Rihana Ersanu  | UGR/8031/13 | 2       |
| Robel Tesfaye  | UGR/8429/13 | 1       |
