package main

import (
	"fmt"
	"os"

	"github.com/riandyrn/go-env"
)

const (
	envValString  = "STRING"
	envValInt     = "INT"
	envValBool    = "BOOL"
	envValSeconds = "SECONDS"
	envValStrings = "STRINGS"
)

func main() {
	// set env variables
	os.Setenv(envValString, "abc")
	os.Setenv(envValInt, "123")
	os.Setenv(envValBool, "true")
	os.Setenv(envValSeconds, "5")
	os.Setenv(envValStrings, "item1,item2,item3")

	// get string value
	strVal := env.GetString(envValString)
	fmt.Printf("strVal: %v\n", strVal) // should output `abc`

	// get int value
	intVal := env.GetInt(envValInt)
	fmt.Printf("intVal: %v\n", intVal) // should output `1`

	// get bool value
	boolVal := env.GetBool(envValBool)
	fmt.Printf("boolVal: %v\n", boolVal) // should output `true`

	// get seconds value
	secVal := env.GetSeconds(envValSeconds)
	fmt.Printf("secVal: %v\n", secVal) // should output `5s`

	// get strings value
	strsVal := env.GetStrings(envValStrings, ",")
	fmt.Printf("strsVal: %v\n", strsVal) // should output `[item1 item2 item3]`
}
