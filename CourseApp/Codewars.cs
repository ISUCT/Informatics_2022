namespace CourseApp
{
    using System.Linq;

    public class Codewars
    {
        public static string EvenOrOdd(int number) // чётное или нечётное
        {
            return number % 2 == 0 ? "Even" : "Odd";
        }

        public static int CountSheeps(bool[] sheeps) // овцы
        {
            int count = 0;
            foreach (bool sheep in sheeps)
            {
                if (sheep == true)
                {
                    count++;
                }
            }

            return count;
        }

        public static int[] MonkeyCount(int n) // почситать обезьян
        {
            int elem = 1;
            int[] mass = new int[n];
            for (byte index = 0; index < mass.Length; index++)
            {
                if (elem <= n)
                {
                    mass[index] = elem++;
                }
            }

            return mass;
        }

        public static int Paperwork(int n, int m) // посчитать количество страниц
        {
            int result = 0;
            if (n > 0 && m > 0)
            {
                result = m * n;
            }

            return result;
        }

        public static bool Hero(int bullets, int dragons) => bullets / dragons >= 2; // драконы

        public static int[] FindAll(int[] a, int n) => Enumerable.Range(0, a.Length).Where(i => a[i] == n).ToArray(); // индексы вхожддений

        public static int SumOfMinimums(int[,] numbers) // сумма минимумов
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

        public static string PolishToEnglish(string sentence) // перевод польский на русский
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

            return new string(result);
        }
    }
}