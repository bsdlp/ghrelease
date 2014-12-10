package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"

	"code.google.com/p/goauth2/oauth"

	flag "github.com/docker/docker/pkg/mflag"
	"github.com/google/go-github/github"
)

type Config struct {
	BuildFlags []string `json:"build_flags,omitempty"`
	AuthToken  string   `json:"auth_token"`
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
	GoxPath     string
)

func LoadConfig(ConfigPath *string) Config {
	var ret Config
	ConfigFile, err := ioutil.ReadFile(*ConfigPath)
	err = json.Unmarshal(ConfigFile, &ret)
	if err != nil {
		log.Fatalln(err)
	}
	return ret
}

func init() {
	CurrentUser, _ := user.Current()
	DefaultConfigPath := path.Join(CurrentUser.HomeDir, ".config/ghrelease/config.json")
	ConfigPath = flag.String([]string{"c", "-config"}, DefaultConfigPath, "Set ghrelease config file")
	GoxPath = flag.String([]string{"g", "-gox"}, FindGox(), "Path to gox")

	flag.Parse()
}

func FindGox() string {
	path, err = exec.LookPath("gox")
	if err != nil {
		log.Fatalln("gox is not in your $PATH")
	}
	return path
}

func Gox(BuildFlags *[]string) *string {
	cmd := exec.Command(GoxPath, *BuildFlags)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	return *out.String()
}

func main() {
	GHRConfig := LoadConfig(ConfigPath)

	EnvAuthToken := os.Getenv("GHRELEASE_AUTH_TOKEN")
	if EnvAuthToken != "" {
		GHRConfig.AuthToken = EnvAuthToken
	}
	if GHRConfig.AuthToken == "" {
		log.Println("auth_token is required to upload release assets to GitHub")
	}

	transport := &oauth.Transport{
		Token: &oauth.Token{AccessToken: GHRConfig.AuthToken},
	}
	client := github.NewClient(transport.Client())

	BuildOutput := Gox(BuildFlags)
}
