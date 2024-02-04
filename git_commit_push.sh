#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(common): Add handshake utility functions

Introduce utility functions in the common package for handling BitTorrent protocol handshakes. Includes ReadHandShake function to read a handshake from a reader and EstablishHandShake function to establish a handshake with a peer. Update import paths to use the correct package names for handshake, peer, and torrentmodels.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
