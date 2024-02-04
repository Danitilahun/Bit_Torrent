#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
docs(handshake): Explain HandShake struct and methods

Provide documentation for the HandShake struct in the handshake package, detailing the purpose of each field (HeaderText, InfoHash, PeerId) and explaining the New constructor and ToBytes method. Clarify the structure of a BitTorrent protocol handshake message in the comments.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
