package simpleGinn

import (
	"testing"
)

func TestMyfunc(t *testing.T) {
	router := SimpleGinnFunc()
	router.Run("0.0.0.0:8023")
}

func ExampleMyfunc() {
	router := SimpleGinnFunc()
	router.Run("0.0.0.0:8023")
	// Output: 7
}