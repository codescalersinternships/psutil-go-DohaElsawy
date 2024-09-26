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
