namespace CourseApp
{
    using System.Linq;

    public class Codewars
    {
        public static int[] FindAll(int[] a, int n) => Enumerable.Range(0, a.Length).Where(i => a[i] == n).ToArray();

        public static int SumOfMinimums(int[,] numbers)
        {
            var result = 0;
            for (int i = 0; i < numbers.GetLength(0); i++)
            {
                var min = numbers[i, 0];
                for (int j = 1; j < numbers.GetLength(1); j++)
                {
                    if (numbers[i, j] < min)
                    {
                        min = numbers[i, j];
                    }
                }

                result += min;
            }

            return result;
        }

        public static string PolishToEnglish(string sentence)
        {
            char[] result = sentence.ToCharArray();
            char[] polish = new char[] { 'ą', 'ć', 'ę', 'ł', 'ń', 'ó', 'ś', 'ź', 'ż' };
            char[] english = new char[] { 'a', 'c', 'e', 'l', 'n', 'o', 's', 'z', 'z' };
            for (byte i = 0; i < result.Length; i++)
            {
                for (byte j = 0; j < polish.Length; j++)
                {
                    if (result[i] == polish[j])
                    {
                        result[i] = english[j];
                    }
                }
            }

            return result.ToString();
        }
    }
}