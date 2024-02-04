#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peerinteraction): Add readHandShake and sendChoke functions

Introduce readHandShake function to handle reading the handshake from a peer and verifying its correctness. The function checks the handshake's info hash against the manifest and prints status messages accordingly.

Add sendChoke function to send a choke message to a peer. The function utilizes SendMessageWithRetry from messageutils package and prints status messages based on the outcome.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
