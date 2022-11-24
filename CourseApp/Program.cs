namespace CourceApp
{
    using System;

    internal class Program
    {
         private static void Main()
         {
            Console.WriteLine("A:");
            for (double x = 1.28; x <= 3.28; x += 0.4)
            {
                Console.WriteLine(Task.Alg(x));
            }

            Console.WriteLine("B:");

            double[] x_massB = new double[] { 1.1, 2.4, 3.6, 1.7, 3.9 };
            foreach (double x in x_massB)
            {
                Console.WriteLine(Task.Alg(x));
            }

            Console.ReadLine();
         }
    }
}