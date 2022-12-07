namespace CourseApp.Tests
{
    using CourceApp;
    using Xunit;

    public class DemoTest
    {
        [Fact]
        public void Test1()
        {
            var taskAMass = Task.TaskA(1.28, 3.28, 0.4);
            Assert.Equal(6, taskAMass.Count);
        }
    }
}
