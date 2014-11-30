package main

import (
	flag "github.com/docker/docker/pkg/mflag"
)

var (
	AuthFile    = flag.String([]string{"a", "-authfile"}, "", "Set the auth file for GitHub credentials")
	CreateAsset = flag.Bool([]string{"c", "-create"}, false, "Create release if it does not exist")
	BuildFlags  = flag.String([]string{"b", "-buildflags"}, "", "Pass build flags to gox")
	ForceCreate = flag.Bool([]string{"f", "-force"}, false, "Replace the asset if it already exists")
	AssetName   = flag.String([]string{"n", "-name"}, "", "Set the name for the release asset")
	Prerelease  = flag.Bool([]string{"p", "-prerelease"}, false, "With -create, create as a dev release")
	ReleaseName = flag.String([]string{"r", "-release"}, "", "Set the release name")
	Version     = flag.Bool([]string{"v", "-version"}, false, "Print the name and version")
)

func main() {
}
