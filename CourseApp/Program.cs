namespace CourceApp
{
    using System;

    internal class Program
    {
         private static void Main()
         {
            Console.WriteLine("A:");
            var t = Task.TaskA(1.28, 3.28, 0.4);
            foreach (double x in t)
            {
                Console.WriteLine(x);
            }

            Console.WriteLine("B:");

            var m = Task.TaskB(new double[] { 1.1, 2.4, 3.6, 1.7, 3.9 });
            foreach (double x in m)
            {
                Console.WriteLine(x);
            }

            Console.ReadLine();
         }
    }
}