package msisdn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// import (
// 	"context"
// 	"testing"

// 	"bitbucket.org/ascensionlab/sdk/sdkerr"
// 	"bitbucket.org/ascensionlab/sdk/service/basenlp"
// 	"github.com/stretchr/testify/assert"
// )

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
