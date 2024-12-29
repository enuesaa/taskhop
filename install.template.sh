#!/bin/bash

# Replace `${VERSION}` with the latest value in GitHub Actions.
VERSION="${VERSION}"; # like `v0.0.3`

if [[ "$VERSION" != v* ]]; then
  echo "Error: Incorrect install script.";
  exit 1;
fi

# strip `v`
VERSION_NUMBER="${VERSION#v}"; # like `0.0.3`

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

URL="https://github.com/enuesaa/taskhop/releases/download/v${VERSION_NUMBER}/taskhop_${VERSION_NUMBER}_${OS}_${ARCH}.tar.gz";
echo "Downloading $URL";
echo "";

curl -L "$URL" | tar -xz -C ~/tmp;

echo "";
echo "Installation completed!";
