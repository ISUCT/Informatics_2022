import unittest
from src import main


class FormulaCalculationTestCase(unittest.TestCase):

    def test_x_one(self):
        result = main.lb3_values(1, 1, [3])
        self.assertEqual(result[0], 3.4861795442282557)

    def test_x_several(self):
        result = main.lb3_values(1, 1, values_list=[10, 20, 30])
        self.assertEqual(len(result), 3)

    def test_x_negative(self):
        result = main.lb3_values(1, 1, values_list=[-1])
        self.assertEqual(None in result, True)

    def test_x_zero(self):
        result = main.lb3_values(1, 1, values_list=[0])
        self.assertEqual(None in result, True)

    def test_values_empty(self):
        try:
            main.lb3_values(1, 1, values_list=[])
        except:
            self.assertTrue(True)
            return
        self.fail()

    def test_x_gap_several(self):
        result = main.lb3_start_end(1, 1, x_start=5, x_end=9, x_step=1)
        self.assertEqual(len(result), 5)

    def test_start_step_end_zero(self):
        result = main.lb3_start_end(1, 1, x_start=0, x_end=0, x_step=0)
        self.assertEqual(None in result, True)

    def test_a_b_negative(self):
        result = main.lb3_start_end(-1, -2, x_start=5, x_end=9, x_step=1)
        self.assertEqual(len(result), 5)


if __name__ == '__main__':
    unittest.main()
