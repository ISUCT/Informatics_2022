namespace CourceApp
{
    using System;

    public class Task
    {
        public static double Alg(double x, double b = 2.5)
        {
            double b_cube = Math.Pow(b, 3);
            double x_cube = Math.Pow(x, 3);
            double verx = Math.Pow(Math.Sin(b_cube + x_cube), 2);
            double niz = 1 + Math.Pow(x_cube + b_cube, 1 / 3);
            Console.WriteLine(verx / niz);
            return verx / niz;
        }
    }
}
