package webhelpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListerServerParser(t *testing.T) {
	t.Log("Test - generate server listner string")

	testCases := []struct {
		host   string
		port   int
		expect string
	}{
		{
			host:   "localhost",
			port:   80,
			expect: "localhost:80",
		},
		{
			host:   "",
			port:   8081,
			expect: ":8081",
		},
	}

	for _, test := range testCases {
		actualValue := GetListenServerString(test.host, test.port)
		assert.Equal(t, actualValue, test.expect)
	}
}
