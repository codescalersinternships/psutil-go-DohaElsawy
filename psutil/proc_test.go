package psutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListProc(t *testing.T) {
	t.Run("return valid proc list form priavte list", func(t *testing.T) {

		fakeproc := newfakeProcFile("./testdata")

		listproc, err := listProc(&fakeproc)

		assert.NoError(t, err)
		assert.NotEmpty(t, listproc)

		for _, proc := range listproc {

			assert.NotEmpty(t, proc.Name)
			assert.NotEmpty(t, proc.Name)
		}
	})
	t.Run("return error proc list form priavte list", func(t *testing.T) {

		fakeproc := newfakeProcFile("./unknown")

		listproc, err := listProc(&fakeproc)

		assert.Error(t, err)
		assert.Empty(t, listproc)
	})
}

func TestGetDetails(t *testing.T) {
	t.Run("valid details", func(t *testing.T) {
		path := "psutil/testdata/1/status"
		data, err := getProcData(path)

		assert.Error(t, err)
		assert.Empty(t, data)
	})
}

type fakeProcFile struct {
	fileName string
}

func newfakeProcFile(path string) fakeProcFile {
	return fakeProcFile{
		fileName: path,
	}
}

func (procf fakeProcFile) openProc() ([]string, error) {
	dirProc, err := os.Open(procf.fileName)
	if err != nil {
		return nil, err
	}
	defer dirProc.Close()

	pids, err := dirProc.Readdirnames(0)

	if err != nil {
		return nil, err
	}

	return pids, nil
}
