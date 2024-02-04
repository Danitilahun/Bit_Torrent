#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
docs(piecehandler): Add package for handling and processing pieces in BitTorrent protocol

Introduce a new package named 'piecehandler' containing utility functions for managing and processing pieces within the BitTorrent protocol. Each function is documented to explain its purpose and usage.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
