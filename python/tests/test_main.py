import unittest
from src import main


class LogTests(unittest.TestCase):

    def test_positive(self):
        res = main.Log(100, 10)
        self.assertEqual(0.5, res)


if __name__ == '__main__':
    unittest.main()
