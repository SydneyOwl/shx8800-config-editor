using System;
using System.IO;
using Newtonsoft.Json;
using System.Runtime.Serialization.Formatters.Binary;
namespace SQ5R
{
#pragma warning disable SYSLIB0011
    public class TransJson
    {
        static void Main(string[] args)
        {
            try
            {
                if (args[0] == "transjson")
                {
                    Stream s = new FileStream(args[1], FileMode.Open);
                    ClassTheRadioData deserializedObject = ClassTheRadioData.CreatObjFromFile(s);
                    string jsonData = JsonConvert.SerializeObject(deserializedObject);
                    Console.WriteLine(jsonData);
                }
                if (args[0]=="transfile"){
                    var radioData = JsonConvert.DeserializeObject<ClassTheRadioData>(@args[2]);
                    using FileStream fs = File.Create(@args[1]);
                    BinaryFormatter binaryFormatter = new BinaryFormatter();
                    binaryFormatter.Serialize(fs, radioData);
                }
            }
            catch
            {
                Console.WriteLine("-1");
            }
        }
    }

    [Serializable]
    public class ClassTheRadioData
    {
        private string[][] channelData = new string[128][];

        public FormFunCFGData funCfgData = new FormFunCFGData();

        public DTMFData dtmfData = new DTMFData();

        public OtherImfData otherImfData = new OtherImfData();

        public string[][] ChannelData
        {
            get
            {
                return channelData;
            }
            set
            {
                channelData = value;
            }
        }
        public static ClassTheRadioData CreatObjFromFile(Stream s)
        {
            BinaryFormatter binaryFormatter = new BinaryFormatter();
            return binaryFormatter.Deserialize(s) as ClassTheRadioData;
        }
    }

    [Serializable]
    public class DTMFData
    {
        private string groupOfDTMF_1 = "20202";

        private string groupOfDTMF_2 = "";

        private string groupOfDTMF_3 = "";

        private string groupOfDTMF_4 = "";

        private string groupOfDTMF_5 = "";

        private string groupOfDTMF_6 = "";

        private string groupOfDTMF_7 = "";

        private string groupOfDTMF_8 = "";

        private string groupOfDTMF_9 = "";

        private string groupOfDTMF_A = "";

        private string groupOfDTMF_B = "";

        private string groupOfDTMF_C = "";

        private string groupOfDTMF_D = "";

        private string groupOfDTMF_E = "";

        private string groupOfDTMF_F = "30303";

        private int lastTimeSend = 1;

        private int lastTimeStop = 1;

        private int groupCall = 14;

        private string theIDOfLocalHost = "80808";

        private bool sendOnPTTPressed = false;

        private bool sendOnPTTReleased = true;

        public string GroupOfDTMF_1
        {
            get
            {
                return groupOfDTMF_1;
            }
            set
            {
                groupOfDTMF_1 = value;
            }
        }

        public string GroupOfDTMF_2
        {
            get
            {
                return groupOfDTMF_2;
            }
            set
            {
                groupOfDTMF_2 = value;
            }
        }

        public string GroupOfDTMF_3
        {
            get
            {
                return groupOfDTMF_3;
            }
            set
            {
                groupOfDTMF_3 = value;
            }
        }

        public string GroupOfDTMF_4
        {
            get
            {
                return groupOfDTMF_4;
            }
            set
            {
                groupOfDTMF_4 = value;
            }
        }

        public string GroupOfDTMF_5
        {
            get
            {
                return groupOfDTMF_5;
            }
            set
            {
                groupOfDTMF_5 = value;
            }
        }

        public string GroupOfDTMF_6
        {
            get
            {
                return groupOfDTMF_6;
            }
            set
            {
                groupOfDTMF_6 = value;
            }
        }

        public string GroupOfDTMF_7
        {
            get
            {
                return groupOfDTMF_7;
            }
            set
            {
                groupOfDTMF_7 = value;
            }
        }

        public string GroupOfDTMF_8
        {
            get
            {
                return groupOfDTMF_8;
            }
            set
            {
                groupOfDTMF_8 = value;
            }
        }

        public string GroupOfDTMF_9
        {
            get
            {
                return groupOfDTMF_9;
            }
            set
            {
                groupOfDTMF_9 = value;
            }
        }

        public string GroupOfDTMF_A
        {
            get
            {
                return groupOfDTMF_A;
            }
            set
            {
                groupOfDTMF_A = value;
            }
        }

        public string GroupOfDTMF_B
        {
            get
            {
                return groupOfDTMF_B;
            }
            set
            {
                groupOfDTMF_B = value;
            }
        }

        public string GroupOfDTMF_C
        {
            get
            {
                return groupOfDTMF_C;
            }
            set
            {
                groupOfDTMF_C = value;
            }
        }

        public string GroupOfDTMF_D
        {
            get
            {
                return groupOfDTMF_D;
            }
            set
            {
                groupOfDTMF_D = value;
            }
        }

        public string GroupOfDTMF_E
        {
            get
            {
                return groupOfDTMF_E;
            }
            set
            {
                groupOfDTMF_E = value;
            }
        }

        public string GroupOfDTMF_F
        {
            get
            {
                return groupOfDTMF_F;
            }
            set
            {
                groupOfDTMF_F = value;
            }
        }

        public int LastTimeSend
        {
            get
            {
                return lastTimeSend;
            }
            set
            {
                lastTimeSend = value;
            }
        }

        public int LastTimeStop
        {
            get
            {
                return lastTimeStop;
            }
            set
            {
                lastTimeStop = value;
            }
        }

        public int GroupCall
        {
            get
            {
                return groupCall;
            }
            set
            {
                groupCall = value;
            }
        }

        public string TheIDOfLocalHost
        {
            get
            {
                return theIDOfLocalHost;
            }
            set
            {
                theIDOfLocalHost = value;
            }
        }

        public bool SendOnPTTPressed
        {
            get
            {
                return sendOnPTTPressed;
            }
            set
            {
                sendOnPTTPressed = value;
            }
        }

        public bool SendOnPTTReleased
        {
            get
            {
                return sendOnPTTReleased;
            }
            set
            {
                sendOnPTTReleased = value;
            }
        }
    }


    [Serializable]
    public class FormFunCFGData
    {
        private int cbB_AutoBackLight = 5;

        private int cbB_VoicSwitch = 1;

        private int cbB_Language = 1;

        private int cbB_VOX = 0;

        private int cbB_SQL = 3;

        private int cbB_TOT = 0;

        private int cbB_PowerOnMsg = 0;

        private int cbB_VoxDelay = 5;

        private int cbB_TimerMenuQuit = 1;

        private int cbB_MicGain = 2;

        private int cbB_BackgroundColor = 0;

        private string tB_CountsOfCH = "128";

        private int cbB_WorkModeA = 0;

        private int cbB_WorkModeB = 0;

        private int cbB_CH_B_DisplayMode = 1;

        private int cbB_CH_A_DisplayMode = 1;

        private int cbB_SendIDDelay = 5;

        private int cbB_PTTID = 0;

        private int cbB_Scan = 1;

        private int cbB_SaveMode = 1;

        private int cbB_DTMF = 3;

        private int cbB_KeySide = 0;

        private int cbB_KeySideL = 1;

        private bool cB_SoundOfBi = true;

        private bool cB_StopSendOnBusy = false;

        private bool cB_AutoLock = false;

        private bool cB_LockKeyBoard = false;

        private string tB_A_RangeOfFreq = "136-173.99750";

        private string tB_A_CurFreq = "146.02500";

        private int cbB_A_SignalingEnCoder = 0;

        private string tB_A_RemainFreq = "00.0000";

        private int cbB_A_RemainDir = 0;

        private int cbB_A_FreqStep = 5;

        private int cbB_A_CHBand = 0;

        private string cbB_A_TxQT = "OFF";

        private string cbB_A_RxQT = "OFF";

        private int cbB_A_Power = 0;

        private int cbB_A_Band = 0;

        private int cbB_A_Fhss = 0;

        private string tB_B_RangeOfFreq = "400-519.99750";

        private string tB_B_CurFreq = "440.02500";

        private int cbB_B_SignalingEnCoder = 0;

        private string tB_B_RemainFreq = "00.0000";

        private int cbB_B_RemainDir = 0;

        private int cbB_B_FreqStep = 5;

        private int cbB_B_CHBand = 0;

        private string cbB_B_TxQT = "OFF";

        private string cbB_B_RxQT = "OFF";

        private int cbB_B_Power = 0;

        private int cbB_B_Band = 1;

        private int cbB_B_Fhss = 0;

        private int cbB_DisplayTypeOfPowerUp = 0;

        private int cbB_PassRepetNoiseDetect = 0;

        private int cbB_PassRepetNoiseClear = 5;

        private int cbB_TailNoiseClear = 1;

        private int cbB_1750Hz = 2;

        private int cbB_TxUnderTDRStart = 0;

        private int cbB_SoundOfTxEnd = 0;

        private int cbB_AlarmMode = 1;

        private bool cB_AlarmSound = true;

        private bool cB_FMRadioEnable = true;

        private bool cB_TDR = false;

        public int CbB_AutoBackLight
        {
            get
            {
                return cbB_AutoBackLight;
            }
            set
            {
                cbB_AutoBackLight = value;
            }
        }

        public int CbB_VOX
        {
            get
            {
                return cbB_VOX;
            }
            set
            {
                cbB_VOX = value;
            }
        }

        public int CbB_SQL
        {
            get
            {
                return cbB_SQL;
            }
            set
            {
                cbB_SQL = value;
            }
        }

        public int CbB_TOT
        {
            get
            {
                return cbB_TOT;
            }
            set
            {
                cbB_TOT = value;
            }
        }

        public string TB_CountsOfCH
        {
            get
            {
                return tB_CountsOfCH;
            }
            set
            {
                tB_CountsOfCH = value;
            }
        }

        public int CbB_CH_B_DisplayMode
        {
            get
            {
                return cbB_CH_B_DisplayMode;
            }
            set
            {
                cbB_CH_B_DisplayMode = value;
            }
        }

        public int CbB_CH_A_DisplayMode
        {
            get
            {
                return cbB_CH_A_DisplayMode;
            }
            set
            {
                cbB_CH_A_DisplayMode = value;
            }
        }

        public int CbB_SendIDDelay
        {
            get
            {
                return cbB_SendIDDelay;
            }
            set
            {
                cbB_SendIDDelay = value;
            }
        }

        public int CbB_PTTID
        {
            get
            {
                return cbB_PTTID;
            }
            set
            {
                cbB_PTTID = value;
            }
        }

        public int CbB_Scan
        {
            get
            {
                return cbB_Scan;
            }
            set
            {
                cbB_Scan = value;
            }
        }

        public int CbB_SaveMode
        {
            get
            {
                return cbB_SaveMode;
            }
            set
            {
                cbB_SaveMode = value;
            }
        }

        public int CbB_DTMF
        {
            get
            {
                return cbB_DTMF;
            }
            set
            {
                cbB_DTMF = value;
            }
        }

        public bool CB_SoundOfBi
        {
            get
            {
                return cB_SoundOfBi;
            }
            set
            {
                cB_SoundOfBi = value;
            }
        }

        public bool CB_StopSendOnBusy
        {
            get
            {
                return cB_StopSendOnBusy;
            }
            set
            {
                cB_StopSendOnBusy = value;
            }
        }

        public bool CB_AutoLock
        {
            get
            {
                return cB_AutoLock;
            }
            set
            {
                cB_AutoLock = value;
            }
        }

        public bool CB_LockKeyBoard
        {
            get
            {
                return cB_LockKeyBoard;
            }
            set
            {
                cB_LockKeyBoard = value;
            }
        }

        public string TB_A_RangeOfFreq
        {
            get
            {
                return tB_A_RangeOfFreq;
            }
            set
            {
                tB_A_RangeOfFreq = value;
            }
        }

        public string TB_A_CurFreq
        {
            get
            {
                return tB_A_CurFreq;
            }
            set
            {
                tB_A_CurFreq = value;
            }
        }

        public int CbB_A_SignalingEnCoder
        {
            get
            {
                return cbB_A_SignalingEnCoder;
            }
            set
            {
                cbB_A_SignalingEnCoder = value;
            }
        }

        public int CbB_A_RemainDir
        {
            get
            {
                return cbB_A_RemainDir;
            }
            set
            {
                cbB_A_RemainDir = value;
            }
        }

        public int CbB_A_FreqStep
        {
            get
            {
                return cbB_A_FreqStep;
            }
            set
            {
                cbB_A_FreqStep = value;
            }
        }

        public int CbB_A_CHBand
        {
            get
            {
                return cbB_A_CHBand;
            }
            set
            {
                cbB_A_CHBand = value;
            }
        }

        public string CbB_A_TxQT
        {
            get
            {
                return cbB_A_TxQT;
            }
            set
            {
                cbB_A_TxQT = value;
            }
        }

        public string CbB_A_RxQT
        {
            get
            {
                return cbB_A_RxQT;
            }
            set
            {
                cbB_A_RxQT = value;
            }
        }

        public int CbB_A_Power
        {
            get
            {
                return cbB_A_Power;
            }
            set
            {
                cbB_A_Power = value;
            }
        }

        public int CbB_A_Band
        {
            get
            {
                return cbB_A_Band;
            }
            set
            {
                cbB_A_Band = value;
            }
        }

        public string TB_B_RangeOfFreq
        {
            get
            {
                return tB_B_RangeOfFreq;
            }
            set
            {
                tB_B_RangeOfFreq = value;
            }
        }

        public string TB_B_CurFreq
        {
            get
            {
                return tB_B_CurFreq;
            }
            set
            {
                tB_B_CurFreq = value;
            }
        }

        public int CbB_B_SignalingEnCoder
        {
            get
            {
                return cbB_B_SignalingEnCoder;
            }
            set
            {
                cbB_B_SignalingEnCoder = value;
            }
        }

        public int CbB_B_RemainDir
        {
            get
            {
                return cbB_B_RemainDir;
            }
            set
            {
                cbB_B_RemainDir = value;
            }
        }

        public int CbB_B_FreqStep
        {
            get
            {
                return cbB_B_FreqStep;
            }
            set
            {
                cbB_B_FreqStep = value;
            }
        }

        public int CbB_B_CHBand
        {
            get
            {
                return cbB_B_CHBand;
            }
            set
            {
                cbB_B_CHBand = value;
            }
        }

        public string CbB_B_TxQT
        {
            get
            {
                return cbB_B_TxQT;
            }
            set
            {
                cbB_B_TxQT = value;
            }
        }

        public string CbB_B_RxQT
        {
            get
            {
                return cbB_B_RxQT;
            }
            set
            {
                cbB_B_RxQT = value;
            }
        }

        public int CbB_B_Power
        {
            get
            {
                return cbB_B_Power;
            }
            set
            {
                cbB_B_Power = value;
            }
        }

        public int CbB_B_Band
        {
            get
            {
                return cbB_B_Band;
            }
            set
            {
                cbB_B_Band = value;
            }
        }

        public int CbB_DisplayTypeOfPowerUp
        {
            get
            {
                return cbB_DisplayTypeOfPowerUp;
            }
            set
            {
                cbB_DisplayTypeOfPowerUp = value;
            }
        }

        public int CbB_PassRepetNoiseDetect
        {
            get
            {
                return cbB_PassRepetNoiseDetect;
            }
            set
            {
                cbB_PassRepetNoiseDetect = value;
            }
        }

        public int CbB_PassRepetNoiseClear
        {
            get
            {
                return cbB_PassRepetNoiseClear;
            }
            set
            {
                cbB_PassRepetNoiseClear = value;
            }
        }

        public int CbB_TailNoiseClear
        {
            get
            {
                return cbB_TailNoiseClear;
            }
            set
            {
                cbB_TailNoiseClear = value;
            }
        }

        public int CbB_TxUnderTDRStart
        {
            get
            {
                return cbB_TxUnderTDRStart;
            }
            set
            {
                cbB_TxUnderTDRStart = value;
            }
        }

        public int CbB_SoundOfTxEnd
        {
            get
            {
                return cbB_SoundOfTxEnd;
            }
            set
            {
                cbB_SoundOfTxEnd = value;
            }
        }

        public int CbB_AlarmMode
        {
            get
            {
                return cbB_AlarmMode;
            }
            set
            {
                cbB_AlarmMode = value;
            }
        }

        public bool CB_AlarmSound
        {
            get
            {
                return cB_AlarmSound;
            }
            set
            {
                cB_AlarmSound = value;
            }
        }

        public bool CB_FMRadioEnable
        {
            get
            {
                return cB_FMRadioEnable;
            }
            set
            {
                cB_FMRadioEnable = value;
            }
        }

        public bool CB_TDR
        {
            get
            {
                return cB_TDR;
            }
            set
            {
                cB_TDR = value;
            }
        }

        public int CbB_VoicSwitch
        {
            get
            {
                return cbB_VoicSwitch;
            }
            set
            {
                cbB_VoicSwitch = value;
            }
        }

        public int CbB_Language
        {
            get
            {
                return cbB_Language;
            }
            set
            {
                cbB_Language = value;
            }
        }

        public int CbB_KeySide
        {
            get
            {
                return cbB_KeySide;
            }
            set
            {
                cbB_KeySide = value;
            }
        }

        public int CbB_KeySideL
        {
            get
            {
                return cbB_KeySideL;
            }
            set
            {
                cbB_KeySideL = value;
            }
        }

        public int CbB_BackgroundColor
        {
            get
            {
                return cbB_BackgroundColor;
            }
            set
            {
                cbB_BackgroundColor = value;
            }
        }

        public int CbB_PowerOnMsg
        {
            get
            {
                return cbB_PowerOnMsg;
            }
            set
            {
                cbB_PowerOnMsg = value;
            }
        }

        public int CbB_VoxDelay
        {
            get
            {
                return cbB_VoxDelay;
            }
            set
            {
                cbB_VoxDelay = value;
            }
        }

        public int CbB_TimerMenuQuit
        {
            get
            {
                return cbB_TimerMenuQuit;
            }
            set
            {
                cbB_TimerMenuQuit = value;
            }
        }

        public int CbB_MicGain
        {
            get
            {
                return cbB_MicGain;
            }
            set
            {
                cbB_MicGain = value;
            }
        }

        public int CbB_A_Fhss
        {
            get
            {
                return cbB_A_Fhss;
            }
            set
            {
                cbB_A_Fhss = value;
            }
        }

        public int CbB_B_Fhss
        {
            get
            {
                return cbB_B_Fhss;
            }
            set
            {
                cbB_B_Fhss = value;
            }
        }

        public string TB_B_RemainFreq
        {
            get
            {
                return tB_B_RemainFreq;
            }
            set
            {
                tB_B_RemainFreq = value;
            }
        }

        public string TB_A_RemainFreq
        {
            get
            {
                return tB_A_RemainFreq;
            }
            set
            {
                tB_A_RemainFreq = value;
            }
        }

        public int CbB_WorkModeA
        {
            get
            {
                return cbB_WorkModeA;
            }
            set
            {
                cbB_WorkModeA = value;
            }
        }

        public int CbB_WorkModeB
        {
            get
            {
                return cbB_WorkModeB;
            }
            set
            {
                cbB_WorkModeB = value;
            }
        }

        public int CbB_1750Hz
        {
            get
            {
                return cbB_1750Hz;
            }
            set
            {
                cbB_1750Hz = value;
            }
        }
    }
    [Serializable]
    public class OtherImfData
    {
        private int theRangeOfVHF = 0;

        private string theMinFreqOfVHF = "136";

        private string theMaxFreqOfVHF = "174";

        private int theRangeOfUHF = 0;

        private string theMinFreqOfUHF = "400";

        private string theMaxFreqOfUHF = "520";

        private bool enableTxUHF = true;

        private bool enableTxVHF = true;

        private string powerUpChar1 = "BAOFENG";

        private string powerUpChar2 = "UV-5R";

        private bool enableTxOver480M = false;

        public string TheMinFreqOfVHF
        {
            get
            {
                return theMinFreqOfVHF;
            }
            set
            {
                theMinFreqOfVHF = value;
            }
        }

        public string TheMaxFreqOfVHF
        {
            get
            {
                return theMaxFreqOfVHF;
            }
            set
            {
                theMaxFreqOfVHF = value;
            }
        }

        public int TheRangeOfUHF
        {
            get
            {
                return theRangeOfUHF;
            }
            set
            {
                theRangeOfUHF = value;
            }
        }

        public string TheMinFreqOfUHF
        {
            get
            {
                return theMinFreqOfUHF;
            }
            set
            {
                theMinFreqOfUHF = value;
            }
        }

        public string TheMaxFreqOfUHF
        {
            get
            {
                return theMaxFreqOfUHF;
            }
            set
            {
                theMaxFreqOfUHF = value;
            }
        }

        public bool EnableTxUHF
        {
            get
            {
                return enableTxUHF;
            }
            set
            {
                enableTxUHF = value;
            }
        }

        public bool EnableTxVHF
        {
            get
            {
                return enableTxVHF;
            }
            set
            {
                enableTxVHF = value;
            }
        }

        public string PowerUpChar1
        {
            get
            {
                return powerUpChar1;
            }
            set
            {
                powerUpChar1 = value;
            }
        }

        public string PowerUpChar2
        {
            get
            {
                return powerUpChar2;
            }
            set
            {
                powerUpChar2 = value;
            }
        }

        public bool EnableTxOver480M
        {
            get
            {
                return enableTxOver480M;
            }
            set
            {
                enableTxOver480M = value;
            }
        }

        public int TheRangeOfVHF
        {
            get
            {
                return theRangeOfVHF;
            }
            set
            {
                theRangeOfVHF = value;
            }
        }
    }
}