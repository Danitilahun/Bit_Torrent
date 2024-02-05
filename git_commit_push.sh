#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
refactor: Organize code into separate files

This commit refactors the main application code by organizing it into separate files for better code structure and maintainability. The main functionality remains unchanged, but the code is now divided into logical components.

- The main application logic is kept in `main.go`.
- The optimistic unchoking logic is moved to `optimistic_unchoking.go`.
- The seeding server logic is moved to `seeding_server.go`.
- Piece-related functions (creating work for pieces and counting downloaded pieces) are moved to `piece_work.go`.

These changes aim to improve code readability, separation of concerns, and ease of future modifications.

"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
