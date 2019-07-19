package resource

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestNewNamespaceCmd_NonEmptyName(t *testing.T) {
	rootCmd := &cobra.Command{Annotations: make(map[string]string)}

	nsc := NewNamespaceCmd(rootCmd, "foo")

	assert.Equal(t, "foo", nsc.Name)
	assert.True(t, rootCmd.HasSubCommands())
	assert.Equal(t, "namespace", rootCmd.Annotations["foo"])
}

func TestNewNamespaceCmd_EmptyName(t *testing.T) {
	rootCmd := &cobra.Command{Annotations: make(map[string]string)}

	nsc := NewNamespaceCmd(rootCmd, "")

	assert.Equal(t, "", nsc.Name)
	assert.False(t, rootCmd.HasSubCommands())
	_, ok := rootCmd.Annotations[""]
	assert.False(t, ok)
}
