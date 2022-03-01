#!/usr/bin/env bash

# shellcheck disable=SC2034
version=2.0.6

cd ../
make Makefile build-linux-drawin
# shellcheck disable=SC2164
cd ./build/package
rm -rf ./jobor-darwin-amd64/bin/jobor
rm -rf ./jobor-linux-amd64/bin/jobor
mv linux-amd64 ./jobor-linux-amd64/bin/jobor
mv darwin-amd64 ./jobor-darwin-amd64/bin/jobor

tar zcvf jobor-linux-amd64-${version}.tar.gz ./jobor-linux-amd64/
tar zcvf jobor-darwin-amd64-${version}.tar.gz ./jobor-darwin-amd64/