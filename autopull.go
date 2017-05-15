package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/BurntSushi/toml"
)

var (
	conf = make(map[string]string)
)

func init() {
	toml.DecodeFile("/etc/autopull/conf.toml", &conf)
}

func main() {
	fmt.Println("Listening on port:8920")
	fmt.Println("Thr trigger url is: /trigger-git?project=xxx")
	http.HandleFunc("/trigger-git", trigger)
	http.ListenAndServe(":8920", nil)
}

func trigger(w http.ResponseWriter, r *http.Request) {
	fmt.Println("======== New request =======")

	r.ParseForm()
	form := new(bytes.Buffer)
	e1 := toml.NewEncoder(form)
	e1.Encode(r.Form)
	fmt.Println("Form Data:")
	fmt.Println(form.String())

	header := new(bytes.Buffer)
	e2 := toml.NewEncoder(header)
	e2.Encode(r.Header)
	fmt.Println("Header:")
	fmt.Println(header.String())

	projectName := r.FormValue("project")
	if projectUrl, find := conf[projectName]; find {
		cmd := exec.Command("go", "get", "-u", projectUrl)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()
	} else {
		fmt.Println("project:", projectName, "not found")
	}

	fmt.Println("\n\n")

	w.Write([]byte("hello world"))
}
