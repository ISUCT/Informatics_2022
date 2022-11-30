import unittest
from src import main


class FormulaCalculationTestCase(unittest.TestCase):

    def test_x_one(self):
        result = main.lb3_value(1, 1, value_list=[3])
        self.assertEqual(result[0], 3.4861795442282557)

    def test_x_several(self):
        result = main.lb3_value(1, 1, value_list=[1, 2, 3])
        self.assertEqual(len(result), 3)

    def test_x_gap_several(self):
        result = main.lb3_gap(1, 1, x_start=1, x_end=5, x_step=1)
        self.assertEqual(len(result), 5)

    def test_x_negative(self):
        result = main.lb3_value(1, 1, value_list=[-1])
        sub = [i for i in result if "Inc" in i]
        self.assertEqual(len(sub), True)

    def test_x_zero(self):
        result = main.lb3_value(1, 1, value_list=[0])
        sub = [i for i in result if "Inc" in i]
        self.assertEqual(len(sub), True)

    def test_step_more_end(self):
        result = main.lb3_gap(1, 1, x_start=1, x_end=2, x_step=3)
        self.assertEqual(len(result), 1)

    def test_start_step_end_zero(self):
        result = main.lb3_gap(1, 1, x_start=0, x_end=0, x_step=0)
        self.assertEqual(len(result), True)

    def test_a_b_negative(self):
        result = main.lb3_gap(-1, -2, x_start=1, x_end=2, x_step=1)
        if len(result) > 0:
            self.assertTrue(True)
        else:
            self.fail()

    def test_values_empty(self):
        result = main.lb3_value(1, 1, value_list=[])
        self.assertEqual(result, 'Optional params are missing')


if __name__ == '__main__':
    unittest.main()
