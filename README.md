# yaml2go-cli

![build_badge](https://github.com/Icemap/yaml2go-cli/workflows/Go/badge.svg)

a command tool for [yaml2go](https://github.com/PrasadG193/yaml2go)

let it can use `go install` command

## Install

```bash
go install github.com/Icemap/yaml2go-cli@latest
```

## Show Help

```bash
yaml2go --help
yaml2go converts YAML specs to Go type definitions

Usage:
    yaml2go < /path/to/yamlspec.yaml

Examples:
    yaml2go < test/example1.yaml
    yaml2go < test/example1.yaml > example1.go
```