# go-info

go-info is a Go module for storing build and version information about the binary in the executable. It provides a basic flag support, when using Go's `flag` package to output this data and provides a simple API to get access to the data in case needed in the application.

## Use

### Building

Using go-info requires slightly modifying your build command to include some additional ldflags. 

```sh
go build -ldflags "-X github.com/jrozner/go-info.version=1.0.0 -X github.com/jrozner/go-info.commitHash=$(git rev-list -1 HEAD) -X github.com/jrozner/go-info.buildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
```

### Code Examples

To use go-info simply import the package into your main package or wherever you intend to use it's data and make sure to use `flag.Parse()` if you are not already using the flag package in your code.

```go
import "github.com/jrozner/go-info"

flag.Parse()
```

This will automatically register the "-version" flag when using the `flag` module. You can check for the presence of it in your application and print the info to stdout.

```go
if info.Flag() {
    info.Print()
}
```

## License

This code is licensed under the [MIT](LICENSE) license