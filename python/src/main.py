import math


def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def Calculation(x: float, a: float, b: float):
    if x > 0 and x != 1:
        return ((a * x + b) ** (1 / 3)) / ((math.log(x, 10))**2)
    return None


def Aformula(x=2.1, a=1.35, b=0.98, x_start=1.14, x_finish=4.24, x_delta=0.62):
    list_reply: list = []
    if x_start > x or x > x_finish:
        list_reply.append(None)
    while x_start <= x <= x_finish:
        x += x_delta
        y = Calculation(x, a, b)
        list_reply.append(y)
    return list_reply


def Bformula(variables=[0.35, 1.28, 3.51, 5.21, 4.16], a=1.35, b=0.98):
    list_reply: list = []
    for i in range(len(variables)):
        x = variables[i]
        y = Calculation(x, a, b)
        list_reply.append(y)
    return list_reply
