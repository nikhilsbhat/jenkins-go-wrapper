package jenkins

import (
	"fmt"
	jen "github.com/bndr/gojenkins"
	//log "github.com/nikhilsbhat/neuron/logger"
	err "github.com/nikhilsbhat/neuron/error"
	//"os"
	"github.com/nikhilsbhat/neuron/app/config"
	"reflect"
)

type jenkinsCred struct {
	JenkinsDomain   string `json:"jenkinsdomain,omitempty"`
	JenkinsUsername string `json:"jenkinsusername,omitempty"`
	JenkinsPassword string `json:"jenkinspassword,omitempty"`
}

func readCiConfig(pathtofile string) (jenkinsCred, error) {
	confile, confneuerr := ioutil.ReadFile(pathtofile)
	if confneuerr != nil {
		log.Error(err.InvalidConfig())
		return jenkinsCred{}, confneuerr
	}

	var confdata jenkinsCred
	if decodneuerr := config.JsonDecode([]byte(confile), &confdata); decodneuerr != nil {
		fmt.Println(err.JsonDecodeError())
		return jenkinsCred{}, decodneuerr
	}
	return confdata, nil
}

func (auth *jenkinsCred) isCfgEmpty() bool {

	if reflect.DeepEqual(auth, jenkinsCred{}) {
		return true
	} else {
		if (auth.JenkinsDomain != "") || (auth.JenkinsUsername != "") || (auth.JenkinsPassword != "") {
			return true
		}
	}
	return false
}

func (auth *jenkinsCred) AuthJenkins() (*jen.Jenkins, error) {

	if res := auth.isCfgEmpty(); res != true {
		jenkins := jen.CreateJenkins(nil, auth.JenkinsDomain, auth.JenkinsUsername, auth.JenkinsPassword)
		_, jenerr := jenkins.Init()
		if jenerr != nil {
			return nil, jenerr
		}
		return jenkins, nil
	}
	return nil, fmt.Errorf("It seems the jenkins config you provided is empty")
}