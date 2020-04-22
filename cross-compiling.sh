#!/bin/bash

export GOPATH=$PWD

echo
echo "This is a demo on cross compiling"
echo
echo "This script will try to build the release for yq"
echo
echo "Typically, as a good maintainer in an open source project"
echo "You need to build different pre-built binaries in different OS and Arch for all users"
echo "In this script, it will demostrate how easy you can build binary in common OS and Arch"

echo
echo "Now, getting yq package... please wait a moment to let go download all the packages..."
echo

go get github.com/mikefarah/yq

echo
echo "yq has already downloaded in $GOPATH/src/github.com/mikefarah/yq"
echo "And you are supposed to see a binary built by go get in $GOPATH/src/bin"
echo


echo
echo "Starting from now, we will try to build different binaries with GOOS and GOARCH"
echo

cd $GOPATH/src/github.com/mikefarah/yq
for arch in "amd64" "386"; do
  for os in "darwin" "linux" "windows"; do
    echo "Building yq with $os/$arch"
    if [[ $os == $windows ]]; then
        GOOS=$os GOARCH=$arch go build -o yq-$os-$arch.exe
    else
        GOOS=$os GOARCH=$arch go build -o yq-$os-$arch
    fi
  done
done

echo
echo "You can now see all the binaries in $PWD"
echo

ls -l yq-*

echo
echo "When you try to examine the binaries with command file, you can verify their arch and os"
echo

for binary in yq-*; do
    echo "For $binary, you can get this message with file command:"
    file $binary
    echo
done
