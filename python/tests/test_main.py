import unittest
from src import main


class SummTests(unittest.TestCase):

    def test_positive(self):
        res = main.summ(2, 3)
        self.assertEqual(5, res)

    def test_zero(self):
        res = main.summ(0, 0)
        self.assertEqual(0, res)

    def test_one_negative(self):
        res = main.summ(-2, 3)
        self.assertEqual(1, res)

    def test_both_negative(self):
        res = main.summ(-2, -4)
        self.assertEqual(-6, res)

    def test_one_negative_zero_res(self):
        res = main.summ(-2, 2)
        self.assertEqual(0, res)

    def test_one_negative_and_text(self):
        try:
            main.summ(-2, "2")
        except:
            self.assertTrue(True)
            return
        self.fail()


class FormulaTestsA(unittest.TestCase):
    def test_A(self):
        res = main.Aformula(3.1)
        self.assertAlmostEqual(5.583, res[0], places=3)

    def test_A_zero(self):
        res = main.Aformula(0)
        self.assertEqual([None], res)

    def test_A_one(self):
        res = main.Aformula(1)
        self.assertEqual([None], res)

    def test_A_negative(self):
        res = main.Aformula(-434.222)
        self.assertEqual([None], res)


class FormulaTestsB(unittest.TestCase):

    def test_B(self):
        res = main.Bformula([2.2, 1.2], 4, 2.12)
        self.assertAlmostEqual(18.921, res[0], places=3)

    def test_B_zero(self):
        res = main.Bformula([0])
        self.assertEqual([None], res)

    def test_B_one(self):
        res = main.Bformula([1])
        self.assertEqual([None], res)

    def test_B_negative(self):
        res = main.Bformula([-434.222])
        self.assertEqual([None], res)


if __name__ == '__main__':
    unittest.main()
