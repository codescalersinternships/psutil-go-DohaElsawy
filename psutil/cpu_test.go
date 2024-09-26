package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpu(t *testing.T) {

	t.Run("valid case", func(t *testing.T) {

		expect := &Cpu{
			vendor:    "GenuineIntel",
			mdoelName: "Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz",
			cacheSize: "8192 KB",
			cPUMHZ:    1100.752,
			whenQuit:  0,
			fileName: "./testdata/fakecpufile.txt",
		}

		c := InitCPU()
		c.fileName = "./testdata/fakecpufile.txt"

		err := c.GetCpuInfo()

		assert.Equal(t, expect, c)
		assert.NoError(t, err)

	})
}

func TestLoaddata(t *testing.T) {
	data, err := loadCpudata("./testdata/fakecpufile.txt")

	assert.NoError(t, err)

	assert.NotEmpty(t, data)

}
