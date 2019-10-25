package msisdn

import (
	"regexp"
)

type phoneFormat struct {
	find    *regexp.Regexp
	replace string
}

type country struct {
	providers    []provider
	areaCode     string
	localPrefix  string
	localFormats []phoneFormat

	minLength int
	maxLength int
}

type provider struct {
	startRange int64
	endRange   int64
	name       string
}

var countries = map[string]*country{
	"MY": {
		areaCode:    "60",
		localPrefix: "0",
		minLength:   8,
		maxLength:   12,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(0[345679]|08[0-9]|01[0-9])(\d{4})(\d{4})$`),
				replace: `$1-$2 $3`,
			},
			{
				find:    regexp.MustCompile(`^(0[345679]|08[0-9]|01[0-9])(\d{3})(\d{3,4})$`),
				replace: `$1-$2 $3`,
			},
		},
		providers: []provider{
			// 7-digit numbers
			{startRange: 60100000000, endRange: 60109999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 60120000000, endRange: 60129999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 60130000000, endRange: 60139999999, name: "Celcom Axiata Berhad"},
			{startRange: 60142000000, endRange: 60142999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 60143000000, endRange: 60143999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 60146000000, endRange: 60146999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 60149000000, endRange: 60149999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 60160000000, endRange: 60169999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 60170000000, endRange: 60179999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 60180000000, endRange: 60189999999, name: "U Mobile Sdn Bhd"},
			{startRange: 60190000000, endRange: 60199999999, name: "Celcom Axiata Berhad"},
			// 6011 8-digit numbers
			{startRange: 601110000000, endRange: 601110499999, name: "Webe Digital Sdn Bhd"},
			{startRange: 601110500000, endRange: 601110999999, name: "Red One Network Sdn. Bhd"},
			{startRange: 601111000000, endRange: 601111499999, name: "U Mobile Sdn Bhd"},
			{startRange: 601111500000, endRange: 601111999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601112000000, endRange: 601112499999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601112500000, endRange: 601112999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601113000000, endRange: 601113499999, name: "Xox Com Sdn Bhd"},
			{startRange: 601114000000, endRange: 601114499999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601114500000, endRange: 601114999999, name: "Celcom Axiata Berhad"},
			{startRange: 601115000000, endRange: 601115499999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601115500000, endRange: 601115999999, name: "Celcom Axiata Berhad"},
			{startRange: 601116000000, endRange: 601116499999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601116500000, endRange: 601116999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601117000000, endRange: 601117499999, name: "YTL Communications Sdn Bhd"},
			{startRange: 601117500000, endRange: 601117999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601118000000, endRange: 601118499999, name: "Telekom Malaysia Berhad"},
			{startRange: 601118500000, endRange: 601118999999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601119000000, endRange: 601119499999, name: "Celcom Axiata Berhad"},
			{startRange: 601119500000, endRange: 601119999999, name: "Celcom Axiata Berhad"},
			{startRange: 601120000000, endRange: 601120499999, name: "Talk Focus Sdn Bhd"},
			{startRange: 601120500000, endRange: 601120999999, name: "Xox Com Sdn Bhd"},
			{startRange: 601121000000, endRange: 601121999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601123000000, endRange: 601123999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601124000000, endRange: 601124499999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601124500000, endRange: 601124999999, name: "Celcom Axiata Berhad"},
			{startRange: 601125000000, endRange: 601125999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601126000000, endRange: 601126999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601127000000, endRange: 601127499999, name: "U Mobile Sdn Bhd"},
			{startRange: 601127500000, endRange: 601127999999, name: "Maxis Broadband Sdn Bhd"},
			{startRange: 601128000000, endRange: 601128999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601129000000, endRange: 601129999999, name: "Celcom Axiata Berhad"},
			{startRange: 601130000000, endRange: 601130999999, name: "YTL Communications Sdn Bhd"},
			{startRange: 601131000000, endRange: 601131999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601132000000, endRange: 601132499999, name: "Altel Communications Sdn Bhd"},
			{startRange: 601132500000, endRange: 601132999999, name: "Altel Communications Sdn Bhd"},
			{startRange: 601133000000, endRange: 601133999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601134000000, endRange: 601134499999, name: "Enabling Asia Tech Sdn Bhd"},
			{startRange: 601134500000, endRange: 601134999999, name: "Enabling Asia Tech Sdn Bhd"},
			{startRange: 601135000000, endRange: 601135499999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601135500000, endRange: 601135999999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601136000000, endRange: 601136999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601137000000, endRange: 601137999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601138000000, endRange: 601138999999, name: "Ceres Telecom Sdn Bhd"},
			{startRange: 601139000000, endRange: 601139999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601140000000, endRange: 601140499999, name: "Celcom Axiata Berhad"},
			{startRange: 601140500000, endRange: 601140999999, name: "Celcom Axiata Berhad"},
			{startRange: 601141000000, endRange: 601141499999, name: "Celcom Axiata Berhad"},
			{startRange: 601141500000, endRange: 601141999999, name: "Celcom Axiata Berhad"},
			{startRange: 601142000000, endRange: 601142499999, name: "Telekomunikasi Indonesia International (M) Sdn Bhd"},
			{startRange: 601142500000, endRange: 601142999999, name: "Telekomunikasi Indonesia International (M) Sdn Bhd"},
			{startRange: 601150000000, endRange: 601150499999, name: "Talk Focus Sdn Bhd"},
			{startRange: 601151000000, endRange: 601151499999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601151500000, endRange: 601151999999, name: "Digi Telecommunications Sdn Bhd"},
			{startRange: 601152000000, endRange: 601152499999, name: "Altel Communications Sdn Bhd"},
			{startRange: 601152500000, endRange: 601152999999, name: "Altel Communications Sdn Bhd"},
			{startRange: 601153000000, endRange: 601153499999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601153500000, endRange: 601153999999, name: "Tune Talk Sdn Bhd"},
			{startRange: 601154000000, endRange: 601154499999, name: "Celcom Axiata Berhad"},
			{startRange: 601154500000, endRange: 601154999999, name: "Celcom Axiata Berhad"},
			{startRange: 601155000000, endRange: 601155499999, name: "Red One Network Sdn. Bhd"},
			{startRange: 601155500000, endRange: 601155999999, name: "Red One Network Sdn. Bhd"},
			{startRange: 601156000000, endRange: 601156499999, name: "Celcom Axiata Berhad"},
			{startRange: 601156500000, endRange: 601156999999, name: "Xox Com Sdn Bhd"},
			{startRange: 601157000000, endRange: 601157499999, name: "YTL Communications Sdn Bhd"},
			{startRange: 601157500000, endRange: 601157999999, name: "Xox Com Sdn Bhd"},
			{startRange: 601158000000, endRange: 601158499999, name: "YTL Communications Sdn Bhd"},
			{startRange: 601158500000, endRange: 601158999999, name: "Xox Com Sdn Bhd"},
			{startRange: 601159500000, endRange: 601159999999, name: "Xox Com Sdn Bhd"},
			{startRange: 601160500000, endRange: 601160599999, name: "U Mobile Sdn Bhd"},
			{startRange: 601160600000, endRange: 601160699999, name: "U Mobile Sdn Bhd"},
			{startRange: 601160700000, endRange: 601160799999, name: "U Mobile Sdn Bhd"},
			{startRange: 601160800000, endRange: 601160899999, name: "U Mobile Sdn Bhd"},
			{startRange: 601160900000, endRange: 601160999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161000000, endRange: 601161099999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161100000, endRange: 601161199999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161200000, endRange: 601161299999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161300000, endRange: 601161399999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161400000, endRange: 601161499999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161500000, endRange: 601161599999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161600000, endRange: 601161699999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161700000, endRange: 601161799999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161800000, endRange: 601161899999, name: "U Mobile Sdn Bhd"},
			{startRange: 601161900000, endRange: 601161999999, name: "U Mobile Sdn Bhd"},
			{startRange: 601162000000, endRange: 601162099999, name: "U Mobile Sdn Bhd"},
			{startRange: 601162100000, endRange: 601162199999, name: "U Mobile Sdn Bhd"},
			{startRange: 601162200000, endRange: 601162299999, name: "U Mobile Sdn Bhd"},
			{startRange: 601162300000, endRange: 601162399999, name: "U Mobile Sdn Bhd"},
			{startRange: 601162400000, endRange: 601162499999, name: "U Mobile Sdn Bhd"},
		},
	},
	"ID": {
		areaCode:    "62",
		localPrefix: "",
		minLength:   8,
		maxLength:   12,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(\d{3})(\d{3})(\d{4})$`),
				replace: `$1 $2 $3`, // +62 yyy xxx xxxx
			},
			{
				find:    regexp.MustCompile(`^(\d{3})(\d{4})(\d{4})$`),
				replace: `$1 $2 $3`, // +62 yyy xxxx xxxx
			},
		},
	},
	"SG": {
		areaCode:    "65",
		localPrefix: "",
		minLength:   8,
		maxLength:   12,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(\d{4})(\d{4})$`),
				replace: `$1 $2`,
			},
		},
	},
	"AU": {
		areaCode:    "61",
		localPrefix: "",
		minLength:   8,
		maxLength:   12,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(4|5)(\d{4})(\d{4})$`),
				replace: `0$1 $2 $3`, // 0(4|5) yyyy xxxx
			},
		},
	},
	"HK": {
		areaCode:    "852",
		localPrefix: "",
		minLength:   8,
		maxLength:   11,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(4|5|6|7|8|9)(\d{4})(\d{4})$`),
				replace: `0$1 $2 $3`, // 0(4|5|6|7|8|9) yyyy xxxx
			},
		},
	},
	"JP": {
		areaCode:    "81",
		localPrefix: "",
		minLength:   11,
		maxLength:   14,
		localFormats: []phoneFormat{
			{
				find:    regexp.MustCompile(`^(70|80|90)(\d{4})(\d{4})$`),
				replace: `0$1 $2 $3`, // 0(7|8|9) yyyy xxxx
			},
		},
	},
}
