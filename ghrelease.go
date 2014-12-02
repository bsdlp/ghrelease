package main

import (
	"fmt"
	"os/user"

	flag "github.com/docker/docker/pkg/mflag"
)

type Config struct {
	BuildFlags string `json:"build_flags,omitempty"`
	AuthToken  string `json:"auth_token"`
}

// initialize flags
var (
	// Config needs to be initted in init()
	ConfigPath  *string
	BuildFlags  = flag.String([]string{"b", "-buildflags"}, "", "Pass build flags to gox")
	ForceCreate = flag.Bool([]string{"f", "-force"}, false, "Replace the asset if it already exists")
	AssetName   = flag.String([]string{"n", "-name"}, "", "Set the name for the release asset")
	Prerelease  = flag.Bool([]string{"p", "-prerelease"}, false, "With -create, create as a dev release")
	ReleaseName = flag.String([]string{"r", "-release"}, "", "Create the release if it does not exist and set the release name")
	Version     = flag.Bool([]string{"v", "-version"}, false, "Print the name and version")
)

func init() {
	CurrentUser, _ := user.Current()
	DefaultConfigPath := fmt.Sprintf("%s/.config/ghrelease/config.json", CurrentUser.HomeDir)
	ConfigPath = flag.String([]string{"c", "-config"}, DefaultConfigPath, "Set ghrelease config file")

	flag.Parse()
}

func main() {
}
