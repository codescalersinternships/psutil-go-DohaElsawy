package psutil

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Process represent process structure
type Process struct {
	PID  int
	Name string
}

type iProcFile interface {
	openProc() ([]string, error)
	// newProcFile(path string)
}

var _ iProcFile = (*procFile)(nil)

type procFile struct {
	fileName string
}

// ListProc return processes list and error if found
func ListProc() ([]Process, error) {

	procf := newProcFile("/proc")

	return listProc(&procf)
}

// GetProcDetails returns details about procsess and error if found
func GetProcDetails(pid int) (string, error) {

	path := fmt.Sprintf("/proc/%d/status", pid)
	return getProcData(path)
}

func getProcData(path string) (string, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func listProc(procf iProcFile) ([]Process, error) {

	fpid, err := procf.openProc()

	if err != nil {
		return nil, err
	}

	listProc, err := parseProcDir(fpid)

	if err != nil {
		return nil, err
	}

	return listProc, nil

}

func newProcFile(path string) procFile {
	return procFile{
		fileName: path,
	}
}

func (procf *procFile) openProc() ([]string, error) {
	dirProc, err := os.Open(procf.fileName)
	if err != nil {
		return nil, err
	}
	defer dirProc.Close()

	pids, err := dirProc.Readdirnames(0)

	if err != nil {
		return nil, err
	}

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

		path := fmt.Sprintf("/proc/%d/status", id)

		data, err := getProcData(path)

		if err != nil {
			return nil, err
		}

		name := getProcName(data)

		proc.PID = id
		proc.Name = name

		processList = append(processList, proc)
	}
	return processList, nil
}

func getProcName(data string) string {

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.SplitN(line, ":", 2)

		key := keyVal[0]
		value := keyVal[1]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "Name" {
			return value
		}
	}
	return ""
}
