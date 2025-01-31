package debug

import (
	"github.com/urfave/cli"

	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
	command.Subcommands = append(command.Subcommands, cli.Command{
		Name:    name,
		Aliases: aliases,
		Usage:   "Debugging and troubleshooting commands",
		Subcommands: []cli.Command{

			cli.Command{
				Name:      "export-validators",
				Aliases:   []string{"x"},
				Usage:     "Exports a TSV file of validators",
				UsageText: "rocketpool api debug export-validators",
				Action: func(c *cli.Context) error {

					// Validate args
					if err := cliutils.ValidateArgCount(c, 0); err != nil {
						return err
					}

					// Export TSV of validators
					ExportValidators(c)
					return nil

				},
			},
		},
	})
}
