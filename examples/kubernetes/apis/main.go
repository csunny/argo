package main

import (
	"encoding/json"
	// "net"
	"net/http"
	"io/ioutil"
	"bytes"
	log "github.com/sirupsen/logrus"
)

const (
	apiServerHost = "10.103.113.127"
	port = "8080"
)

func main()  {
	// 获取k8s集群的版本
	getK8sVersion()
	// 获取k8s api 路由
	getPath()
}

func getPath(){
	log.Info("Get k8s api path")
	addr := "http://" + apiServerHost + ":" + port
	resp, err := http.Get(addr)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Error(err)
	}
	var res PathRes
	err = json.Unmarshal(body, &res)
	if err != nil{
		log.Error(err)
	}
	log.Info(res.Paths)
	defer resp.Body.Close()
}	

func getK8sVersion()  {
	log.Info("Get k8s version...")
	versionAddr := "http://" + apiServerHost + ":" + port + "/version"
	resp, err := http.Get(versionAddr)
	if err != nil{
		log.Error(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Error(err)
	}
	log.Info(bytes.NewBuffer(body).String())
}
