package local

import (
	"log"

	"github.com/spf13/cobra"
)

func RunLocal(cmd *cobra.Command, args []string) {
	log.Println("this is local")
}
