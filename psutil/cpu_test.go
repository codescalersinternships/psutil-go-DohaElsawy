package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuInfo(t *testing.T) {

	t.Run("get actual data", func(t *testing.T) {

		err := GetCpuInfo()

		assert.NoError(t, err)
	})

}
