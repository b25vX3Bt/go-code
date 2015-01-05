go-code
=======

golang example codes.

## go get

    $ go get github.com/b25vX3Bt/go-code

## my memo

### golang init

    $ export GOPATH=$HOME/work/golang      # add .bashrc
    $ export PATH=$GOPATH/bin:$PATH        # add .bashrc
    $ mkdir $GOPATH

### tool

    $ go get golang.org/x/tools/cmd/...
    $ go get github.com/nsf/gocode
    $ go get github.com/peco/peco/cmd/peco
    $ go get github.com/motemen/ghq
    $ vi ~/.gitconfig
    ...
    [ghq]
        root = ~/work/golang/src


### before git push ?

'go get' to git push

    $ git remote rename origin go_origin
    $ git remote add origin https://b25vX3Bt@github.com/b25vX3Bt/go-code.git

