#!/bin/bash
# this script is used to install all dependent tools
# based on tools.go file content from $1 argument
PATTERN="_ \"(.+)\""
while read line ; do
  if [[ $line =~ $PATTERN ]]; then
    PKG=${BASH_REMATCH[1]}
    echo "installing" $PKG "with GO111MODULE" $GO111MODULE
    go install $PKG
  fi
done < $1