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
	})
	t.Run("return error proc list form priavte list", func(t *testing.T) {

		fakeproc := newfakeProcFile("./unknown")

		listproc, err := listProc(&fakeproc)

		assert.Error(t, err)
		assert.Empty(t, listproc)
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
