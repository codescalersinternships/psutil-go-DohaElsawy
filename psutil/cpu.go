package psutil

import (
	"errors"

	"github.com/codescalersinternships/psutil-go-DohaElsawy/internal"
)

var (
	ErrNotFindEnoughFields = errors.New("cpuinfo: did't find all data")
)

func GetCpuInfo() error {

	var c internal.CpuInfo
	err := c.ParseCpuData("/proc/cpuinfo")

	if err != nil {
		return err
	}

	return nil

}
