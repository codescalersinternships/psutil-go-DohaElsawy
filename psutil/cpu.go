package psutil

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotFindEnoughFields = errors.New("cpuinfo: did't find all data")
)


type Cpu struct {
	Vendor    string
	MdoelName string
	CacheSize string
	CPUMHZ    float64
	whenQuit  int
	fileName  string
}

func InitCPU() *Cpu{
	return &Cpu{
		Vendor:    "",
		MdoelName: "",
		CacheSize: "",
		whenQuit:  4,
		fileName:  "/proc/cpuinfo",
	}
}



func loadCpudata(fileName string) (string, error) {
	fcpu, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(fcpu), nil
}

func (c *Cpu) CpuGetInfo() (error) {

	data, err := loadCpudata(c.fileName)

	if err != nil {
		return err
	}

	lines := strings.Split(data, "\n")

	for _, line := range lines {

		keyVal := strings.Split(line, ":")
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

	return ErrNotFindEnoughFields
}
