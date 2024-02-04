#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(torrentmodels): add TorrentManifest struct and related functionality
  Introduces a new Go struct, TorrentManifest, within the torrentmodels package. This struct represents the metadata information of a torrent, including fields such as PieceHashes, Announce, InfoHash, etc. Also includes a FileMetadata struct and decoding functionality for torrent manifests.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
