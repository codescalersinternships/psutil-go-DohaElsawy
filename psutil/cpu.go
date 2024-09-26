package psutil

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotFindEnoughFields = errors.New("cpuinfo: did't find all data")
)

type ICpu interface {
	InitCPU()
	CpuGetInfo()
}

type Cpu struct {
	vendor    string
	mdoelName string
	cacheSize string
	cPUMHZ    float64
	whenQuit  int
	fileName  string
}

func InitCPU() *Cpu {
	return &Cpu{
		vendor:    "",
		mdoelName: "",
		cacheSize: "",
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

func (c *Cpu) GetCpuInfo() error {

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
			c.vendor = value
			c.whenQuit--
		}
		if key == "cache size" {
			c.cacheSize = value
			c.whenQuit--
		}
		if key == "model name" {
			c.mdoelName = value
			c.whenQuit--
		}
		if key == "cpu MHz" {
			convVal, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			c.cPUMHZ = convVal
			c.whenQuit--
		}

		if c.whenQuit < 1 {
			return nil
		}

	}

	return ErrNotFindEnoughFields
}

func (c *Cpu) String() string {

	return fmt.Sprintf("cpd vecder is: %s\ncpu model name is: %s\ncache size is: %s\ncpu MHz is: %f\n", c.vendor, c.mdoelName, c.cacheSize, c.cPUMHZ)

}
