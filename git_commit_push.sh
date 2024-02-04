#!/bin/bash

# Add all changes
git add .

# Commit changes with the specified commit message
git commit -m "
feat(seed): Add HandleSeedingRequest function

Introduce HandleSeedingRequest function in the seed package to handle seeding requests from peers. The function validates the request, checks for choke status, and processes the requested piece data. Update import paths to use the correct package names for common and models.
"
# Push changes to the remote repository
git push origin HEAD

echo "Changes committed and pushed successfully."
