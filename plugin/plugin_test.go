package plugin

import (
	"os"
	"testing"

	"github.com/vincenzopalazzo/cln4go/client"
)

func TestCallFistMethod(t *testing.T) {
	path := os.Getenv("CLN_UNIX_SOCKET")
	client, err := client.NewUnix(path)
	if err != nil {
		panic(err)
	}
	response, err := client.Call("hello", make(map[string]interface{}))
	if err != nil {
		panic(err)
	}

	message, found := response["message"]
	if !found {
		t.Error("The message is not found")
	}

	if message != "hello from go 1.18" {
		t.Errorf("message received %s different from expected %s", message, "hello from go 1.18")
	}
}

func TestOptionValueExist(t *testing.T) {
	path := os.Getenv("CLN_UNIX_SOCKET")
	client, err := client.NewUnix(path)
	if err != nil {
		panic(err)
	}
	response, err := client.Call("foo_bar", make(map[string]interface{}))
	if err != nil {
		panic(err)
	}

	message, found := response["message"]
	if !found {
		t.Error("The message is not found")
	}

	if message != "Hello Go" {
		t.Errorf("message received %s different from expected %s", message, "Hello Go")
	}
}
