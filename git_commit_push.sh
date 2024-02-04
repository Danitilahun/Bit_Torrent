#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
chore(peers): Organize code, move PeerAddress to separate file

Move the definition of the PeerAddress type to its own file (peer_address.go) to improve code organization and maintainability. No functional changes.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
