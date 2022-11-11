import math


def Problem(a: float, b: float, x: float):
    result = math.log(x * x - 1) / Log(5, (a * x * x - b))
    message = "a={a}, b={b} x = {x}, result={result}"
    print(message.format(a=a, b=b, x=x, result=result))


def Log(base: float, x: float):
    return math.log(x) / math.log(base)


if __name__ == "__main__":
    a = 1.1
    b = 0.09

    print("Задача А\n")
    i = 1.2
    while i < 2.2:
        Problem(a, b, i)
        i += 0.2

    print("-------------------------------\n")

    print("Задача B\n")
    Problem(a, b, 1.21)
    Problem(a, b, 1.76)
    Problem(a, b, 2.53)
    Problem(a, b, 3.48)
    Problem(a, b, 4.52)
