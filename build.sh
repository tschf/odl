#!/bin/bash
BASEDIR=$GOPATH/src/github.com/tschf/odl/dist

rm -rf ${BASEDIR}

echo "Building for OS X (darwin) 32 bit"
export GOOS=darwin
export GOARCH=386
mkdir -p ${BASEDIR}/osx_32
go build
mv odl ${BASEDIR}/osx_32/odl
echo "Done"

echo "Building for OS X (darwin) 64 bit"
export GOOS=darwin
export GOARCH=amd64
mkdir -p ${BASEDIR}/osx_64
go build
mv odl ${BASEDIR}/osx_64/odl
echo "Done"

echo "Building for Linux 32 bit"
export GOOS=linux
export GOARCH=386
mkdir -p ${BASEDIR}/linux_32
go build
mv odl ${BASEDIR}/linux_32/odl
echo "Done"

echo "Building for Linux 64 bit"
export GOOS=linux
export GOARCH=amd64
mkdir -p ${BASEDIR}/linux_64
go build
mv odl ${BASEDIR}/linux_64/odl
echo "Done"

echo "Building for Windows 32 bit"
export GOOS=windows
export GOARCH=386
mkdir -p ${BASEDIR}/win_32
go build
mv odl.exe ${BASEDIR}/win_32/odl.exe
echo "Done"

echo "Building for Windows 64 bit"
export GOOS=windows
export GOARCH=amd64
mkdir -p ${BASEDIR}/win_64
go build
mv odl.exe ${BASEDIR}/win_64/odl.exe
echo "Done"
