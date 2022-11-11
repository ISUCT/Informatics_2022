namespace Informatics_2022
{
    using System;

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
                        double b_cube = Math.Pow(2.5f, 3);
                        double x_cube = Math.Pow(x, 3);
                        double verx = Math.Pow(Math.Sin(b_cube + x_cube), 2);
                        double niz = 1 + Math.Pow(x_cube + b_cube, 1 / 3);
                        Console.WriteLine(verx / niz);
                    }

                    Console.ReadLine();
                    break;
                case "B":
                    double[] x_massB = new double[] { 1.1f, 2.4f, 3.6f, 1.7f, 3.9f };
                    foreach (double x in x_massB)
                    {
                        double b_cube = Math.Pow(2.5f, 3);
                        double x_cube = Math.Pow(x, 3);
                        double verx = Math.Pow(Math.Sin(b_cube + x_cube), 2);
                        double niz = 1 + Math.Pow(x_cube + b_cube, 1 / 3);
                        Console.WriteLine(verx / niz);
                    }

                    Console.ReadLine();
                    break;
            }
        }
    }
}