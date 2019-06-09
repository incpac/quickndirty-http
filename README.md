# QuickNDirty HTTP server

A quick'n'dirty HTTP server designed to temporarily serve static files.

## Compile

Simply download, grab dependencies and compile.

```
go get github.com/incpac/quickndirty-http
dep ensure
go build -ldflags "-X main.Version=0.0.1"
```

## Running

```
$ ./qndhttp --content ./myawesomecontent --bind 0.0.0.0:4000
```

See `qndhttp --help` for all options