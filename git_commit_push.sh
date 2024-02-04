#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(tracker): Add GetPeersList function and related utilities

Introduce GetPeersList function in the tracker package to retrieve a list of peers from specified trackers in the torrent manifest. This includes functions to generate the tracker request URL, send an HTTP request to the tracker, and extract peer information from the tracker response. Additionally, update imports to use the correct package names.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
