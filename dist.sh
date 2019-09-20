#!/bin/bash
PWD=`pwd`
dir=$(PWD)
export GO111MODULE=on
BUILD_ID=`git describe --always --long`

if [[ -n ${VERSION} ]]; then
echo ${VERSION}
else
VERSION=`git describe --always --long`
fi

set -eu

rm -rf bin/dist
mkdir bin/dist

for os in "linux" "windows" "darwin"
do
  for arch in amd64 #386
  do
    export GOOS=$os
    export GOARCH=$arch
    go get -v all

        echo "Building for $os $arch"

        if [[ "$os" == "windows" ]]; then
          export ext='.exe'
        else
          export ext=''
        fi
        cd ${dir}
        name="microspector_${VERSION}_${os}_${arch}${ext}"
        go build -i -v -o ${PWD}/bin/dist/${name} -ldflags="-X main.Version=${VERSION} -X main.Build=${BUILD_ID}" ${PWD}/cmd
        cd ${PWD}/bin/dist
        tar -cvzf ${name}.tar.gz ${name}

  done
done

echo "All done."