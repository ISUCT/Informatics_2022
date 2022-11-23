import unittest
from src import main


class FormulaCalculationTestCase(unittest.TestCase):

    def test_x_two(self):
        result = main.lb3_main(1, 1, values_list=[2])
        sub = [i for i in result if "Inc" in i]
        self.assertEqual(len(sub), True)

    def test_x_negative(self):
        result = main.lb3_main(1, 1, values_list=[-1])
        sub = [i for i in result if "Inc" in i]
        self.assertEqual(len(sub), True)

    def test_x_zero(self):
        result = main.lb3_main(1, 1, values_list=[0])
        sub = [i for i in result if "Inc" in i]
        self.assertEqual(len(sub), True)

    def test_step_more_end(self):
        result = main.lb3_main(1, 1, x_start=1, x_end=2, x_step=3)
        self.assertEqual(len(result), 1)

    def test_start_step_end_zero(self):
        result = main.lb3_main(1, 1, x_start=0, x_end=0, x_step=0)
        self.assertEqual(result, "Optional parameters are missing")

    def test_no_x(self):
        result = main.lb3_main(1, 1)
        self.assertEqual(result, "Optional parameters are missing")


if __name__ == '__main__':
    unittest.main()
