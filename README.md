# heroku-golang-echoServer

```bash
gvm use go1.6
gvm pkgset use vs-code

go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/tpng/gopkgs
go get -u -v github.com/newhook/go-symbols
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v github.com/cweill/gotests/...
go get -u -v golang.org/x/tools/cmd/godoc

go get -u github.com/kardianos/govendor

go build main.go

PORT=8080 heroku local

```

```bash 

# use govendor

# link project to ~/.gvm/pkgsets/__use__/__ps_name/src/__name__

cd ~/.gvm/pkgsets/__use__/__ps_name/src/__name__

govendor init
govendor add github.com/ant0ine/go-json-rest/rest
govendor fetch github.com/ant0ine/go-json-rest/rest

govendor test +local
govendor install +local


```
