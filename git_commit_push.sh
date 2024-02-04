#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peercommunication): Define MessageType Enumeration

Create a MessageType enumeration to represent different types of messages in peer-to-peer communication. This enum includes constants for various message types used in P2P protocols, facilitating clear and consistent identification of message types.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
