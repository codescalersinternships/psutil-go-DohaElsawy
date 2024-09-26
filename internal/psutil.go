package internal

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	ErrNotFindEnoughFields = errors.New("cpuinfo: did't find all data")
)

type CpuInfo struct {
	vendor    string
	mdoelName string
	cacheSize string
	cPUMHZ    float64
	whenQuit  int
}

var (

	numOfAttribute = 4
)

func loadCpudata(fileName string) (string, error) {
	fcpu, err := os.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(fcpu), nil
}



func (c *CpuInfo) ParseCpuData(fileName string) error {
	
	c.whenQuit = numOfAttribute

	data, err := loadCpudata(fileName)

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

		if c.whenQuit == 0 {
			return nil
		}

	}

	return ErrNotFindEnoughFields

}
