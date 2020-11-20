# Go search example

## Requirements

either 
- Docker
- VSCode and [Remote Container Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)

or 
- [Go 1.15](https://golang.org/dl/)

## Get started

### VSCode

Open project, enable "Container Extension" and press `F5`.

### Shell

```
go mod download
go build -o ./main ./cmd/main.go
./main
```

## Attribution

> Copyright (c) 2019 vanGato (https://github.com/phpSoftware/search-engine-template) \
Distributed under https://github.com/phpSoftware/search-engine-template/blob/master/LICENSE

> Rob Pike - Google I/O 2012 - Go Concurrency Patterns \
https://www.youtube.com/watch?v=f6kdp27TYZs