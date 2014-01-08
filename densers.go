// The denser server. It's a simple little web server that
// accepts IP address reports from the denser client every
// so often. It provides an API endpoint for retrieving
// that IP address.
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os/user"
	"path/filepath"
)

type Config struct {
	Port int
}

func ConfigPath() string {
	user, _ := user.Current()
	return filepath.Join(user.HomeDir, ".denser")
}

func ReadConfig() (config Config) {
	data, _ := ioutil.ReadFile(ConfigPath())
	toml.Decode(string(data), &config)
	return
}

var Ip = "IP address has not been set yet. Wait a few minutes."

func GetIp(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(Ip))
}

func SetIp(w http.ResponseWriter, req *http.Request) {
	Ip = mux.Vars(req)["ip"]
	fmt.Println("Setting IP address to", Ip)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	config := ReadConfig()
	r := mux.NewRouter()
	r.HandleFunc("/get", GetIp).Methods("GET")
	r.HandleFunc("/set/{ip}", SetIp).Methods("PUT")
	http.Handle("/", r)

	var port int
	if config.Port != 0 {
		port = config.Port
	} else {
		port = 3245
	}

	fmt.Printf("Listening on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
