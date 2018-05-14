package eureka

import (
	"encoding/json"
	"go-router/eureka/model"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"strings"
)

// InstanceInfo holds information about one instance of one app
type InstanceInfo struct {
	Hostname   string
	Port       int
	InstanceID string
	Secure     bool
}

// NextAvailableInstance finds the next instance to be used
func NextAvailableInstance(appName string) (*InstanceInfo, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8761/eureka/apps/"+appName, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var jsonResponse model.EurekaResponse
	err = json.NewDecoder(res.Body).Decode(&jsonResponse)
	if err != nil {
		panic(err)
	}
	println(jsonResponse.Application.Name)
	target := jsonResponse.Application.Instance[rand.Int()%len(jsonResponse.Application.Instance)]
	return &InstanceInfo{Hostname: target.HostName, Port: target.Port.Value, InstanceID: target.InstanceID}, nil
}

// Director reverse proxy director for eureka
func Director(req *http.Request) {
	var head string
	head, req.URL.Path = shiftPath(req.URL.Path)
	info, _ := NextAvailableInstance(head)
	/*if err != nil {
		http.Error(res, "Not Found", http.StatusNotFound)
	}*/
	req.URL.Host = info.Hostname + ":" + strconv.Itoa(info.Port)
	if info.Secure {
		req.URL.Scheme = "https"
	} else {
		req.URL.Scheme = "http"
	}
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
