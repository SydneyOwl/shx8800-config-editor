package radio

type ClassTheRadioData struct {
	ChannelData  [128][]string  `json:"channelData"`
	FunCfgData   FormFunCFGData `json:"funCfgData"`
	DtmfData     DTMFData       `json:"dtmfData"`
	OtherImfData OtherImfData   `json:"otherImfData"`
}
