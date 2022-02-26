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
./yaml2go-cli -h                                    
yaml2go-cli is a cli-tool for yaml to go struct

Usage:
  yaml2go-cli [flags]

Flags:
  -h, --help             help for yaml2go-cli
  -i, --input string     input yaml file path
  -o, --output string    output go file path
  -p, --package string   package name (default "main")
  -s, --struct string    struct name (default "Default")
```

## Example

```
./yaml2go-cli -i test/test.yaml -o test/test.bean.file
```

- [test/test.yaml](test/test.yaml)
- [test/test.bean.file](test/test.bean.file)
