#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
Implement the main BitTorrent client logic with concurrency, downloading, seeding, tracker communication, and optimistic unchoking. Handles work distribution, peer connections, piece handling, and updates the bitfield and progress accordingly. Utilizes channels for communication between goroutines and manages seeding requests. Implements a listener for seeding server, optimistic unchoking, and processes results as pieces are downloaded, updating the total downloaded count and sending 'have' messages to peers. Implements graceful shutdown and write completed blob to files upon finishing the download.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
