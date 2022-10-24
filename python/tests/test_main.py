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


if __name__ == '__main__':
    unittest.main()
