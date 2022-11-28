namespace CourseApp.Tests
{
    using CourceApp;
    using Xunit;

    public class DemoTest
    {
        [Fact]
        private void Test1()
        {
            var t = Task.Znamen(1.1);
            Assert.NotEqual(0, t);
        }

        [Fact]
        private void Test2()
        {
            var t = Task.Znamen(2.4);
            Assert.NotEqual(0, t);
        }

        [Fact]
        private void Test3()
        {
            var t = Task.Znamen(3.6);
            Assert.NotEqual(0, t);
        }

        [Fact]
        private void Test4()
        {
            var t = Task.Znamen(1.7);
            Assert.NotEqual(0, t);
        }

        [Fact]
        private void Test5()
        {
            var t = Task.Znamen(3.9);
            Assert.NotEqual(0, t);
        }

        [Fact]
        private void Test6()
        {
            var t = Task.TaskA(1.28, 3.28, 0.4);
            Assert.Equal(6, t.Count);
        }
    }
}
