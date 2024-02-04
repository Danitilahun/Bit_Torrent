#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
refactor(peers): Refactor package structure and naming

Update the package name to "peers" to reflect the entities represented in the code. Organize the code into separate files for Peer and PeerAddress to enhance clarity and maintainability.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
