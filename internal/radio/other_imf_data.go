package radio

type OtherImfData struct {
	TheRangeOfVHF    int    `json:"theRangeOfVHF"`
	TheMinFreqOfVHF  string `json:"theMinFreqOfVHF"`
	TheMaxFreqOfVHF  string `json:"theMaxFreqOfVHF"`
	TheRangeOfUHF    int    `json:"theRangeOfUHF"`
	TheMinFreqOfUHF  string `json:"theMinFreqOfUHF"`
	TheMaxFreqOfUHF  string `json:"theMaxFreqOfUHF"`
	EnableTxUHF      bool   `json:"enableTxUHF"`
	EnableTxVHF      bool   `json:"enableTxVHF"`
	PowerUpChar1     string `json:"powerUpChar1"`
	PowerUpChar2     string `json:"powerUpChar2"`
	EnableTxOver480M bool   `json:"enableTxOver480M"`
}

func CreateOIMFData() OtherImfData {
	return OtherImfData{
		TheRangeOfVHF:    0,
		TheMinFreqOfVHF:  "136",
		TheMaxFreqOfVHF:  "174",
		TheRangeOfUHF:    0,
		TheMinFreqOfUHF:  "400",
		TheMaxFreqOfUHF:  "520",
		EnableTxUHF:      true,
		EnableTxVHF:      true,
		PowerUpChar1:     "BAOFENG",
		PowerUpChar2:     "UV-5R",
		EnableTxOver480M: false,
	}
}
