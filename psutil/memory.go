package psutil

import (
	"os"
	"strconv"
	"strings"
)

type MemInfo struct {
	Total     string
	Used      string
	Available string
	whenQuit  int
}

type memFile struct {
	readData func(path string) (string, error)
}

func GetMemInfo() (MemInfo, error) {

	memf := newCpuFile()

	data, err := memf.readData("/proc/meminfo")

	if err != nil {
		return MemInfo{}, nil
	}

	mem := newMem()

	err = mem.parseMemData(data)

	if err != nil {
		return MemInfo{}, err
	}
	return mem, nil
}

func newMemFile() memFile {
	return memFile{
		readData: loadMemdata,
	}
}

func loadMemdata(path string) (string, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func newMem() MemInfo {
	return MemInfo{
		Total:     "",
		Used:      "",
		Available: "",
		whenQuit:  3,
	}
}

func (m *MemInfo) parseMemData(data string) error {

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.SplitN(line, ":", 2)
		key := keyVal[0]
		value := keyVal[1]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "MemTotal" {
			m.Total = value
			m.whenQuit--
		}
		if key == "MemAvailable" {
			m.Available = value
			m.whenQuit--
		}
		if key == "MemFree" {
			totalmem := strings.Trim(m.Total, " kB")
			freeMem := strings.Trim(value, " kB")

			total, err := strconv.Atoi(totalmem)
			if err != nil {
				return err
			}

			free, err := strconv.Atoi(freeMem)
			if err != nil {
				return err
			}

			m.Used = strconv.Itoa(total - free)
			m.Used += " kB"

			m.whenQuit--
		}
		if m.whenQuit == 0 {
			return nil
		}

	}

	return ErrFindEnoughFields

}
