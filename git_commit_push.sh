#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peerinteraction): Implement StartPeerWorker function

Add StartPeerWorker function to manage the interaction with a peer. It establishes the connection, handshake, and handles piece jobs. The function sends interested and unchoke messages, receives the bitfield, and interacts with the peer based on the piece job progress. It also uses the processIncomingMessages function for handling incoming messages and communicates with other channels like workChannel, pieceJobResultChannel, and seedRequestChannel.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
