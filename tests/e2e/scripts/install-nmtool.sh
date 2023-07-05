#!/bin/bash

if hash nmtool 2>/dev/null; then
  echo "[install-nmtool.sh] nmtool is already installed. skipping installation."
  exit 0
fi

echo "[install-nmtool.sh] installing nmtool."
cd nmtool || exit 1
make install
