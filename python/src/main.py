import math


def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def Aformula(x: float = 2.1, a: float = 1.35, b: float = 0.98):
    list_reply = []
    if 1.14 > x or x > 4.24:
        list_reply.append(None)
    while 1.14 <= x <= 4.24:
        x += 0.62
        y = ((a * x + b) ** (1 / 3)) / ((math.log(x, 10))**2)
        list_reply.append(y)
    return list_reply


def Bformula(variables=[0.35, 1.28, 3.51, 5.21, 4.16], a=1.35, b=0.98):
    list_reply = []
    for i in range(len(variables)):
        x = variables[i]
        if x != 1 and x > 0:
            y = ((a * x + b) ** (1 / 3)) / ((math.log(x, 10))**2)
            list_reply.append(y)
        else:
            list_reply.append(None)
    return list_reply
