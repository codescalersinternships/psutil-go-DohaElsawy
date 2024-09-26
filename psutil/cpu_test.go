package psutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuInfo(t *testing.T) {

	t.Run("get actual data", func(t *testing.T) {

		c, err := GetCpuInfo()

		assert.NoError(t, err)
		assert.NotEmpty(t, c)
		assert.NotEmpty(t, c.Vendor)
		assert.NotEmpty(t, c.MdoelName)
		assert.NotEmpty(t, c.CacheSize)
		assert.NotEmpty(t, c.CPUMHZ)

	})

}



type fakeCpuFile struct {
	fileName string
}

func newfakeCpuFile() fakeCpuFile {
	return fakeCpuFile{
		fileName: "./testdata/fakecpufile.txt",
	}
}

func (fcpu *fakeCpuFile) loadCpudata() (string, error) {

	data, err := os.ReadFile(fcpu.fileName)

	if err != nil {
		return "", err
	}

	return string(data), nil
}


func TestLoadData(t *testing.T) {

	file := newfakeCpuFile()

	data, err := file.loadCpudata()

	assert.NoError(t, err)

	assert.NotEmpty(t, data)

}



