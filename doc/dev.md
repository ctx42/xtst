```shell
VERSION=$(cat ../VER) || exit 1
GOPROXY=proxy.golang.org go list -m github.com/ctx42/xtst@${VERSION}
```
