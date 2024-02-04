#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peerinteraction): Add processIncomingMessages function

Introduce processIncomingMessages function in the peerinteraction package to handle incoming messages from peers. The function reads and processes various message types, updating peer state and handling piece job results. It also communicates with the seed package for handling seed requests. Update import paths to use the correct package names for download, messageutils, peer, peercommunication, piecehandler, and seed.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
