#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m " 
feat(torrentmodels): 
add FileMetadata struct for representing file metadata in torrents

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
