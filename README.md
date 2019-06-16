# Go Env

Fetching primitive values from env without hassle in Go!

## Installation

```bash
go get github.com/riandyrn/go-env
```

## Code Usage

```go
...

import "github.com/riandyrn/go-env"

...
strVal := env.GetString("StrKey") // fetch string value
intVal := env.GetInt("IntKey") // fetch int value
boolVal := env.GetBool("BoolKey") // fetch bool value
...
```

No need to care about error! All errors would be simply converted to its zero value.

For full working example, check `/example` directory.

---
