package msisdn

import (
	"fmt"
	"strconv"
	"strings"
)

//
type MSISDN struct {
	msisdn      string
	countryCode string
	landline    bool
	provider    string
}

//
func (m *MSISDN) GetMSISDN() string {
	return m.msisdn
}

//
func (m *MSISDN) GetProvider() string {
	return m.provider
}

//
func (m *MSISDN) GetCountryCode() string {
	return m.countryCode
}

//
func (m *MSISDN) IsLandLine() bool {
	return m.landline
}

//
func (m *MSISDN) Parse(msisdn string) error {
	m.msisdn = strings.TrimLeft(msisdn, " +")
	m.msisdn = strings.TrimSpace(m.msisdn)

	if !m.validate() {
		return fmt.Errorf("MSISDN is invalid")
	}

	prefix := m.msisdn[:2]
	m.countryCode = countries[prefix]

	if m.countryCode != "" {
		m.provider = m.getProvider()
	}

	m.landline = m.isLandLine()

	return nil
}

//
func (m *MSISDN) LocalToMSISDN(phone string) string {
	if strings.HasPrefix(phone, "60") ||
		strings.HasPrefix(phone, "65") ||
		strings.HasPrefix(phone, "62") ||
		strings.HasPrefix(phone, "1800") {
		return phone
	}
	// FIXME: handle by country code. Now only support MY
	return fmt.Sprintf("6%s", phone)
}

func (m *MSISDN) getProvider() string {
	prefix := m.msisdn[:4]
	subscriberNumber := m.msisdn[4:]
	sn, _ := strconv.Atoi(subscriberNumber)
	switch prefix {
	case "6011":
		for _, p := range myProviders2 {
			if sn >= p.startRange && sn <= p.endRange {
				return p.name
			}
		}
	case "6014":
		prefix = m.msisdn[:5]
		return myProviders[prefix].name
	default:
		return myProviders[prefix].name
	}
	return ""
}

func (m *MSISDN) isLandLine() bool {
	if strings.HasPrefix(m.msisdn, "1800") {
		return true
	}
	switch m.countryCode {
	case "MY":
		if !strings.HasPrefix(m.msisdn, "601") {
			return true
		}
	default:
		return false
	}
	return false
}

func (m *MSISDN) validate() bool {
	// must be number and not empty
	_, err := strconv.Atoi(m.msisdn)
	if err != nil {
		return false
	}
	// length must greater than reserved digit for country code
	if len(m.msisdn) < 4 {
		return false
	}
	// must not exceed 15 digits
	if len(m.msisdn) > 15 {
		return false
	}
	return true
}
