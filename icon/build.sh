#!/bin/sh

if [ -z "$GOPATH" ]; then
    echo "GOPATH environment variable not set"
    exit 1
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    if ! go install github.com/cratonica/2goarray@latest; then
        echo "Failure executing go install github.com/cratonica/2goarray@latest"
        exit 1
    fi
fi

# generate <output-file> <build-tag> <array-name> <source-image>
generate() {
    output="$1"
    build_tag="$2"
    name="$3"
    src="$4"

    echo "Generating $output"
    echo "$build_tag" > "$output"
    echo >> "$output"
    if ! "$GOPATH/bin/2goarray" "$name" icon < "$src" >> "$output"; then
        echo "Failure generating $output"
        exit 1
    fi
    gofmt -s "$output" > "$output.formatted"
    mv "$output.formatted" "$output"
}

UNIX_TAG="//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd"
WINDOWS_TAG="//go:build windows"

generate large_icon_unix.go    "$UNIX_TAG"    DataLarge ../Icon.png
generate icon_unix.go          "$UNIX_TAG"    Data      fortunate.png
generate icon_light_unix.go    "$UNIX_TAG"    DataLight fortunate-light.png
generate icon_dark_unix.go     "$UNIX_TAG"    DataDark  fortunate-dark.png

generate large_icon_windows.go "$WINDOWS_TAG" DataLarge ../Icon.ico
generate icon_windows.go       "$WINDOWS_TAG" Data      fortunate.ico
generate icon_light_windows.go "$WINDOWS_TAG" DataLight fortunate-light.ico
generate icon_dark_windows.go  "$WINDOWS_TAG" DataDark  fortunate-dark.ico

echo "Finished"
