package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHandler(t *testing.T) {
	in := events.APIGatewayProxyRequest{
		Body: "hello world",
	}

	out := events.APIGatewayProxyResponse{
		Body: "Your POST body: " + in.Body,
	}

	res, err := Handler(in)
	if err != nil {
		fmt.Println("Error: Handler call failed: ", err)
		t.Fatal(err)
	}

	assert.Equal(t, res.Body, out.Body, "They should be equal")
}
