package jenkins

import (
	"fmt"
	err "github.com/nikhilsbhat/neuron/error"
	"github.com/spf13/cobra"
)

var (
	cmds map[string]*cobra.Command
)

type jbcmds struct {
	commands []*cobra.Command
}

// Register will help in registering all the subcommands to the main commands.
func Register(name string, fn *cobra.Command) {
	if cmds == nil {
		cmds = make(map[string]*cobra.Command)
	}

	if cmds[name] != nil {
		panic(fmt.Errorf("Command %q is already registered", name))
	}
	cmds[name] = fn
}

func getCmds() *cobra.Command {
	jenkinscmd := new(jbcmds)
	jenkinscmd.commands = append(jenkinscmd.commands, getJobCmds())
	//future subcommands will go here

	// This gets the full and final command with all subcommands and flags for jenkins
	cmd := jenkinscmd.prepareCmds()
	return cmd
}

func (c *jbcmds) prepareCmds() *cobra.Command {
	rootCmd := getJenkinsCmds()
	for _, cm := range c.commands {
		rootCmd.AddCommand(cm)
	}
	return rootCmd
}

// SetNeuronCmds helps in gathering all the subcommands so that it can be used while registering it with main command.
func SetJenkinsCmds() *cobra.Command {
	cmd := getCmds()
	return cmd
}

func getJenkinsCmds() *cobra.Command {

	var jenkinsCmd = &cobra.Command{
		Use:   "jenkins [command]",
		Short: "command to deal with 'go' wrapper of jenkins and its activities",
		Long:  `This will help user to use go wrapper of jenkins and to get things done in various aspects including creating/deleting/updating/copying job,nodes and etc of jenkins etc.`,
		RunE:  cc.echoJenkins,
	}
	jenkinsCmd.SetUsageTemplate(getUsageTemplate())
	registerFlags(jenkinsCmd)
	return jenkinsCmd
}

func (cm *cliMeta) echoJenkins(cmd *cobra.Command, args []string) error {
	if cm.CliSet == false {
		return err.CliNoStart()
	}
	cmd.Usage()
	return nil
}

// This function will return the custom template for usage function,
// only functions/methods inside this package can call this.

func getUsageTemplate() string {
	return `{{printf "\n"}}Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if gt (len .Aliases) 0}}{{printf "\n" }}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}{{printf "\n" }}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{printf "\n"}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}{{printf "\n"}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}{{printf "\n"}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}{{printf "\n"}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}{{printf "\n"}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}"
{{printf "\n"}}`
}