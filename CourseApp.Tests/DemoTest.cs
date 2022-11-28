namespace CourseApp.Tests
{
    using CourceApp;
    using Xunit;

    public class DemoTest
    {
        /*
        [Fact]
        public void Test1()
        {
            var znamen1 = Task.GetZnamen[0];
            Assert.NotEqual(0, znamen1);
        }

        [Fact]
        public void Test2()
        {
            var znamen2 = Task.GetZnamen[1];
            Assert.NotEqual(0, znamen2);
        }

        [Fact]
        public void Test3()
        {
            var znamen3 = Task.GetZnamen[2];
            Assert.NotEqual(0, znamen3);
        }

        [Fact]
        public void Test4()
        {
            var znamen4 = Task.GetZnamen[3];
            Assert.NotEqual(0, znamen4);
        }

        [Fact]
        public void Test5()
        {
            var znamen5 = Task.GetZnamen[4];
            Assert.NotEqual(0, znamen5);
        }
        */
        [Fact]
        public void Test1()
        {
            var taskAMass = Task.TaskA(1.28, 3.28, 0.4);
            Assert.Equal(6, taskAMass.Count);
        }
    }
}
