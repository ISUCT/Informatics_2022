import math


def summ(a: int, b: int) -> int:
    return a + b


def lab_var_13_1(a: float, b: float, x_start: float, x_end: float, step: float):
    y = None
    while x_start <= x_end:
        try:
            y = (a * (x_start ** 0.5) - b * math.log(x_start, 5)) / (math.log10(abs(x_start - 1)))
        except Exception as Error:
            print(Error, x_start)
        x_start += step
        print(y)


def lab_var_13_2(app: list[...], a: float, b: float):
    for number in app:
        y = (a * (number ** 0.5) - b * math.log(number, 5)) / (math.log10(abs(number - 1)))
        print(y)


if __name__ == "__main__":
    lab_var_13_1(4.1, 2.7, 1.2, 5.2, 0.8)
    print("-" * 30)
    lab_var_13_2([1.9, 2.15, 2.34, 2.73, 3.16], 4.1, 2.7)
