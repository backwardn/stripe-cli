package resource

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewOperationCmd(t *testing.T) {
	parentCmd := &cobra.Command{Annotations: make(map[string]string)}

	oc := NewOperationCmd(parentCmd, "foo", "/v1/bars/{id}", "get")

	assert.Equal(t, "foo", oc.Name)
	assert.Equal(t, "/v1/bars/{id}", oc.Path)
	assert.Equal(t, "GET", oc.HTTPVerb)
	assert.Equal(t, []string{"{id}"}, oc.URLParams)
	assert.True(t, parentCmd.HasSubCommands())
	assert.Equal(t, "operation", parentCmd.Annotations["foo"])
}
