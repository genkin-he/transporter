#!/usr/bin/env fish

# original version courtesy of https://github.com/oklog/oklog

function version_prompt
	echo "Semantic version: "
end

read --prompt version_prompt VERSION

if test (echo $VERSION | grep '^v')
	echo Use the raw semantic version, without a v prefix
	exit
end

set REV (git rev-parse --short HEAD)
echo Tagging $REV as v$VERSION
git tag --annotate v$VERSION -m "Release v$VERSION"
echo Be sure to: git push --tags
echo

set DISTDIR dist/v$VERSION
mkdir -p $DISTDIR

for pair in linux/amd64 darwin/amd64 dragonfly/amd64 freebsd/amd64 netbsd/amd64 openbsd/amd64 windows/amd64
	set GOOS   (echo $pair | cut -d'/' -f1)
	set GOARCH (echo $pair | cut -d'/' -f2)
	set BIN    $DISTDIR/transporter-$VERSION-$GOOS-$GOARCH
	echo $BIN
	env GOOS=$GOOS GOARCH=$GOARCH go build -o $BIN -ldflags="-X main.version=$VERSION" transporter/cmd/transporter
end
