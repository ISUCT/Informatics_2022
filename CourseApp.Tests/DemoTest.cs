namespace CourseApp.Tests
{
    using CourceApp;
    using Xunit;

    public class DemoTest
    {
        [Fact]
        public void Test1()
        {
            {
                var t = Task.Znamen(1.1);
                Assert.NotEqual(0, t);
            }
        }

        [Fact]
        public void Test2()
        {
            {
                var t = Task.Znamen(2.4);
                Assert.NotEqual(0, t);
            }
        }

        [Fact]
        public void Test3()
        {
            var t = Task.Znamen(3.6);
            Assert.NotEqual(0, t);
        }

        [Fact]
        public void Test4()
        {
            var t = Task.Znamen(1.7);
            Assert.NotEqual(0, t);
        }

        [Fact]
        public void Test5()
        {
            var t = Task.Znamen(3.9);
            Assert.NotEqual(0, t);
        }
    }
}
