package psutil

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Generic file system errors.
var (
	ErrFindEnoughFields = errors.New("info: did't find all data")
)

// MemInfo represent memory info structure
type CpuInfo struct {
	Vendor    string
	MdoelName string
	CacheSize string
	CPUMHZ    float64
	whenQuit  int
}

type cpuFile struct {
	readData func(path string) (string, error)
}

// GetCpuInfo return CPU info struncture and error if found
func GetCpuInfo() (CpuInfo, error) {

	defaultPath := "/proc/cpuinfo"

	return getCpuInfo(defaultPath)
}

func getCpuInfo(path string) (CpuInfo, error) {

	cpuf := newCpuFile()

	data, err := cpuf.readData(path)

	if err != nil {
		return CpuInfo{}, nil
	}

	cpu := newCpu()

	err = cpu.parseCpuData(data)

	if err != nil {
		return CpuInfo{}, err
	}
	return cpu, nil

}

func loadData(path string) (string, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func newCpuFile() cpuFile {
	return cpuFile{
		readData: loadData,
	}
}

func newCpu() CpuInfo {
	return CpuInfo{
		Vendor:    "",
		MdoelName: "",
		CPUMHZ:    0,
		CacheSize: "",
		whenQuit:  4,
	}
}

func (c *CpuInfo) parseCpuData(data string) error {

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.SplitN(line, ":", 2)
		key := keyVal[0]
		value := keyVal[1]

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if key == "vendor_id" {
			c.Vendor = value
			c.whenQuit--
		}
		if key == "cache size" {
			c.CacheSize = value
			c.whenQuit--
		}
		if key == "model name" {
			c.MdoelName = value
			c.whenQuit--
		}
		if key == "cpu MHz" {
			convVal, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			c.CPUMHZ = convVal
			c.whenQuit--
		}

		if c.whenQuit == 0 {
			return nil
		}

	}

	return ErrFindEnoughFields

}
