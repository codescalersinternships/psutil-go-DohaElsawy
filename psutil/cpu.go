package psutil

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrFindEnoughFields = errors.New("cpuinfo: did't find all data")
)

type CpuInfo struct {
	Vendor    string
	MdoelName string
	CacheSize string
	CPUMHZ    float64
	whenQuit  int
}

type cpuFile struct {
	fileName string
}


type icpu interface {
	loadCpudata() (string, error)
}

var _ icpu = (*cpuFile)(nil)

func GetCpuInfo() (CpuInfo, error) {
	file := newCpuFile("")

	data, err := file.loadCpudata()

	if err != nil {
		return CpuInfo{},nil
	}

	cpu := newCpu()

	err = cpu.parseCpuData(data)

	if err != nil {
		return CpuInfo{} , err
	}
	return cpu, nil
}

func newCpuFile(path string) cpuFile {
	return cpuFile{
		fileName: path,
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

func (fcpu *cpuFile) loadCpudata() (string, error) {

	data, err := os.ReadFile(fcpu.fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *CpuInfo) parseCpuData(data string) error {

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.SplitN(line, ":",2)
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
