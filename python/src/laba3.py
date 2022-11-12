import math


def lab_var_13_1(a: float, b: float, x_start: float, x_end: float, step: float):
    print("Les 1")
    y = None
    while x_start <= x_end:
        try:
            y = (a * (x_start ** 0.5) - b * math.log(x_start, 5)) / (math.log10(abs(x_start - 1)))
        except Exception as Error:
            print(Error, x_start)
        x_start += step
        print(y)


def lab_var_13_2(app: list, a: float, b: float):
    print("Les 2")
    for number in app:
        y = (a * (number ** 0.5) - b * math.log(number, 5)) / (math.log10(abs(number - 1)))
        print(y)
