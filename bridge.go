package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

func parseCommand(cmd []string) (string, []string) {
	return cmd[0], cmd[1:]
}

func setEnvVars(vars map[string]string, debug bool) {

	for key, value := range vars {
		if debug {
			log.Printf("Exporting %s as %s\n", key, value)
		} else {
			log.Printf("Exporting %s\n", key)
		}

		os.Setenv(key, value)
	}

}

func etcdClient(etcdHost string) *etcd.Client {
	return etcd.NewClient([]string{etcdHost})
}

func retrieveVars(c *etcd.Client, p string) (map[string]string, error) {
	resp, err := c.Get(p, false, false)

	if err != nil {
		return nil, err
	}

	vars := make(map[string]string)

	for _, n := range resp.Node.Nodes {
		key, value := path.Base(n.Key), n.Value
		vars[key] = value
	}

	return vars, nil
}

func executeCommand(binary string, params []string) {
	cmd := exec.Command(binary, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	var path string
	var etcdHost string
	var debug bool

	flag.StringVar(&path, "path", "/example.com/app_name", "Path to application variables")
	flag.StringVar(&etcdHost, "etcd_host", "http://127.0.0.1:4001", "etcd cluster endpoint")
	flag.BoolVar(&debug, "debug", false, "log environment variable values")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("Command not specified")
	}

	binary, params := parseCommand(args)

	log.Printf("Application path: %s, etcd endpoint %s", path, etcdHost)

	vars, err := retrieveVars(etcdClient(etcdHost), path)
	if err != nil {
		log.Println("etcd not available, skipping configuration and launching command...")
	} else {
		setEnvVars(vars, debug)
	}

	log.Printf("Executing %s %s", binary, strings.Join(params, " "))
	executeCommand(binary, params)
}
