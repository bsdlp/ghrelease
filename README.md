# ghrelease

Golang GitHub release tool. Uses the execellent
[mitchellh/gox](https://github.com/mitchellh/gox) tool.

## Usage

```
> ghrel
Usage of ghrel:
  -authfile="": Set the auth file for GitHub credentials
  -create=false: Create release if it does not exist
  -force=false: Replace the asset if it already exists
  -name: Set the name for the release asset
  -prerelease=false: With -create, create as a dev release
  -release="": Set the release name
  -version: Print the name and version
```
