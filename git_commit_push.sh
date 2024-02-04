#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(download): Add DownloadJobProgress struct

Introduce the DownloadJobProgress struct representing the progress of downloading a piece.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
