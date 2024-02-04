#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(download): Add DownloadJobResult struct

Introduce the DownloadJobResult struct representing the result of a successfully downloaded piece.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
