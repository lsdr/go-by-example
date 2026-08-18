# Go by Example

Repos with code and other by products of reading [Go by Example][goex]

## Up & Running

### IDE

Install [GoLand][goland] by Jetbrains and have fun

### Vim / NeoVim

LazyVim has a pretty sweet [out-of-the-box setup for Go][lazygo]

### building binaries to `dist` dir

```sh
go build -o dist/hello-world 01-hello-world/01-hello-world.go
```

## Troubleshooting

### `go: cannot determine module path for source directory`

Newer versions of go require a module to be set.

#### Solution

Add a `go.mod` to the project:

```sh
cd PROJECT_ROOT
touch go.mod
echo 'github.com/lsdr/go-by-example' > go.mod
```

## Links and References

### See Also

- [Code Organization section][gocorg], How to Write Go Code
- [Effective Go][efgo], Go Programming Language official docs
- [Getting started with VS Code Go][vsog]

### `vgo` and build problems

- <https://stackoverflow.com/questions/50803978/golang-how-to-use-vgo-error-cannot-find-package>
- <https://github.com/golang/go/issues/25176>
- <https://github.com/golang/go/wiki/Modules>

:school_satchel:

[goex]: https://gobyexample.com/
[goland]: https://www.jetbrains.com/go/
[efgo]: https://golang.org/doc/effective_go.html
[gocorg]: https://go.dev/doc/code#Organization
[vsog]: https://code.visualstudio.com/docs/languages/go
[lazygo]: https://www.lazyvim.org/extras/lang/go
