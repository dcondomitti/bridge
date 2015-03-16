package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

func extractData(n *etcd.Node) (string, string) {
	key := n.Key
	value := n.Value

	tokens := strings.Split(key, "/")
	return tokens[len(tokens)-1], value
}

func parseCommand(cmd []string) (string, []string) {
	return cmd[0], cmd[1:]
}

func setEnvVars(etcdHost string, path string, debug bool) {
	client := etcd.NewClient([]string{etcdHost})
	resp, err := client.Get(path, false, false)

	if err != nil {
		log.Fatal(err)
	}

	for _, n := range resp.Node.Nodes {
		key, value := extractData(n)
		if debug {
			log.Printf("Exporting %s as %s\n", key, value)
		} else {
			log.Printf("Exporting %s\n", key)
		}

		os.Setenv(key, value)
	}

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

	setEnvVars(etcdHost, path, debug)

	log.Printf("Executing %s %s", binary, strings.Join(params, " "))
	cmd := exec.Command(binary, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
