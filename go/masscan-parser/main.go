package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	var filePath, prefix string
	flag.StringVar(&filePath, "f", "", "example.json")
	flag.StringVar(&prefix, "p", "", "lines prefix, for example: http://")

	flag.Parse()

	if filePath == "" {
		flag.PrintDefaults()
		log.Fatal("IP not set")
	}

	rawPortsFileBytes, _ := ioutil.ReadFile(filePath)
	jsonPortsString := `{ "data":` + string(rawPortsFileBytes) + `}`
	result := gjson.Get(jsonPortsString, "data")
	portsMap := make(map[string][]int)

	for _, name := range result.Array() {
		jsonLine := name.String()
		ip := gjson.Get(jsonLine, "ip").String()
		port := gjson.Get(jsonLine, "ports.0.port").String()
		portNumber, _ := strconv.Atoi(port)
		portsMap[ip] = append(portsMap[ip], portNumber)
	}

	for ip, ports := range portsMap {
		for _, port := range ports {
			fmt.Println(prefix + ip + ":" + strconv.Itoa(port))
		}
	}
}
