#!/bin/sh
#
# Builds simurg for Linux and macOS on amd64 and arm64 into dist/, with
# VERSION written into the binary. Then runs the binary built for this
# machine and checks that it reports VERSION, which proves the injection.
# Runs inside the Go image; nothing is published.
#
# Usage: VERSION=<version> build.sh

set -eu

readonly version="${VERSION:?VERSION is not set}"
readonly out=dist

rm -rf "$out"
for os in linux darwin; do
    for arch in amd64 arm64; do
        CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -trimpath \
            -ldflags "-s -w -X main.version=$version" \
            -o "$out/simurg-$os-$arch" .
        echo "built $out/simurg-$os-$arch"
    done
done

native="$out/simurg-$(go env GOOS)-$(go env GOARCH)"
got=$("$native" version)
want="simurg $version"
if [ "$got" != "$want" ]; then
    echo "FAIL: $native version printed '$got', want '$want'" >&2
    exit 1
fi
echo "$native version: $got"
