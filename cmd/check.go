package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check site is available",
	Long:  `Checks site is available via proxy`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				go func() {
					time.Sleep(5 * time.Second)
					fmt.Println(time.Now())
				}()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
