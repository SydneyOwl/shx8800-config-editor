namespace SQ5R
{
    [Serializable]
    public class OtherImfData
    {
        public int theRangeOfVHF = 0;

        public string theMinFreqOfVHF = "136";

        public string theMaxFreqOfVHF = "174";

        public int theRangeOfUHF = 0;

        public string theMinFreqOfUHF = "400";

        public string theMaxFreqOfUHF = "520";

        public bool enableTxUHF = true;

        public bool enableTxVHF = true;

        public string powerUpChar1 = "BAOFENG";

        public string powerUpChar2 = "UV-5R";

        public bool enableTxOver480M = false;
    }
}