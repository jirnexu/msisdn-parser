package msisdn

import (
	"fmt"
	"regexp"
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
	m.msisdn = m.trim(msisdn)
	m.countryCode = m.getCountryCode()
	if !m.validate() {
		return fmt.Errorf("MSISDN is invalid")
	}
	m.provider = m.getProvider()
	m.landline = m.isLandLine()
	return nil
}

//
func (m *MSISDN) LocalToMSISDN(phone string) string {
	phone = m.trim(phone)
	if m.isMSISDN(phone) {
		return phone
	}
	// FIXME: handle by country code. Now only support MY
	msisdn := fmt.Sprintf("6%s", phone)
	err := m.Parse(msisdn)
	if err != nil {
		return ""
	}
	return msisdn
}

func (m *MSISDN) getCountryCode() string {
	prefix := m.msisdn[:2]
	for cc, country := range countries {
		if prefix == country.countryCode {
			return cc
		}
	}
	return ""
}

func (m *MSISDN) getProvider() string {
	if m.countryCode == "" {
		return ""
	}

	subscriberNumber, _ := strconv.Atoi(m.msisdn)
	sn := int64(subscriberNumber)

	for _, p := range countries[m.countryCode].providers {
		if sn >= p.startRange && sn <= p.endRange {
			return p.name
		}
	}
	return ""
}

func (m *MSISDN) isMSISDN(s string) bool {
	if strings.HasPrefix(s, "1800") {
		return true
	}
	for _, country := range countries {
		if strings.HasPrefix(s, country.countryCode) {
			err := m.Parse(s)
			if err != nil {
				return false
			}
			return true
		}
	}
	return false
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
	// must not exceed 15 digits
	if len(m.msisdn) > 15 {
		return false
	}
	country := countries[m.countryCode]
	if len(m.msisdn) < country.minLength || len(m.msisdn) > country.maxLength {
		return false
	}
	return true
}

func (m *MSISDN) trim(s string) string {
	notWordRE := regexp.MustCompile(`[\W_]`)
	return notWordRE.ReplaceAllString(s, "")
}
