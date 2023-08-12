package radio

type FormFunCFGData struct {
	CbB_AutoBackLight        int    `json:"cbB_AutoBackLight"`
	CbB_VoicSwitch           int    `json:"cbB_VoicSwitch"`
	CbB_Language             int    `json:"cbB_Language"`
	CbB_VOX                  int    `json:"cbB_VOX"`
	CbB_SQL                  int    `json:"cbB_SQL"`
	CbB_TOT                  int    `json:"cbB_TOT"`
	CbB_PowerOnMsg           int    `json:"cbB_PowerOnMsg"`
	CbB_VoxDelay             int    `json:"cbB_VoxDelay"`
	CbB_TimerMenuQuit        int    `json:"cbB_TimerMenuQuit"`
	CbB_MicGain              int    `json:"cbB_MicGain"`
	CbB_BackgroundColor      int    `json:"cbB_BackgroundColor"`
	TB_CountsOfCH            string `json:"tB_CountsOfCH"`
	CbB_WorkModeA            int    `json:"cbB_WorkModeA"`
	CbB_WorkModeB            int    `json:"cbB_WorkModeB"`
	CbB_CH_B_DisplayMode     int    `json:"cbB_CH_B_DisplayMode"`
	CbB_CH_A_DisplayMode     int    `json:"cbB_CH_A_DisplayMode"`
	CbB_SendIDDelay          int    `json:"cbB_SendIDDelay"`
	CbB_PTTID                int    `json:"cbB_PTTID"`
	CbB_Scan                 int    `json:"cbB_Scan"`
	CbB_SaveMode             int    `json:"cbB_SaveMode"`
	CbB_DTMF                 int    `json:"cbB_DTMF"`
	CbB_KeySide              int    `json:"cbB_KeySide"`
	CbB_KeySideL             int    `json:"cbB_KeySideL"`
	CB_SoundOfBi             bool   `json:"cB_SoundOfBi"`
	CB_StopSendOnBusy        bool   `json:"cB_StopSendOnBusy"`
	CB_AutoLock              bool   `json:"cB_AutoLock"`
	CB_LockKeyBoard          bool   `json:"cB_LockKeyBoard"`
	TB_A_RangeOfFreq         string `json:"tB_A_RangeOfFreq"`
	TB_A_CurFreq             string `json:"tB_A_CurFreq"`
	CbB_A_SignalingEnCoder   int    `json:"cbB_A_SignalingEnCoder"`
	TB_A_RemainFreq          string `json:"tB_A_RemainFreq"`
	CbB_A_RemainDir          int    `json:"cbB_A_RemainDir"`
	CbB_A_FreqStep           int    `json:"cbB_A_FreqStep"`
	CbB_A_CHBand             int    `json:"cbB_A_CHBand"`
	CbB_A_TxQT               string `json:"cbB_A_TxQT"`
	CbB_A_RxQT               string `json:"cbB_A_RxQT"`
	CbB_A_Power              int    `json:"cbB_A_Power"`
	CbB_A_Band               int    `json:"cbB_A_Band"`
	CbB_A_Fhss               int    `json:"cbB_A_Fhss"`
	TB_B_RangeOfFreq         string `json:"tB_B_RangeOfFreq"`
	TB_B_CurFreq             string `json:"tB_B_CurFreq"`
	CbB_B_SignalingEnCoder   int    `json:"cbB_B_SignalingEnCoder"`
	TB_B_RemainFreq          string `json:"tB_B_RemainFreq"`
	CbB_B_RemainDir          int    `json:"cbB_B_RemainDir"`
	CbB_B_FreqStep           int    `json:"cbB_B_FreqStep"`
	CbB_B_CHBand             int    `json:"cbB_B_CHBand"`
	CbB_B_TxQT               string `json:"cbB_B_TxQT"`
	CbB_B_RxQT               string `json:"cbB_B_RxQT"`
	CbB_B_Power              int    `json:"cbB_B_Power"`
	CbB_B_Band               int    `json:"cbB_B_Band"`
	CbB_B_Fhss               int    `json:"cbB_B_Fhss"`
	CbB_DisplayTypeOfPowerUp int    `json:"cbB_DisplayTypeOfPowerUp"`
	CbB_PassRepetNoiseDetect int    `json:"cbB_PassRepetNoiseDetect"`
	CbB_PassRepetNoiseClear  int    `json:"cbB_PassRepetNoiseClear"`
	CbB_TailNoiseClear       int    `json:"cbB_TailNoiseClear"`
	CbB_1750Hz               int    `json:"cbB_1750Hz"`
	CbB_TxUnderTDRStart      int    `json:"cbB_TxUnderTDRStart"`
	CbB_SoundOfTxEnd         int    `json:"cbB_SoundOfTxEnd"`
	CbB_AlarmMode            int    `json:"cbB_AlarmMode"`
	CB_AlarmSound            bool   `json:"cB_AlarmSound"`
	CB_FMRadioEnable         bool   `json:"cB_FMRadioEnable"`
	CB_TDR                   bool   `json:"cB_TDR"`
}

func CreateFFCData() FormFunCFGData {
	return FormFunCFGData{
		CbB_AutoBackLight:        5,
		CbB_VoicSwitch:           1,
		CbB_Language:             1,
		CbB_VOX:                  0,
		CbB_SQL:                  3,
		CbB_TOT:                  0,
		CbB_PowerOnMsg:           0,
		CbB_VoxDelay:             5,
		CbB_TimerMenuQuit:        1,
		CbB_MicGain:              2,
		CbB_BackgroundColor:      0,
		TB_CountsOfCH:            "128",
		CbB_WorkModeA:            0,
		CbB_WorkModeB:            0,
		CbB_CH_B_DisplayMode:     1,
		CbB_CH_A_DisplayMode:     1,
		CbB_SendIDDelay:          5,
		CbB_PTTID:                0,
		CbB_Scan:                 1,
		CbB_SaveMode:             1,
		CbB_DTMF:                 3,
		CbB_KeySide:              0,
		CbB_KeySideL:             1,
		CB_SoundOfBi:             true,
		CB_StopSendOnBusy:        false,
		CB_AutoLock:              false,
		CB_LockKeyBoard:          false,
		TB_A_RangeOfFreq:         "136-173.99750",
		TB_A_CurFreq:             "146.02500",
		CbB_A_SignalingEnCoder:   0,
		TB_A_RemainFreq:          "00.0000",
		CbB_A_RemainDir:          0,
		CbB_A_FreqStep:           5,
		CbB_A_CHBand:             0,
		CbB_A_TxQT:               "OFF",
		CbB_A_RxQT:               "OFF",
		CbB_A_Power:              0,
		CbB_A_Band:               0,
		CbB_A_Fhss:               0,
		TB_B_RangeOfFreq:         "400-519.99750",
		TB_B_CurFreq:             "440.02500",
		CbB_B_SignalingEnCoder:   0,
		TB_B_RemainFreq:          "00.0000",
		CbB_B_RemainDir:          0,
		CbB_B_FreqStep:           5,
		CbB_B_CHBand:             0,
		CbB_B_TxQT:               "OFF",
		CbB_B_RxQT:               "OFF",
		CbB_B_Power:              0,
		CbB_B_Band:               1,
		CbB_B_Fhss:               0,
		CbB_DisplayTypeOfPowerUp: 0,
		CbB_PassRepetNoiseDetect: 0,
		CbB_PassRepetNoiseClear:  5,
		CbB_TailNoiseClear:       1,
		CbB_1750Hz:               2,
		CbB_TxUnderTDRStart:      0,
		CbB_SoundOfTxEnd:         0,
		CbB_AlarmMode:            1,
		CB_AlarmSound:            true,
		CB_FMRadioEnable:         true,
		CB_TDR:                   false,
	}
}
