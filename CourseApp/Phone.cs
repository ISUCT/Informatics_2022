namespace CourseApp
{
    using System;

    public class Phone
    {
        private float diaonal;

        public Phone(string name, float diagonal)
        {
            Name = name;
            Diagonal = diagonal;
        }

        public string Name { get; set; }

        public float Diagonal
        {
            get
            {
                return diaonal;
            }

            set
            {
                if (value > 0 && value < 20)
                {
                    this.diaonal = value;
                }
            }
        }

        public void Show()
        {
            Console.WriteLine($"{Name} with diagonal {diaonal}");
        }
    }
}