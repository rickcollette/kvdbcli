#!/bin/bash

## Vars
VERSION_FILE="VERSION"

# Get the current version from the VERSION file
function GetVersion() {
    # Check if the VERSION file exists
    if [ ! -f "$VERSION_FILE" ]; then
        echo "$VERSION_FILE not found!"
        exit 1
    fi

    # Read the version from the VERSION file
    AppVersion="v$(cat ${VERSION_FILE})"

    # Check if version was found
    if [ -z "$AppVersion" ]; then
        echo "Version not found in $VERSION_FILE"
        exit 1
    fi

    echo "Current version: $AppVersion"
}

# Git functions
function GitFunctions() {
    git add .
    git commit -m"Preparing for ${AppVersion} release"
    git tag ${AppVersion}
    git push origin ${AppVersion}
    go list -m github.com/rickcollette/kvdbcli@${AppVersion}
}


# Run the commands

GetVersion
GitFunctions
