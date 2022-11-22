namespace CourseApp.Tests
{
    using System;
    using Xunit;

    public class DemoTest
    {
        [Fact]
        public void Test1()
        {
            {
                double xCube = Math.Pow(1.1, 3);
                double bCube = Math.Pow(2.5, 3);
                double znamen = 1 + Math.Pow(bCube + xCube, 1 / 3);
                Assert.NotEqual(0, znamen);
            }
        }

        [Fact]
        public void Test2()
        {
            {
                double xCube = Math.Pow(2.4, 3);
                double bCube = Math.Pow(2.5, 3);
                double znamen = 1 + Math.Pow(bCube + xCube, 1 / 3);
                Assert.NotEqual(0, znamen);
            }
        }

        [Fact]
        public void Test3()
        {
            double xCube = Math.Pow(3.6, 3);
            double bCube = Math.Pow(2.5, 3);
            double znamen = 1 + Math.Pow(bCube + xCube, 1 / 3);
            Assert.NotEqual(0, znamen);
        }

        [Fact]
        public void Test4()
        {
            double xCube = Math.Pow(1.7, 3);
            double bCube = Math.Pow(2.5, 3);
            double znamen = 1 + Math.Pow(bCube + xCube, 1 / 3);
            Assert.NotEqual(0, znamen);
        }

        [Fact]
        public void Test5()
        {
            double xCube = Math.Pow(3.9, 3);
            double bCube = Math.Pow(2.5, 3);
            double znamen = 1 + Math.Pow(bCube + xCube, 1 / 3);
            Assert.NotEqual(0, znamen);
        }
    }
}
