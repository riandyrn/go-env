package main

import (
	"fmt"
	"os"

	"github.com/riandyrn/go-env"
)

const (
	envValString = "STRING"
	envValInt    = "INT"
	envValBool   = "BOOL"
)

func main() {
	// set env variables
	os.Setenv(envValString, "abc")
	os.Setenv(envValInt, "123")
	os.Setenv(envValBool, "true")

	// get string value
	strVal := env.GetString(envValString)
	fmt.Printf("strVal: %v\n", strVal) // should output `abc`

	// get int value
	intVal := env.GetInt(envValInt)
	fmt.Printf("intVal: %v\n", intVal) // should output `1`

	// get bool value
	boolVal := env.GetBool(envValBool)
	fmt.Printf("boolVal: %v\n", boolVal) // should output `true`
}
