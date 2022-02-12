# Godbolt CLI

Godbolt console wrapper for easily execute local file without any security risk and compiler.

## Install
- Compile the source code and add to your PATH.
OR
- `go install gitlab.com/aiocat/godbolt-cli@latest`

## Usage

- `godbolt-cli file.extension`
- `godbolt-cli test.c`
- `godbolt-cli https://raw.githubusercontent.com/dfellis/llvm-hello-world/master/helloWorld.ll`
- `godbolt-cli https://raw.githubusercontent.com/dfellis/llvm-hello-world/master/helloWorld.ll output.txt`

## Supported Languages

- [Languages Map](https://gitlab.com/aiocat/godbolt-cli/-/blob/main/compiler.go#L14)


## License

This project is distributed under MIT license. for more information:

- https://gitlab.com/aiocat/godbolt-cli/-/blob/main/LICENSE
