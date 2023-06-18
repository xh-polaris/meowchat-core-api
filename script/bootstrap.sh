#!/bin/bash
CURDIR=$(cd "$(dirname "$0")" || exit; pwd)
BinaryName=meowchat.core-api
echo "$CURDIR/bin/${BinaryName}"
exec "$CURDIR/bin/${BinaryName}"
