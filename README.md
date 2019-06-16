# Go Env

Fetching primitive values from env without hassle in Go!

## Why

Have you been annoyed when your code become ugly because you need to fetch value from env but you also need to take care of its error when what you need is just its value?

Especially if the value you need to fetch is requiring some parsing first (e.g `int`).

```go
...
strVal, _ := os.Getenv("IntKey")
intVal, _ := strconv.Atoi(strVal) // too much hassle!
...
```

This module is offering you solution to that!

```go
intVal := env.GetInt("IntKey") // so simple!
```

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

No need to care that troublesome errors!

All errors in the process would simply make the functions return their returned type zero value (e.g `env.GetInt("abc")` would returns `0`).

Check `env_test.go` to see possible scenarios handled by this module.

For full working example, check out `/example` directory.

---
