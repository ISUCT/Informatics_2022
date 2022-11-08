using System;

namespace Informatics_2022
{
    class Count_LAB 
    {
        static void Main()
        {
            Console.Write("Введите, ответ на какую задачу хотите получить (A или B):");
            string zadacha = Console.ReadLine();
            switch (zadacha)
            {
                case "A":
                    double[] x_massA = new double[] { 2.5f, 1.28f, 3.28f, 0.4f };
                    foreach (double x in x_massA)
                    {
                        double b_cube = Math.Pow(2.5f, 3);
                        double x_cube = Math.Pow(x, 3);
                        double jopa1 = Math.Pow(Math.Sin(b_cube + x_cube), 2);
                        double jopa2 = 1 + Math.Pow((x_cube + b_cube), 1 / 3);
                        Console.WriteLine(jopa1 / jopa2);
                    }
                    Console.ReadLine();
                    break;
                case "B":
                    double[] x_massB = new double[] { 1.1f, 2.4f, 3.6f, 1.7f, 3.9f };
                    foreach (double x in x_massB)
                    {
                        double b_cube = Math.Pow(2.5f, 3);
                        double x_cube = Math.Pow(x, 3);
                        double jopa1 = Math.Pow(Math.Sin(b_cube + x_cube), 2);
                        double jopa2 = 1 + Math.Pow((x_cube + b_cube), 1 / 3);
                        Console.WriteLine(jopa1 / jopa2);
                    }
                    Console.ReadLine();
                    break;

            }
        }
    }
}