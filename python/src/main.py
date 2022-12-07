import math


def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def laba_v_17_1(a: float, b: float, x_start: float, x_end: float, step: float):
    y = None
    while x_start <= x_end:
        try:
            y = (a + (math.tan(b * x_start) ** 2)) / (b + ((math.sin(a * x_start)) / (math.cos(a * x_start))) ** 2)
        except Exception as Error:
            print(Error, x_start)
        x_start += step
        print(y)


def laba_v_17_2(app: list, a: float, b: float):
    for x in app:
        y = (a + (math.tan(b * x) ** 2)) / (b + ((math.sin(a * x)) / (math.cos(a * x))) ** 2)
        print(y)


if "__main__" != __name__:
    pass
else:
    print('1')
    laba_v_17_1(0.1, 0.5, 0.15, 1.37, 0.25)

    print('2')

    laba_v_17_2([0.2, 0.3, 0.44, 0.6, 0.56], 0.1, 0.5)
