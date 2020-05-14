# ftop

Convert a fully qualified domain name to a filesystem path.

## example

```
% ftop muir.remulac.fr
fr/remulac/muir
%
```

## building

You will need a [Go](https://golang.org/) tool suite installed.  There are a variety of ways to build a `Go` command.

You can build the command directly from this directory:
```
% go build ftop.go
```

You can even fetch it from the source code repository, build, and install it in one step.
```
% go get github.com/ranger6/ftop
```
