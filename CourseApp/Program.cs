namespace Informatics_2022
{
    using System;
    using CourseApp;

    internal class Program
    {
         private static void Main()
         {
            Console.Write("Введите, ответ на какую задачу хотите получить (A или B):");
            string zadacha = Console.ReadLine();
            switch (zadacha)
            {
                case "A":
                    for (double x = 1.28; x <= 3.28; x += 0.4)
                    {
                        Task.Alg(x);
                    }

                    Console.ReadLine();
                    break;

                case "B":
                    double[] x_massB = new double[] { 1.1f, 2.4f, 3.6f, 1.7f, 3.9f };
                    foreach (double x in x_massB)
                    {
                        Task.Alg(x);
                    }

                    Console.ReadLine();
                    break;
            }
        }
    }
}