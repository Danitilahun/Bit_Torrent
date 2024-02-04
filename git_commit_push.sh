#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(common): Update message utility functions and imports

Extend message utility functions in the common package to support custom message types defined in the peercommunication package. Update import paths to use the correct package names for peer and peercommunication. The functions include reading and sending various message types such as Have, Unchoke, Choke, and custom messages. Additionally, there are functions for reading request messages.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
