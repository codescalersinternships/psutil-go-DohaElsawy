package psutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Process struct {
	PID  int
	Name string
}

var (
	ErrNoNameForProcess = errors.New("impossible no name for the process")
)

func ListProc() ([]Process, error) {

	fpid, err := openProc()

	if err != nil {
		return nil, err
	}

	listProc, err := parseProcDir(fpid)

	if err != nil {
		return nil, err
	}

	return listProc, nil
}

func openProc() ([]string, error) {
	dirProc, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer dirProc.Close()

	pids, err := dirProc.Readdirnames(0)

	if err != nil {
		return nil, err
	}
	log.Printf("%v", pids)

	return pids, nil
}

func parseProcDir(pids []string) ([]Process, error) {
	var processList []Process
	var proc Process

	for _, pid := range pids {
		log.Printf("%v", pid)
		id, err := strconv.Atoi(pid)
		if err != nil {
			continue
		}

		name, err := getProcName(id)

		if err != nil {
			return nil, err
		}

		proc.PID = id
		proc.Name = name

		processList = append(processList, proc)
	}
	return processList, nil
}

func getProcDetails(pid int) (string, error) {
	path := fmt.Sprintf("/proc/%d/status", pid)

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func getProcName(pid int) (string, error) {
	data, err := getProcDetails(pid)

	if err != nil {
		return "", err
	}

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.SplitN(line, ":", 2)

		key := keyVal[0]
		value := keyVal[1]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "Name" {
			return value, nil
		}
	}
	return "", ErrNoNameForProcess
}
