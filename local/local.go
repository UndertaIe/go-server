package local

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunLocal(cmd *cobra.Command, args []string) {
	fmt.Println("this is local")
}
