package main

import (
	"bytes"
	"fmt"
	"github.com/juju/errors"
	"io/ioutil"
	"os"
	"os/exec"
)

func ReadDir(dir string) (map[string]string, error) {
	envs := map[string]string{}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return envs, errors.New("Could not read directory: " + dir)
	}

	for _, f := range files {
		file, err := ioutil.ReadFile(dir + "/" + f.Name())
		if err != nil {
			return envs, errors.New("Could not read from file: " + f.Name())
		}
		envs[f.Name()] = string(file)
	}
	return envs, nil
}

func RunCmd(command string, env map[string]string) int {
	for key, value := range env {
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("Could not set ENV variable %s with value %s \n", key, value)
		}
	}

	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(command)

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Println(exitError.ExitCode())
			return exitError.ExitCode()
		}
	}

	stdout := outbuf.String()
	fmt.Println(stdout)

	return 0
}

func main() {
	path := os.Args[1]
	command := os.Args[2]

	envs, err := ReadDir(path)
	if err != nil {
		fmt.Println("Could not get env-variables")
	}

	os.Exit(RunCmd(command, envs))
}
