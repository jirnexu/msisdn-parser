package msisdn

import (
	"log"
)

type provider struct {
	startRange int
	endRange   int
	name       string
}

var countries map[string]string

// myProvider contains Malaysia mobile number with 7-digit mobile number
var myProviders map[string]provider

// myProviders2 contains Malaysia 6011 universal prefix with 8-digit mobile number
var myProviders2 []provider

func init() {
	err := loadData()
	if err != nil {
		log.Fatalln(err) // fixme: no error
	}
}

func loadData() error {
	countries = make(map[string]string, 3)
	countries = map[string]string{
		"60": "MY",
		"62": "ID",
		"65": "SG",
	}
	myProviders = make(map[string]provider, 11)
	myProviders = map[string]provider{
		"6010":  {name: "Digi Telecommunications Sdn Bhd"},
		"6012":  {name: "Maxis Broadband Sdn Bhd"},
		"6013":  {name: "Celcom Axiata Berhad"},
		"60142": {name: "Maxis Broadband Sdn Bhd"},
		"60143": {name: "Digi Telecommunications Sdn Bhd"},
		"60146": {name: "Digi Telecommunications Sdn Bhd"},
		"60149": {name: "Digi Telecommunications Sdn Bhd"},
		"6016":  {name: "Digi Telecommunications Sdn Bhd"},
		"6017":  {name: "Maxis Broadband Sdn Bhd"},
		"6018":  {name: "U Mobile Sdn Bhd"},
		"6019":  {name: "Celcom Axiata Berhad"},
	}
	myProviders2 = []provider{
		{startRange: 10000000, endRange: 10499999, name: "Webe Digital Sdn Bhd"},
		{startRange: 10500000, endRange: 10999999, name: "Red One Network Sdn. Bhd"},
		{startRange: 11000000, endRange: 11499999, name: "U Mobile Sdn Bhd"},
		{startRange: 11500000, endRange: 11999999, name: "U Mobile Sdn Bhd"},
		{startRange: 12000000, endRange: 12499999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 12500000, endRange: 12999999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 13000000, endRange: 13499999, name: "Xox Com Sdn Bhd"},
		{startRange: 14000000, endRange: 14499999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 14500000, endRange: 14999999, name: "Celcom Axiata Berhad"},
		{startRange: 15000000, endRange: 15499999, name: "Tune Talk Sdn Bhd"},
		{startRange: 15500000, endRange: 15999999, name: "Celcom Axiata Berhad"},
		{startRange: 16000000, endRange: 16499999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 16500000, endRange: 16999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 17000000, endRange: 17499999, name: "YTL Communications Sdn Bhd"},
		{startRange: 17500000, endRange: 17999999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 18000000, endRange: 18499999, name: "Telekom Malaysia Berhad"},
		{startRange: 18500000, endRange: 18999999, name: "Tune Talk Sdn Bhd"},
		{startRange: 19000000, endRange: 19499999, name: "Celcom Axiata Berhad"},
		{startRange: 19500000, endRange: 19999999, name: "Celcom Axiata Berhad"},
		{startRange: 20000000, endRange: 20499999, name: "Talk Focus Sdn Bhd"},
		{startRange: 20500000, endRange: 20999999, name: "Xox Com Sdn Bhd"},
		{startRange: 21000000, endRange: 21999999, name: "U Mobile Sdn Bhd"},
		{startRange: 23000000, endRange: 23999999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 24000000, endRange: 24499999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 24500000, endRange: 24999999, name: "Celcom Axiata Berhad"},
		{startRange: 25000000, endRange: 25999999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 26000000, endRange: 26999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 27000000, endRange: 27499999, name: "U Mobile Sdn Bhd"},
		{startRange: 27500000, endRange: 27999999, name: "Maxis Broadband Sdn Bhd"},
		{startRange: 28000000, endRange: 28999999, name: "U Mobile Sdn Bhd"},
		{startRange: 29000000, endRange: 29999999, name: "Celcom Axiata Berhad"},
		{startRange: 30000000, endRange: 30999999, name: "YTL Communications Sdn Bhd"},
		{startRange: 31000000, endRange: 31999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 32000000, endRange: 32499999, name: "Altel Communications Sdn Bhd"},
		{startRange: 32500000, endRange: 32999999, name: "Altel Communications Sdn Bhd"},
		{startRange: 33000000, endRange: 33999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 34000000, endRange: 34499999, name: "Enabling Asia Tech Sdn Bhd"},
		{startRange: 34500000, endRange: 34999999, name: "Enabling Asia Tech Sdn Bhd"},
		{startRange: 35000000, endRange: 35499999, name: "Tune Talk Sdn Bhd"},
		{startRange: 35500000, endRange: 35999999, name: "Tune Talk Sdn Bhd"},
		{startRange: 36000000, endRange: 36999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 37000000, endRange: 37999999, name: "U Mobile Sdn Bhd"},
		{startRange: 38000000, endRange: 38999999, name: "Ceres Telecom Sdn Bhd"},
		{startRange: 39000000, endRange: 39999999, name: "U Mobile Sdn Bhd"},
		{startRange: 40000000, endRange: 40499999, name: "Celcom Axiata Berhad"},
		{startRange: 40500000, endRange: 40999999, name: "Celcom Axiata Berhad"},
		{startRange: 41000000, endRange: 41499999, name: "Celcom Axiata Berhad"},
		{startRange: 41500000, endRange: 41999999, name: "Celcom Axiata Berhad"},
		{startRange: 42000000, endRange: 42499999, name: "Telekomunikasi Indonesia International (M) Sdn Bhd"},
		{startRange: 42500000, endRange: 42999999, name: "Telekomunikasi Indonesia International (M) Sdn Bhd"},
		{startRange: 50000000, endRange: 50499999, name: "Talk Focus Sdn Bhd"},
		{startRange: 51000000, endRange: 51499999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 51500000, endRange: 51999999, name: "Digi Telecommunications Sdn Bhd"},
		{startRange: 52000000, endRange: 52499999, name: "Altel Communications Sdn Bhd"},
		{startRange: 52500000, endRange: 52999999, name: "Altel Communications Sdn Bhd"},
		{startRange: 53000000, endRange: 53499999, name: "Tune Talk Sdn Bhd"},
		{startRange: 53500000, endRange: 53999999, name: "Tune Talk Sdn Bhd"},
		{startRange: 54000000, endRange: 54499999, name: "Celcom Axiata Berhad"},
		{startRange: 54500000, endRange: 54999999, name: "Celcom Axiata Berhad"},
		{startRange: 55000000, endRange: 55499999, name: "Red One Network Sdn. Bhd"},
		{startRange: 55500000, endRange: 55999999, name: "Red One Network Sdn. Bhd"},
		{startRange: 56000000, endRange: 56499999, name: "Celcom Axiata Berhad"},
		{startRange: 56500000, endRange: 56999999, name: "Xox Com Sdn Bhd"},
		{startRange: 57000000, endRange: 57499999, name: "YTL Communications Sdn Bhd"},
		{startRange: 57500000, endRange: 57999999, name: "Xox Com Sdn Bhd"},
		{startRange: 58000000, endRange: 58499999, name: "YTL Communications Sdn Bhd"},
		{startRange: 58500000, endRange: 58999999, name: "Xox Com Sdn Bhd"},
		{startRange: 59500000, endRange: 59999999, name: "Xox Com Sdn Bhd"},
		{startRange: 60500000, endRange: 60599999, name: "U Mobile Sdn Bhd"},
		{startRange: 60600000, endRange: 60699999, name: "U Mobile Sdn Bhd"},
		{startRange: 60700000, endRange: 60799999, name: "U Mobile Sdn Bhd"},
		{startRange: 60800000, endRange: 60899999, name: "U Mobile Sdn Bhd"},
		{startRange: 60900000, endRange: 60999999, name: "U Mobile Sdn Bhd"},
		{startRange: 61000000, endRange: 61099999, name: "U Mobile Sdn Bhd"},
		{startRange: 61100000, endRange: 61199999, name: "U Mobile Sdn Bhd"},
		{startRange: 61200000, endRange: 61299999, name: "U Mobile Sdn Bhd"},
		{startRange: 61300000, endRange: 61399999, name: "U Mobile Sdn Bhd"},
		{startRange: 61400000, endRange: 61499999, name: "U Mobile Sdn Bhd"},
		{startRange: 61500000, endRange: 61599999, name: "U Mobile Sdn Bhd"},
		{startRange: 61600000, endRange: 61699999, name: "U Mobile Sdn Bhd"},
		{startRange: 61700000, endRange: 61799999, name: "U Mobile Sdn Bhd"},
		{startRange: 61800000, endRange: 61899999, name: "U Mobile Sdn Bhd"},
		{startRange: 61900000, endRange: 61999999, name: "U Mobile Sdn Bhd"},
		{startRange: 62000000, endRange: 62099999, name: "U Mobile Sdn Bhd"},
		{startRange: 62100000, endRange: 62199999, name: "U Mobile Sdn Bhd"},
		{startRange: 62200000, endRange: 62299999, name: "U Mobile Sdn Bhd"},
		{startRange: 62300000, endRange: 62399999, name: "U Mobile Sdn Bhd"},
		{startRange: 62400000, endRange: 62499999, name: "U Mobile Sdn Bhd"},
	}
	return nil
}
