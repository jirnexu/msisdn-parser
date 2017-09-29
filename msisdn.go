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
	m.msisdn = strings.TrimLeft(msisdn, "+")
	m.msisdn = strings.TrimSpace(m.msisdn)
	if !m.validate() {
		return fmt.Errorf("MSISDN is invalid")
	}

	prefix := m.msisdn[:2]
	m.countryCode = countries[prefix]

	if m.countryCode != "" {
		prefix := m.msisdn[:4]
		m.provider = providers[prefix].name
	}

	m.landline = m.isLandLine()

	return nil
}

func (m *MSISDN) isLandLine() bool {
	if strings.HasPrefix(m.msisdn, "1800") {
		return true
	}
	if m.countryCode == "MY" {
		if !strings.HasPrefix(m.msisdn, "601") {
			return true
		}
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
