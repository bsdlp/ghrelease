# ghrelease

Golang GitHub release tool. Uses the execellent
[mitchellh/gox](https://github.com/mitchellh/gox) tool.

## Usage

```
> ghrelease
Usage of ./ghrelease:
  -b, --buildflags=""                               Pass build flags to gox
  -c, --config="~/.config/ghrelease/config.json"    Set ghrelease config file
  -f, --force=false                                 Replace the asset if it already exists
  -n, --name=""                                     Set the name for the release asset
  -p, --prerelease=false                            With -create, create as a dev release
  -r, --release=""                                  Create the release if it does not exist and set the release name
  -v, --version=false                               Print the name and version
```
## Configuration

* `~/.config/ghrelease/config.json`:
```json
{
    "auth_token": "supersekritgithubapptoken",
    "build_flags": [
        "-parallel=30",
        "-verbose"
    ]
}
```

`auth_token` can also be provided by the `GHRELEASE_AUTH_TOKEN` environment
variable.
