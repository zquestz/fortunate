#/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go install github.com/cratonica/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go get github.com/cratonica/2goarray
        exit
    fi
fi

OUTPUT=icon_unix.go
echo Generating $OUTPUT
echo "//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd" > $OUTPUT
echo "// +build linux darwin dragonfly freebsd netbsd openbsd" >> $OUTPUT
echo >> $OUTPUT
cat "lotus.png" | $GOPATH/bin/2goarray Data icon >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi

OUTPUT=icon_light_unix.go
echo Generating $OUTPUT
echo "//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd" > $OUTPUT
echo "// +build linux darwin dragonfly freebsd netbsd openbsd" >> $OUTPUT
echo >> $OUTPUT
cat "lotus-light.png" | $GOPATH/bin/2goarray DataLight icon >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi

OUTPUT=icon_dark_unix.go
echo Generating $OUTPUT
echo "//go:build linux || darwin || dragonfly || freebsd || netbsd || openbsd" > $OUTPUT
echo "// +build linux darwin dragonfly freebsd netbsd openbsd" >> $OUTPUT
echo >> $OUTPUT
cat "lotus-dark.png" | $GOPATH/bin/2goarray DataDark icon >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished

OUTPUT_WINDOWS=icon_windows.go
echo Generating $OUTPUT_WINDOWS
echo "//go:build windows" > $OUTPUT_WINDOWS
echo "// +build windows" >> $OUTPUT_WINDOWS
echo >> $OUTPUT_WINDOWS
cat "lotus.ico" | $GOPATH/bin/2goarray Data icon >> $OUTPUT_WINDOWS
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT_WINDOWS
    exit
fi

OUTPUT_WINDOWS=icon_light_windows.go
echo Generating $OUTPUT_WINDOWS
echo "//go:build windows" > $OUTPUT_WINDOWS
echo "// +build windows" >> $OUTPUT_WINDOWS
echo >> $OUTPUT_WINDOWS
cat "lotus-light.ico" | $GOPATH/bin/2goarray DataLight icon >> $OUTPUT_WINDOWS
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT_WINDOWS
    exit
fi

OUTPUT_WINDOWS=icon_dark_windows.go
echo Generating $OUTPUT_WINDOWS
echo "//go:build windows" > $OUTPUT_WINDOWS
echo "// +build windows" >> $OUTPUT_WINDOWS
echo >> $OUTPUT_WINDOWS
cat "lotus-dark.ico" | $GOPATH/bin/2goarray DataDark icon >> $OUTPUT_WINDOWS
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT_WINDOWS
    exit
fi
echo Finished
