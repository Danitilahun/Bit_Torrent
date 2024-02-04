#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(torrentmodels): add DecodeTorrentManifest function

Introduces a new function DecodeTorrentManifest in the torrentmodels package. This function is responsible for decoding the metadata of a torrent, extracting information such as announce URL, announce list, comments, created by, and detailed file information. The decoded data is used to create a TorrentManifest struct, providing a comprehensive representation of the torrent's metadata.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
