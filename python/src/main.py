import math


def lab_var_22_1(a: float, x_s: float, x_e: float, step: float):
    y = None
    while x_s <= x_e:
        try:
            y = a ** (x_s ** 2 - 1) - math.log((x_s ** 2 - 1), 10) + (x_s ** 2 - 1) ** (1 / 3)
        except Exception as Error:
            print(Error, x_s)
        x_s += step
        print(y)


def lab_var_22_2(app: list, a: float):
    for num in app:
        y = a ** (num ** 2 - 1) - math.log((num ** 2 - 1), 10) + (num ** 2 - 1) ** (1 / 3)
        print(y)


if "__main__" != __name__:
    pass
else:
    print('zadanie 1')
    lab_var_22_1(2.25, 1.2, 2.7, 0.3)

    print('zadanie 2')

    lab_var_22_2([4.48, 3.56, 2.78, 5.28, 3.21], 2.25)