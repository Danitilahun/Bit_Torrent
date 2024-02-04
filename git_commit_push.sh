#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(seed): Add SeedRequest struct

Introduce the SeedRequest struct in the seed package, representing a request from a peer in a seed scenario. This struct includes a reference to the peer making the request and the corresponding peer communication message.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
