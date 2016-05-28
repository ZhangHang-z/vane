Vane - A Package Manager In Golang.
---------------------------------------

vane is a package manager in golang that can be install some front-end packages, inspired from bower.

##Install
make sure you have golang and installed, and set the `$GOPATH` and `$GOBIN`.

```bash
$ go get github.com/ZhangHang-z/vane
$ sudo bash make.sh
```

if that way was failed, you can install manually.

```bash
$ go get github.com/ZhangHang-z/vane
$ go install $GOPATH/src/github.com/ZhangHang-z/vane/bin/vane-cli.go
$ mv $GOBIN/vane-cli /usr/bin/vane
```

##Usage
`vane install jquery`

##License
[The MIT License](./LICENSE)


##Author
[ZhangHang-z](https://github.com/ZhangHang-z)