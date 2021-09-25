// Code generated by github.com/abhimanyu003/sttr/cmd/generate.go. DO NOT EDIT
// generated: Fri, 24 Sep 2021 22:31:53 +0000

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/abhimanyu003/sttr/processors"
	"github.com/spf13/cobra"
)

var zeropad_flag_n int
var zeropad_flag_p string

func init() {	
	zeropadCmd.Flags().IntVarP(&zeropad_flag_n, "number-of-zeros", "n", 5, "Number of zeros to be padded")
	zeropadCmd.Flags().StringVarP(&zeropad_flag_p, "prefix", "p", "", "The number get prefixed with this")
	rootCmd.AddCommand(zeropadCmd)
}

var zeropadCmd = &cobra.Command{
	Use:   "zeropad",
	Short: "Pad a number with zeros",
	Aliases: []string {},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		in, out := "", ""

		if len(args) == 0 {
			all, err := ioutil.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			in = string(all)
		} else {
			in = args[0]
		}

		p := processors.Zeropad{}
		flags := make([]processors.Flag, 0)
		flags = append(flags, processors.Flag{Short: "n", Value: zeropad_flag_n})
		flags = append(flags, processors.Flag{Short: "p", Value: zeropad_flag_p})

		out, err = p.Transform(in, flags...)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, out)
		return err
	},
}
