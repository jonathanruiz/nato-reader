package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

// test function
func TestNormalOutput(t *testing.T) {

	cmd := &cobra.Command{Use: "word"}
	cmd.SetArgs([]string{"--word", "Jonathan"})
	cmd.Execute()

}


