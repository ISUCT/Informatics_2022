namespace CourseApp
{
    using System;

    public class Program
    {
        public static void main(string[] args)
        {
            Phone phone1 = new Phone("iPhone", -7);
            phone1.Show();
            phone1.Diagonal = 7;
            phone1.Show();
            phone1.Diagonal = -16;
            phone1.Show();

            Phone tablet = new Phone("Android", 6);
            tablet.Diagonal = 6;
            tablet.Show();
            tablet.Diagonal = -10;
            tablet.Show();
            tablet.Diagonal = 8;
            tablet.Show();
            Console.WriteLine("Hello World");
        }
    }
}
