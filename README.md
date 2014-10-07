# bolt-ircbot

```bolt-ircbot``` responds to irc msg as like as ```#1234``` and replies github issue information.

Based on work by [GitHub user soh335](https://github.com/soh335/github-issue-ircbot)

## How to install

```
go get github.com/GawainLynch/go-ircevent
cd $GOPATH/src/github.com/GawainLynch/go-ircevent/
git checkout 40cfe292a9577a79503e08c90c00987919499cd9
go get github.com/GawainLynch/bolt-ircbot
go install github.com/GawainLynch/bolt-ircbot
```

## Usage

```
bolt-ircbot --config /path/to/config.json
```

## See also

* http://blog.handlena.me/entry/2013/06/12/234712
* http://soh335.hatenablog.com/entry/2013/06/13/103457
