// +build teste2e

package cmd

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -race -tags teste2e -coverprofile=coverage.txt -covermode=atomic -c github.com/helm/chart-testing/ct/cmd

func TestExecute(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	type testCase struct {
		name string
		args []string
		err  error
	}

	cases := []testCase{
		{
			"version",
			[]string{"ct", "version"},
			nil,
		},
		{
			"install",
			[]string{"ct", "install"},
			nil,
		},
		{
			"lint",
			[]string{"ct", "lint", "--chart-yaml-schema", "./etc/chart_schema.yaml", "--lint-conf", "./etc/lintconf.yaml"},
			nil,
		},
		{
			"lint without config",
			[]string{"ct", "lint"},
			errors.New("'chart_schema.yaml' neither specified nor found in default locations"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.args
			err := Execute()
			if err != nil || tc.err != nil {
				assert.Equal(t, tc.err.Error(), err.Error(), "Expected error strings to match")
			}
		})
	}
}
