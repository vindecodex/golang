# Configuring Path ( Linux )

During Installation Golang docs recommends to put go binary files in the `/usr/local/go` which will make it as our `GOROOT`

Checking our `GOROOT`:

```php
go env GOROOT
```

Checking our `GOPATH`:

```php
go env GOPATH
```

Best advice is not to change `GOROOT`. But we can change `GOPATH` any where we like.

under `~/.profile`

`export GOPATH=/anylocation/` === `/anylocation/src/`

It is now easier to create modules under `/anylocation/src/yourproject` by just targeting your `projectname/modules`


#### Example:

```
|-anylocation/
	|-src/
		|-yourproject/
			|-yourmodules/
				|-test.go
			|-main.go
```

main.go
```go
package main

import "yourproject/yourmodules"

func main() {
	yourmodules.Method("test")
}
```

