package psutil

import (
	"os"
	"strconv"
	"strings"
)


// MemInfo represent memory info structure
type MemInfo struct {
	Total     string
	Used      string
	Available string
	whenQuit  int
}

type iMemFile interface {
	loadData() (string, error)
}

var _ iMemFile = (*memFile)(nil)

type memFile struct {
	fileName string
}


// GetMemInfo return memory info struncture and error if found
func GetMemInfo() (MemInfo, error) {
	memf := newMemFile()

	data, err := memf.loadData()

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

func (m *memFile) loadData() (string, error) {

	data, err := os.ReadFile(m.fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func newMemFile() memFile {
	return memFile{
		fileName: "/proc/meminfo",
	}
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
