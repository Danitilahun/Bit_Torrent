#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(fileUtils): Add utility functions for file and blob handling

Introduce utility functions in the fileUtils package for managing files and blobs in a BitTorrent client. Includes functions to load or create download blobs, write blob data to files based on manifest information, and read a torrent manifest from a file. Update import paths to use the correct package names.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
