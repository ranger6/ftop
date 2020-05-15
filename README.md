# ftop

Convert a fully qualified domain name to a filesystem path, reversing the order.

## example

```
$ ftop muir.remulac.fr
fr/remulac/muir
$
```

## building

You will need a [Go](https://golang.org/) tool suite installed. Go modules are used; but are hardly necessary for such a simple tool.  So, a version of Go >= 1.11 should be used.  However, the tool should build just fine with an older version.

There are a variety of ways to build a `Go` command.

You can build the command directly from this directory:
```
$ go build
```

You can even fetch it from the source code repository, build, and install it in `$GOPATH/bin` in one step. From any directory:
```
$ go get github.com/ranger6/ftop
```
