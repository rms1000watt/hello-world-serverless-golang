package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHandler(t *testing.T) {
	out := Response{
		Message: "Okay so your other function also executed successfully!",
	}

	res, err := Handler()
	if err != nil {
		fmt.Println("Error: Handler call failed: ", err)
		t.Fatal(err)
	}

	assert.Equal(t, res.Message, out.Message, "They should be equal")
}
