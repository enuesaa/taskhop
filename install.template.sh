#!/bin/bash
set -e

# Replace version number on github actions.
# VERSION="{{VERSION}}"
VERSION="0.0.2"

OS=$(uname | tr '[:upper:]' '[:lower:]');
ARCH=$(uname -m);

if [ "$ARCH" = "x86_64" ]; then
  ARCH="amd64";
elif [[ "$ARCH" == arm* ]]; then
  ARCH="arm64";
else
  echo "Error: Unsupported arch $ARCH";
  exit 1;
fi

URL="https://github.com/enuesaa/taskhop/releases/download/v${VERSION}/taskhop_${VERSION}_${OS}_${ARCH}.tar.gz";
echo "Downloading $URL";
echo "";

curl -L "$URL" | tar -xz -C ~/tmp

echo "";
echo "Installation completed!";
