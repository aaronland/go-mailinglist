# go-runtimevar-string

Work in progress.

This is a thin wrapper around the `gocloud.dev/runtimevar` libraries to work with AWS Secrets Manager values.

It does not implement the `runtimevar` interface (there is no `OpenString` method, for example) and assumes strings, mostly because I couldn't make the JSON decoder stuff work (and haven't spent the time to figure it out) and because I haven't read up on how to implement custom decoders.

This package exists to hide the details from other code and may eventually be retired assuming solutions to the outstanding issues mentioned above. Or not. Maybe `runtimevar.OpenString(ctx, *url)` is the "good enough is perfect" solution. We'll see...

## Example

```
package main

import (
	"context"
	"flag"
	"fmt"
	runtimevar "github.com/aaronland/gocloud-runtimevar-string"
	_ "gocloud.dev/runtimevar/blobvar"
	_ "gocloud.dev/runtimevar/constantvar"
	_ "gocloud.dev/runtimevar/filevar"
	"log"
)

func main() {

	url := flag.String("url", "", "...")

	flag.Parse()

	ctx := context.Background()

	s, err := runtimevar.OpenString(ctx, *url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
```

For example:

```
$> go run cmd/runtimevar/main.go -url 'constant://?val=hello+world&decoder=string'
hello world
```

Or this, since `decoder=string` is assumed:

```
$> go run cmd/runtimevar/main.go -url 'constant://?val=hello+world'
hello world
```

Or this, which is not actually part of the `gocloud.dev/runtimevar` package:

```
$> go run cmd/runtimevar/main.go -url 'awssecretsmanager://orthis/{SECRET_ID}?region={AWS_REGION}&credentials={AWS_CREDENTIALS}'
SECRET_DATA...
```

### awssecretsmanager:// URLs

_Please write me._

## See also

* https://gocloud.dev/howto/runtimevar/