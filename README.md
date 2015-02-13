# ghrelease

Golang GitHub release tool.

## Usage

```
> ghrelease
Usage of ./ghrelease:
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
    "auth_token": "supersekritgithubapptoken"
}
```

`auth_token` can also be provided by the `GHRELEASE_AUTH_TOKEN` environment
variable.
