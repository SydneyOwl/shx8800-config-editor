package radio

type DTMFData struct {
	GroupOfDTMF_1     string `json:"groupOfDTMF_1"`
	GroupOfDTMF_2     string `json:"groupOfDTMF_2"`
	GroupOfDTMF_3     string `json:"groupOfDTMF_3"`
	GroupOfDTMF_4     string `json:"groupOfDTMF_4"`
	GroupOfDTMF_5     string `json:"groupOfDTMF_5"`
	GroupOfDTMF_6     string `json:"groupOfDTMF_6"`
	GroupOfDTMF_7     string `json:"groupOfDTMF_7"`
	GroupOfDTMF_8     string `json:"groupOfDTMF_8"`
	GroupOfDTMF_9     string `json:"groupOfDTMF_9"`
	GroupOfDTMF_A     string `json:"groupOfDTMF_A"`
	GroupOfDTMF_B     string `json:"groupOfDTMF_B"`
	GroupOfDTMF_C     string `json:"groupOfDTMF_C"`
	GroupOfDTMF_D     string `json:"groupOfDTMF_D"`
	GroupOfDTMF_E     string `json:"groupOfDTMF_E"`
	GroupOfDTMF_F     string `json:"groupOfDTMF_F"`
	LastTimeSend      int    `json:"lastTimeSend"`
	LastTimeStop      int    `json:"lastTimeStop"`
	GroupCall         int    `json:"groupCall"`
	TheIDOfLocalHost  string `json:"theIDOfLocalHost"`
	SendOnPTTPressed  bool   `json:"sendOnPTTPressed"`
	SendOnPTTReleased bool   `json:"sendOnPTTReleased"`
}

func CreateDTMFData() DTMFData {
	return DTMFData{
		GroupOfDTMF_1:     "20202",
		GroupOfDTMF_2:     "",
		GroupOfDTMF_3:     "",
		GroupOfDTMF_4:     "",
		GroupOfDTMF_5:     "",
		GroupOfDTMF_6:     "",
		GroupOfDTMF_7:     "",
		GroupOfDTMF_8:     "",
		GroupOfDTMF_9:     "",
		GroupOfDTMF_A:     "",
		GroupOfDTMF_B:     "",
		GroupOfDTMF_C:     "",
		GroupOfDTMF_D:     "",
		GroupOfDTMF_E:     "",
		GroupOfDTMF_F:     "30303",
		LastTimeSend:      1,
		LastTimeStop:      1,
		GroupCall:         14,
		TheIDOfLocalHost:  "80808",
		SendOnPTTPressed:  false,
		SendOnPTTReleased: true,
	}
}
