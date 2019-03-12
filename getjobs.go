package jenkins

import (
	"encoding/json"
	"fmt"
	"fmt"
	jen "github.com/bndr/gojenkins"
	//log "github.com/nikhilsbhat/neuron/logger"
	//err "github.com/nikhilsbhat/neuron/error"
	//"os"
	"reflect"
	"github.com/spf13/cobra"
	"strings"
)

var (
	jobcmds map[string]*cobra.Command
)

// The function that helps in registering the subcommands with the respective main command.
// Make sure you call this, and this is the only way to register the subcommands.
func jbRegister(name string, fn *cobra.Command) {
	if jobcmds == nil {
		jobcmds = make(map[string]*cobra.Command)
	}

	if jobcmds[name] != nil {
		panic(fmt.Sprintf("Command %s is already registered", name))
	}
	jobcmds[name] = fn
}

// The only way to create server command is to call this function and
// package commands will take care of calling this.
func getJobCmds() *cobra.Command {

	// Creating "server" happens here.
	var cmdJob = &cobra.Command{
		Use:   "job [flags]",
		Short: "command to carry out ci activities",
		Long:  `This will help user to create/update/get/delete/build jobs in the CI they ask for.`,
		Run:   cc.echoJob,
	}
	registerjbFlags("server", cmdJob)

	for cmdname, cmd := range jobcmds {
		cmdJob.AddCommand(cmd)
		registerjbFlags(cmdname, cmd)
	}
	return cmdJob
}

// Registering all the flags to the subcommands and command job itself.
func registerjbFlags(cmdname string, cmd *cobra.Command) {

	switch strings.ToLower(cmdname) {
	}
}

func (cm *cliMeta) getJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}
	createsv.Cloud.Name = cm.getCloud(cmd)
	createsv.Cloud.Region = cm.getRegion(cmd)
	createsv.Cloud.Profile = cm.getProfile(cmd)
	createsv.Cloud.GetRaw = cm.getGetRaw(cmd)
	servresponse, servresperr := createsv.CreateServer()
	if servresperr != nil {
		cm.NeuronSaysItsError(servresperr.Error())
	} else {
		jsonval, _ := json.MarshalIndent(servresponse, "", " ")
		cm.NeuronSaysItsInfo(string(jsonval))
	}
}

func (cm *cliMeta) createJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}
	deletesv.Cloud.Name = cm.getCloud(cmd)
	deletesv.Cloud.Region = cm.getRegion(cmd)
	deletesv.Cloud.Profile = cm.getProfile(cmd)
	deletesv.Cloud.GetRaw = cm.getGetRaw(cmd)
	deletesvresponse, sverr := deletesv.DeleteServer()
	if sverr != nil {
		cm.NeuronSaysItsError(sverr.Error())
	} else {
		jsonval, _ := json.MarshalIndent(deletesvresponse, "", " ")
		cm.NeuronSaysItsInfo(string(jsonval))
	}
}

func (cm *cliMeta) deleteJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}

	getsv.Cloud.Name = cm.getCloud(cmd)
	getsv.Cloud.Region = cm.getRegion(cmd)
	getsv.Cloud.Profile = cm.getProfile(cmd)
	getsv.Cloud.GetRaw = cm.getGetRaw(cmd)

	if cm.isAll(cmd) {
		getservresponse, svgeterr := getsv.GetAllServers()
		if svgeterr != nil {
			cm.NeuronSaysItsError(svgeterr.Error())
		} else {
			jsonval, _ := json.MarshalIndent(getservresponse, "", " ")
			cm.NeuronSaysItsInfo(string(jsonval))
		}
	}

	getservresponse, svgeterr := getsv.GetServersDetails()
	if svgeterr != nil {
		cm.NeuronSaysItsError(svgeterr.Error())
	} else {
		jsonval, _ := json.MarshalIndent(getservresponse, "", " ")
		cm.NeuronSaysItsInfo(string(jsonval))
	}
}

func (cm *cliMeta) copyJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}
	updatesv.Cloud.Name = cm.getCloud(cmd)
	updatesv.Cloud.Region = cm.getRegion(cmd)
	updatesv.Cloud.Profile = cm.getProfile(cmd)
	updatesv.Cloud.GetRaw = cm.getGetRaw(cmd)
	svupdateresponse, svuperr := updatesv.UpdateServers()
	if svuperr != nil {
		cm.NeuronSaysItsError(svuperr.Error())
	} else {
		jsonval, _ := json.MarshalIndent(svupdateresponse, "", " ")
		cm.NeuronSaysItsInfo(string(jsonval))
	}
}

func (cm *cliMeta) buildJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}
	cm.printMessage()
	cmd.Usage()
}

func (cm *cliMeta) echoJob(cmd *cobra.Command, args []string) {
	if cm.CliSet == false {
		cm.NeuronSaysItsError(err.CliNoStart().Error())
	}
	cm.printMessage()
	cmd.Usage()
}