# go-cgi-reverse-proxy

## Usage

- Build and install `go get github.com/Rompei/go-cgi-reverse-proxy && cd $GOPATH/src/github.com/Rompei/go-cgi-reverse-proxy && go install`
- Create [config file](https://github.com/Rompei/go-cgi-reverse-proxy/tree/master/config/config.yaml)
- Execute command `go-cgi-reverse-proxy -c /path/to/config/file -r /path/to/root/directory`
- Directories and executables are generated.

```
root/
├── backend
│   ├── caption
│   │   ├── aaa
│   │   │   └── index.cgi
│   │   ├── bbb
│   │   │   └── index.cgi
│   │   └── index.cgi
│   ├── deepdream
│   │   └── index.cgi
│   ├── index.cgi
│   └── styletransfer
│       └── index.cgi
├── frontend
│   ├── caption
│   │   └── index.cgi
│   ├── deepdream
│   │   └── index.cgi
│   ├── index.cgi
│   └── styletransfer
│       └── index.cgi
└── index.cgi
```

## License

[BSD 3](https://opensource.org/licenses/BSD-3-Clause)
