# ghrelease

Golang GitHub release tool. Uses the execellent
[mitchellh/gox](https://github.com/mitchellh/gox) tool.

## Usage

```
> ghrelease
Usage of ghrelease:
  -a, --authfile="": Set the auth file for GitHub credentials
  -c, --create=false: Create release if it does not exist
  -b, --buildflags="": Pass build flags to gox
  -f, --force=false: Replace the asset if it already exists
  -n, --name: Set the name for the release asset
  -p, --prerelease=false: With --create, create as a dev release
  -r, --release="": Set the release name
  -v, --version: Print the name and version
```
