#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
Separate handshake package into struct and I/O files

- Move the HandShake struct and its methods to handshake.go
- Move I/O related functions (ToBytes, ReadHandShake) to handshake_io.go
- Organize code for better maintainability

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
