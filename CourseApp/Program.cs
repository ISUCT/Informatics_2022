namespace CourceApp
{
    using System;
    using CourseApp;

    internal class Program
    {
         private static void Main()
         {
            Console.WriteLine("A:");
            foreach (double elem in Task.TaskA(1.28, 3.28, 0.4))
            {
                Console.WriteLine(elem);
            }

            Console.WriteLine("B:");
            foreach (double elem in Task.TaskB(new double[] { 1.1, 2.4, 3.6, 1.7, 3.9 }))
            {
                Console.WriteLine(elem);
            }

            Console.WriteLine(Codewars.PolishToEnglish("Jędrzej Błądziński"));
         }
    }
}