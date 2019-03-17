# Logger
## Description
Logger interface so that your application code does not depend on the implementation. Currently it is supporting logger wrap for:
1. [uber-zap](https://github.com/uber-go/zap)
2. [logrus](https://github.com/sirupsen/logrus)

## How to use
To use in your code you can import directly `import github.com/amitrai48/logger` or download the code

```
$ go get github.com/amitrai48/logger
```

## Examples
Check example from [folder](https://github.com/amitrai48/logger/examples/example.go) `examples/example.go`
The example show how the logger package can use logger from `logrus` or `zap`
```
$ go run examples/example.go
```

## Contributing
We encourage and support an active, healthy community of contributors — including you! Details are in the [contribution guide](https://github.com/amitrai48/logger/blob/master/CONTRIBUTING.md)
