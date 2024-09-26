package psutil

import (
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

func TestLoadData(t *testing.T) {

	file := newCpuFile()

	data, err := file.readData("./testdata/fakecpufile.txt")

	assert.NoError(t, err)

	assert.NotEmpty(t, data)

}
