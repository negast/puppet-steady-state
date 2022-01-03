package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PuppetResource struct {
	Tags        []string      `json:"tags"`
	File        string        `json:"file"`
	Type        string        `json:"type"`
	Title       string        `json:"title"`
	Line        int           `json:"line"`
	Resource    string        `json:"resource"`
	Environment string        `json:"environment"`
	Certname    string        `json:"certname"`
	Parameters  []interface{} `json:"parameters,omitempty"`
	Exported    bool          `json:"exported"`
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func main() {

	//	we'll first get all resources to determine the steady state
	hostname := "ubuntu1"
	base_url := "http://localhost:8080"
	services, _ := nodeSteadyState(hostname, base_url)

	for _, s := range services {
		println(s)
	}
	//for _, f := range files {
	//	println(f)
	//}
}

func nodeSteadyState(hostname string, base_url string) ([]string, []string) {
	services := []string{}
	files := []string{}
	resources := getResourcesForNode(hostname, base_url)
	for _, r := range resources {
		if r.Type == "File" {
			files = append(files, r.Title)
		}
		if r.Type == "Service" {
			services = append(services, r.Title)
		}
	}
	return services, files
}

func getResourcesForNode(node string, base_url string) []PuppetResource {
	url := base_url + "/pdb/query/v4/resources"

	payload := strings.NewReader(fmt.Sprintf("{\"query\":[\"=\",\"certname\",\"%s\"]}", node))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	data := []PuppetResource{}
	json.Unmarshal(body, &data)

	return data

}
