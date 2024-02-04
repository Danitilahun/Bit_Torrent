#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(download): Add DownloadJob struct

Introduce the DownloadJob struct representing a job to download a piece of a file in a torrent.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
