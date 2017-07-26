#!/bin/bash
if [ -z "$GITHUB_TOKEN" ]; then
    echo "Please export GITHUB_TOKEN variable"
    exit 1
fi
go get github.com/goreleaser/goreleaser
git fetch --tags
pwd=`git rev-parse --show-toplevel`
release=`git tag | tail -n 1`
version=`echo $release | awk -F \. '{ print $1 }'`
mayor=`echo $release | awk -F \. '{ print $2 }'`
minor=`echo $release | awk -F \. '{ print $3 }'`
new_minor=$(( $minor + 1 ))
new_release="$version.$mayor.$new_minor"
echo "package utils" > $PWD/utils/version.go
echo "" >> $PWD/utils/version.go
echo "const VERSION = \"$new_release\"" >> $PWD/utils/version.go
echo "+ git add $PWD/utils/version.go"
git add $PWD/utils/version.go
echo "+ git commit -m \"Release $new_release\""
git commit -m "Release $new_release"
echo "+ git tag $new_release `git rev-parse HEAD`"
git tag $new_release `git rev-parse HEAD`
echo "+ git push"
git push
gorelease
