package env

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetString(t *testing.T) {
	testCases := []struct {
		Name     string
		EnvKey   string
		EnvValue string
		ExpValue string
	}{
		{
			Name:     "Test Key Not Set",
			EnvKey:   "",
			EnvValue: "ABC",
			ExpValue: "",
		},
		{
			Name:     "Test Value Empty",
			EnvKey:   "Key",
			EnvValue: "",
			ExpValue: "",
		},
		{
			Name:     "Test Value Non Empty",
			EnvKey:   "Key",
			EnvValue: "ABC",
			ExpValue: "ABC",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if len(testCase.EnvKey) != 0 {
				os.Setenv(testCase.EnvKey, testCase.EnvValue)
			}
			val := GetString(testCase.EnvKey)
			if val != testCase.ExpValue {
				t.Fatalf("unexpected val, expected: %v, got: %v", testCase.EnvValue, val)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	testCases := []struct {
		Name     string
		EnvKey   string
		EnvValue string
		ExpValue int
	}{
		{
			Name:     "Test Key Not Set",
			EnvKey:   "",
			EnvValue: "123",
			ExpValue: 0,
		},
		{
			Name:     "Test Value Empty",
			EnvKey:   "Key",
			EnvValue: "",
			ExpValue: 0,
		},
		{
			Name:     "Test Value Not Integer (String)",
			EnvKey:   "Key",
			EnvValue: "ABC",
			ExpValue: 0,
		},
		{
			Name:     "Test Value Not Integer (Float)",
			EnvKey:   "Key",
			EnvValue: "123.132",
			ExpValue: 0,
		},
		{
			Name:     "Test Value Integer",
			EnvKey:   "Key",
			EnvValue: "123",
			ExpValue: 123,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if len(testCase.EnvKey) != 0 {
				os.Setenv(testCase.EnvKey, testCase.EnvValue)
			}
			val := GetInt(testCase.EnvKey)
			if val != testCase.ExpValue {
				t.Fatalf("unexpected val, expected: %v, got: %v", testCase.ExpValue, val)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	testCases := []struct {
		Name     string
		EnvKey   string
		EnvValue string
		ExpValue bool
	}{
		{
			Name:     "Test Key Not Set",
			EnvKey:   "",
			EnvValue: "true",
			ExpValue: false,
		},
		{
			Name:     "Test Value Not Set",
			EnvKey:   "Key",
			EnvValue: "",
			ExpValue: false,
		},
		{
			Name:     "Test Invalid Value",
			EnvKey:   "Key",
			EnvValue: "ABC",
			ExpValue: false,
		},
		{
			Name:     `Test Valid Value ("true")`,
			EnvKey:   "Key",
			EnvValue: "true",
			ExpValue: true,
		},
		{
			Name:     `Test Valid Value ("TRUE")`,
			EnvKey:   "Key",
			EnvValue: "TRUE",
			ExpValue: true,
		},
		{
			Name:     `Test Valid Value ("1")`,
			EnvKey:   "Key",
			EnvValue: "1",
			ExpValue: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if len(testCase.EnvKey) != 0 {
				os.Setenv(testCase.EnvKey, testCase.EnvValue)
			}
			val := GetBool(testCase.EnvKey)
			if val != testCase.ExpValue {
				t.Fatalf("unexpected val, expected: %v, got: %v", testCase.ExpValue, val)
			}
		})
	}
}

func TestGetSeconds(t *testing.T) {
	testCases := []struct {
		Name     string
		EnvKey   string
		EnvValue string
		ExpValue time.Duration
	}{
		{
			Name:     "Test Key Not Set",
			EnvKey:   "",
			EnvValue: "123",
			ExpValue: 0 * time.Second,
		},
		{
			Name:     "Test Value Empty",
			EnvKey:   "Key",
			EnvValue: "",
			ExpValue: 0 * time.Second,
		},
		{
			Name:     "Test Value Not Integer (String)",
			EnvKey:   "Key",
			EnvValue: "ABC",
			ExpValue: 0 * time.Second,
		},
		{
			Name:     "Test Value Not Integer (Float)",
			EnvKey:   "Key",
			EnvValue: "123.132",
			ExpValue: 0 * time.Second,
		},
		{
			Name:     "Test Value Integer",
			EnvKey:   "Key",
			EnvValue: "60",
			ExpValue: 60 * time.Second,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if len(testCase.EnvKey) != 0 {
				os.Setenv(testCase.EnvKey, testCase.EnvValue)
			}
			val := GetSeconds(testCase.EnvKey)
			if val != testCase.ExpValue {
				t.Fatalf("unexpected val, expected: %v, got: %v", testCase.ExpValue, val)
			}
		})
	}
}

func TestGetStrings(t *testing.T) {
	testCases := []struct {
		Name      string
		EnvKey    string
		EnvValue  string
		Separator string
		ExpValue  []string
		ExpLen    int
	}{
		{
			Name:      "Test Key Not Set",
			EnvKey:    "",
			EnvValue:  "ABC",
			Separator: ",",
			ExpValue:  nil,
			ExpLen:    0,
		},
		{
			Name:      "Test Value Empty",
			EnvKey:    "Key",
			EnvValue:  "",
			Separator: ",",
			ExpValue:  nil,
			ExpLen:    0,
		},
		{
			Name:      "Test Value Not Empty - Single Item",
			EnvKey:    "Key",
			EnvValue:  "ABC",
			Separator: ",",
			ExpValue:  []string{"ABC"},
			ExpLen:    1,
		},
		{
			Name:      "Test Value Not Empty - Mulitple Items",
			EnvKey:    "Key",
			EnvValue:  "ABC,DEF,GHI",
			Separator: ",",
			ExpValue:  []string{"ABC", "DEF", "GHI"},
			ExpLen:    3,
		},
		{
			Name:      "Test Value Not Empty - Mulitple Items With Space",
			EnvKey:    "Key",
			EnvValue:  "ABC, DEF, G HI",
			Separator: ", ",
			ExpValue:  []string{"ABC", "DEF", "G HI"},
			ExpLen:    3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			if len(testCase.EnvKey) != 0 {
				os.Setenv(testCase.EnvKey, testCase.EnvValue)
			}
			val := GetStrings(testCase.EnvKey, testCase.Separator)
			if len(val) != testCase.ExpLen {
				t.Fatalf("unexpected length, expected: %v, got: %v", testCase.ExpLen, len(val))
			}
			if strings.Join(val, testCase.Separator) != strings.Join(testCase.ExpValue, testCase.Separator) {
				t.Fatalf("unexpected val, expected: %v, got: %v", testCase.EnvValue, val)
			}
		})
	}
}
