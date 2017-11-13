package msisdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	{
		m, err := Parse("0312341234", "MY")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "60312341234", m.GetMSISDN())
	}
	{
		m, err := Parse("6031231234", "MY")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "6031231234", m.GetMSISDN())
	}
	{
		m, err := Parse("88262323", "SG")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "6588262323", m.GetMSISDN())
	}
}

func Test_GetLocal(t *testing.T) {
	testMap := map[string]string{
		"+60312341234":    "0312341234",  // Malaysia
		"+6031231234":     "031231234",   // Malaysia
		"+605208 500":     "05208500",    // Malaysia
		"+605208 5000":    "052085000",   // Malaysia
		"+6081123123":     "081123123",   // Malaysia
		"+60111231234":    "0111231234",  // Malaysia
		"+601112341234":   "01112341234", // Malaysia
		"+6011-1234 1234": "01112341234", // Malaysia
		"+601112328670":   "01112328670", // Malaysia
		"+6531234567":     "31234567",    // Singapore
		"+6567895741":     "67895741",    // Singapore
		"+6588262323":     "88262323",    // Singapore
		"+6598268423":     "98268423",    // Singapore
	}
	for msisdnStr, localFormatted := range testMap {
		m, err := ParseMSISDN(msisdnStr)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, localFormatted, m.GetLocal())
	}
}

func Test_GetLocalFormatted(t *testing.T) {
	testMap := map[string]string{
		"+60312341234":    "03-1234 1234",  // Malaysia
		"+6031231234":     "03-123 1234",   // Malaysia
		"+605208 500":     "05-208 500",    // Malaysia
		"+605208 5000":    "05-208 5000",   // Malaysia
		"+6081123123":     "08-112 3123",   // Malaysia
		"+60111231234":    "011-123 1234",  // Malaysia
		"+601112341234":   "011-1234 1234", // Malaysia
		"+6011-1234 1234": "011-1234 1234", // Malaysia
		"+601112328670":   "011-1232 8670", // Malaysia
		"+6531234567":     "3123 4567",     // Singapore
		"+6567895741":     "6789 5741",     // Singapore
		"+6588262323":     "8826 2323",     // Singapore
		"+6598268423":     "9826 8423",     // Singapore
	}
	for msisdnStr, localFormatted := range testMap {
		m, err := ParseMSISDN(msisdnStr)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, localFormatted, m.GetLocalFormatted())
	}
}
