package test

import (
	"fmt"
	"testing"

	"github.com/levigross/grequests"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	defer resp.Close()
	fmt.Println(resp.String())
	fmt.Println(resp.RawResponse)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
}
