package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemInfo(t *testing.T) {

	t.Run("get actual data from private get function", func(t *testing.T) {
		m, err := getMemInfo("./testdata/fakememfile.txt")

		assert.NoError(t, err)
		assert.NotEmpty(t, m)
		assert.NotEmpty(t, m.Available)
		assert.NotEmpty(t, m.Used)
		assert.NotEmpty(t, m.Total)

	})

	t.Run("get actual data from public get function", func(t *testing.T) {
		m, err := GetMemInfo()

		assert.NoError(t, err)
		assert.NotEmpty(t, m)
		assert.NotEmpty(t, m.Available)
		assert.NotEmpty(t, m.Used)
		assert.NotEmpty(t, m.Total)

	})

}

func TestLoadMemData(t *testing.T) {

	file := newMemFile("./testdata/fakememfile.txt")

	data, err := file.loadData()

	assert.NoError(t, err)
	assert.NotEmpty(t, data)

}

func TestParseMemData(t *testing.T) {
	testcase := []struct {
		description string
		expected    MemInfo
		input       string
		err         error
	}{
		{
			description: "valid case",
			expected: MemInfo{
				Total:    "12010608 kB",
				Used: "9718220 kB",
				Available: "4893396 kB",
				whenQuit:  0,
			},
			input: `MemTotal:       12010608 kB
MemFree:         2292388 kB
MemAvailable:    4893396 kB
Buffers:          170816 kB
Cached:          3854016 kB
SwapCached:       177684 kB`,
			err: nil,
		},
		{
			description: "only found one field",
			expected: MemInfo{
				Total:    "12010608 kB",
				Used: "",
				Available: "",
				whenQuit:  2,
			},
			input: `MemTotal:       12010608 kB`,
			err: ErrFindEnoughFields,
		},
		{
			description: "not found any field",
			expected: MemInfo{
				Total:    "",
				Used: "",
				Available: "",
				whenQuit:  3,
			},
			input: `SReclaimable:     488288 kB
SUnreclaim:       380972 kB
KernelStack:       34736 kB
PageTables:        89376 kB
SecPageTables:         0 kB
NFS_Unstable:          0 kB`,
			err: ErrFindEnoughFields,
		},
	}

	for _, test := range testcase {
		t.Run(test.description, func(t *testing.T) {
			m := newMem()

			err := m.parseMemData(test.input)

			assert.Equal(t,test.err, err)
			assert.Equal(t,test.expected,m)
		})
	}

}