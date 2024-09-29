package psutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCpuInfo(t *testing.T) {

	t.Run("get actual data", func(t *testing.T) {

		c, err := getCpuInfo("./testdata/fakecpufile.txt")

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

func TestParseData(t *testing.T) {
	testcase := []struct {
		description string
		expected    CpuInfo
		input       string
		err         error
	}{
		{
			description: "valid case",
			expected: CpuInfo{
				Vendor:    "GenuineIntel",
				MdoelName: "Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz",
				CacheSize: "8192 KB",
				CPUMHZ:    1100.752,
				whenQuit:  0,
			},
			input: `processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 126
model name	: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
stepping	: 5
microcode	: 0xc6
cpu MHz		: 1100.752
cache size	: 8192 KB
physical id	: 0
siblings	: 8
core id		: 0
cpu cores	: 4
apicid		: 0`,
			err: nil,
		},
		{
			description: "only found one field",
			expected: CpuInfo{
				Vendor:    "GenuineIntel",
				MdoelName: "",
				CacheSize: "",
				CPUMHZ:    0,
				whenQuit:  3,
			},
			input: `processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 126`,
			err: ErrFindEnoughFields,
		},
		{
			description: "not found any field",
			expected: CpuInfo{
				Vendor:    "",
				MdoelName: "",
				CacheSize: "",
				CPUMHZ:    0,
				whenQuit:  4,
			},
			input: `processor	: 0
cpu family	: 6
model		: 126
stepping	: 5
microcode	: 0xc6
physical id	: 0
siblings	: 8
core id		: 0
cpu cores	: 4
apicid		: 0`,
			err: ErrFindEnoughFields,
		},
	}

	for _, test := range testcase {
		t.Run(test.description, func(t *testing.T) {
			c := newCpu()

			err := c.parseCpuData(test.input)

			assert.Equal(t,test.err, err)
			assert.Equal(t,c,test.expected)
		})
	}

}
