#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(peercommunication): Implement Message Struct and Methods

Introduce a Message struct representing P2P messages, along with methods to convert messages to bytes and obtain a human-readable string representation of the message type. This provides a foundation for handling communication between peers in a peer-to-peer network.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
