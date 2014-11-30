package main

import flag "github.com/docker/docker/pkg/mflag"

// initialize flags
var (
	Config      = flag.String([]string{"c", "-config"}, "~/.config/ghrelease/config.json", "Set ghrelease config file")
	BuildFlags  = flag.String([]string{"b", "-buildflags"}, "", "Pass build flags to gox")
	ForceCreate = flag.Bool([]string{"f", "-force"}, false, "Replace the asset if it already exists")
	AssetName   = flag.String([]string{"n", "-name"}, "", "Set the name for the release asset")
	Prerelease  = flag.Bool([]string{"p", "-prerelease"}, false, "With -create, create as a dev release")
	ReleaseName = flag.String([]string{"r", "-release"}, "", "Create the release if it does not exist and set the release name")
	Version     = flag.Bool([]string{"v", "-version"}, false, "Print the name and version")
)

func init() {
	flag.Parse()
}

func main() {
}
