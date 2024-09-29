package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProc(t *testing.T) {

	t.Run("return valid proc list", func(t *testing.T) {
		listproc , err := ListProc()

		assert.NoError(t, err)

		assert.NotEmpty(t, listproc)
	})

}
