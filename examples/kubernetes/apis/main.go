package main

import (
	"encoding/json"
	// "net"
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	apiServerHost = "10.103.113.127"
	port          = "8080"
)

func main() {
	// 获取k8s集群的版本
	getK8sVersion()
	// 获取k8s api 路由
	getPath()

	deployDeployment(DeploymentNginx)
}

func getPath() {
	log.Info("Get k8s api path")
	addr := "http://" + apiServerHost + ":" + port
	resp, err := http.Get(addr)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	var res PathRes
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Error(err)
	}
	log.Info(res.Paths[0])
	defer resp.Body.Close()
}

func getK8sVersion() {
	log.Info("Get k8s version...")
	versionAddr := "http://" + apiServerHost + ":" + port + "/version"
	resp, err := http.Get(versionAddr)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(bytes.NewBuffer(body).String())
}

func deployDeployment(data string) {
	// deployment a deployment
	log.Info("Post info to create deploy")
	yaml_value, _ := yaml.Marshal([]byte(data))
	resp, err := http.Post(
		"http://10.103.113.127:8001/apis/apps/v1/namespaces/development/deployments",
		"application/yaml",
		bytes.NewBuffer(yaml_value),
	)
	if err != nil{
		log.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(bytes.NewBuffer(body).String())
	
	defer resp.Body.Close()
}
