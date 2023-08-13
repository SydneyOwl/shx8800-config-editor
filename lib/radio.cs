using System.Runtime.Serialization.Formatters.Binary;

namespace SQ5R
{
#pragma warning disable SYSLIB0011
    [Serializable]
    public class ClassTheRadioData
    {
        public string[][] channelData = new string[128][];

        public FormFunCFGData funCfgData = new FormFunCFGData();

        public DTMFData dtmfData = new DTMFData();

        public OtherImfData otherImfData = new OtherImfData();
        public static ClassTheRadioData CreatObjFromFile(Stream s)
        {
            BinaryFormatter binaryFormatter = new BinaryFormatter();
            return binaryFormatter.Deserialize(s) as ClassTheRadioData;
        }
    }
}