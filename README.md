# comment

Linter `comment` reports commented code.

The linter makes best attempt to find commented code and by doing so may skip obvious commented code or report false
positives.

Check [invalid.go](./testdata/src/comment/invalid.go) for examples of possible reported comments.

## Install

```sh
go get github.com/mneverov/comment
```

## Run

```sh
comment ./...
```

## Note

It is impossible to differentiate between a valid comment and a commented code.

For example, to clarify a flow the following comment is valid and gives a context:

```
// this = that
```

But it also a valid go code.

The following comment is a URL for a website with some info that describes the following code as well as a valid go
label.

```
// https://bit.ly/2EMXD26
```
