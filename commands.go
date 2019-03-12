package jenkins

import (
	"github.com/spf13/cobra"
)

func init() {

	jbRegister("getjobs", &cobra.Command{
		Use:          "get [flags]",
		Short:        "command to get jobs",
		Long:         `This will help user to get list of jobs present in jenkins.`,
		Run:          cc.getJob,
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
	})
	jbRegister("createjobs", &cobra.Command{
		Use:          "create [flags]",
		Short:        "command to create jobs",
		Long:         `This will help user to create jobs in jenkins.`,
		Run:          cc.createJob,
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
	})
	jbRegister("deletejob", &cobra.Command{
		Use:          "delete [flags]",
		Short:        "command to delete jobs",
		Long:         `This will help user in deletion of existing jobs present in jenkins.`,
		Run:          cc.deleteJob,
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
	})
	jbRegister("copyjob", &cobra.Command{
		Use:          "copy [flags]",
		Short:        "command to copy jobs",
		Long:         `This will help user in copying one jenkins job to another, this will come in handy while creating a new jenkins job which is identical to existing.`,
		Run:          cc.copyJob,
		Args:         cobra.MinimumNArgs(2),
		SilenceUsage: true,
	})
	jbRegister("buildjob", &cobra.Command{
		Use:          "build [flags]",
		Short:        "command to build jobs",
		Long:         `This will help user in triggering jenkins build.`,
		Run:          cc.buildJob,
		Args:         cobra.MinimumNArgs(1),
		SilenceUsage: true,
	})
}