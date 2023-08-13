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
                if (args[0] == "transfile")
                {
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

    

    

    
}