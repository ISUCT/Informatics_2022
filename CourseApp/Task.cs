namespace CourceApp
{
    using System;
    using System.Collections.Generic;

    public class Task
    {
        public static double[] GetZnamen
        {
            get
            {
                return GetZnamen;
            }

            private set
            {
                GetZnamen[0] = Znamen(1.1);
                GetZnamen[1] = Znamen(2.4);
                GetZnamen[2] = Znamen(3.6);
                GetZnamen[3] = Znamen(1.7);
                GetZnamen[4] = Znamen(3.9);
            }
        }

        public static double Alg(double x, double b = 2.5)
        {
            double b_cube = Math.Pow(b, 3);
            double x_cube = Math.Pow(x, 3);
            double verx = Math.Pow(Math.Sin(b_cube + x_cube), 2);
            double niz = Znamen(x);
            return verx / niz;
        }

        public static List<double> TaskA(double startX, double finX, double deltaX)
        {
            List<double> doubles = new List<double>();
            for (double x = startX; x <= finX; x += deltaX)
            {
                doubles.Add(Alg(x));
            }

            return doubles;
        }

        public static List<double> TaskB(double[] mass)
        {
            List<double> doubles = new List<double>();
            foreach (double x in mass)
            {
                doubles.Add(Alg(x));
            }

            return doubles;
        }

        private static double Znamen(double x, double b = 2.5)
        {
            double bCube = Math.Pow(b, 3);
            double xCube = Math.Pow(x, 3);
            var znamen = Math.Pow(bCube / xCube, 1 / 3);
            return znamen;
        }
    }
}
