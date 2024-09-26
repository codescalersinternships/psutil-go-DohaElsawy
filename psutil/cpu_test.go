package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuInfo(t *testing.T) {

	t.Run("get actual data", func(t *testing.T) {

		c , err := GetCpuInfo()

		assert.NoError(t, err)
		assert.NotEmpty(t, c)
	})

}


func TestCpu(t *testing.T) {

	t.Run("valid case", func(t *testing.T) {

	})
}

func TestLoaddata(t *testing.T) {
	file := newCpuFile("./testdata/fakecpufile.txt")

	data, err := file.loadCpudata()

	assert.NoError(t, err)

	assert.NotEmpty(t, data)

}
