#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat: Simulate Bitfield Usage in Main Example

Demonstrate the usage of the bitfield package in a simple scenario:
- Simulate a manifest with 10 pieces.
- Load or create a Bitfield from/to a file.
- Mark additional pieces (1, 4, 7).
- Print and visualize the Bitfield status.
- Write the updated Bitfield to the file.

This provides an example of how the Bitfield package can be used to manage file piece availability.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
