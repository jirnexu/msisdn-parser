package msisdn

import (
	"log"
)

type provider struct {
	countryCode string
	name        string
}

var countries map[string]string
var providers map[string]provider

func init() {
	err := loadData()
	if err != nil {
		log.Fatalln(err)
	}
}

func loadData() error {
	countries = make(map[string]string, 3)
	countries = map[string]string{
		"60": "MY",
		"62": "ID",
		"65": "SG",
	}

	providers = make(map[string]provider, 0)
	providers = map[string]provider{
		"6010": {
			countryCode: "MY",
			name:        "Digi",
		},
		"6012": {
			countryCode: "MY",
			name:        "Maxis",
		},
		"6013": {
			countryCode: "MY",
			name:        "Celcom",
		},
		"60142": {
			countryCode: "MY",
			name:        "Maxis",
		},
		"60143": {
			countryCode: "MY",
			name:        "Digi",
		},
		"60146": {
			countryCode: "MY",
			name:        "Digi",
		},
		"60149": {
			countryCode: "MY",
			name:        "Digi",
		},
		"6016": {
			countryCode: "MY",
			name:        "Digi",
		},
		"6017": {
			countryCode: "MY",
			name:        "Maxis",
		},
		"6018": {
			countryCode: "MY",
			name:        "U Mobile",
		},
		"6019": {
			countryCode: "MY",
			name:        "Celcom",
		},
	}
	return nil
}
