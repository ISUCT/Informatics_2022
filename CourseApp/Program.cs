namespace CourceApp
{
    using System;

    internal class Program
    {
         private static void Main()
         {
            Console.WriteLine("A:");
            var massA = Task.TaskA(1.28, 3.28, 0.4);
            foreach (double answer in massA)
            {
                Console.WriteLine(answer);
            }

            Console.WriteLine("B:");

            var massB = Task.TaskB(new double[] { 1.1, 2.4, 3.6, 1.7, 3.9 });
            foreach (double answer in massB)
            {
                Console.WriteLine(answer);
            }
         }
    }
}