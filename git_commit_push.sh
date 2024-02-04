#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peerutils): Add ConnectToPeer and EstablishConnection functions

Introduce ConnectToPeer function to establish a TCP connection to a peer and EstablishConnection function to connect to a peer and return a Peer instance. Additionally, update the import paths to use the correct package names.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
