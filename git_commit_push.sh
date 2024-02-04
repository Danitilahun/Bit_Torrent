#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat: Add TorrentRequestMessage struct to represent a message requesting data from a peer

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
