namespace CourceApp
{
    using System;

    internal class Program
    {
         private static void Main()
         {
            Console.WriteLine("A:");
            Console.WriteLine(Task.TaskA(1.28, 3.28, 0.4));

            Console.WriteLine("B:");
            Console.WriteLine(Task.TaskB(new double[] { 1.1, 2.4, 3.6, 1.7, 3.9 }));

            Console.ReadLine();
         }
    }
}