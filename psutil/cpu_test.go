package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpu(t *testing.T) {

	t.Run("valid case", func(t *testing.T) {

		expect := &Cpu{
			Vendor:    "GenuineIntel",
			MdoelName: "Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz",
			CacheSize: "8192 KB",
			CPUMHZ:    3500,
			whenQuit:  0,
			fileName: "/proc/cpuinfo",
		}

		c := InitCPU()

		err := c.CpuGetInfo()

		assert.Equal(t, expect, c)
		assert.NoError(t, err)

	})
}

func TestLoaddata(t *testing.T) {
	scanner, err := loadCpudata("./testdata/fakecpufile.txt")

	assert.NoError(t, err)

	assert.NotEmpty(t, scanner)

}
