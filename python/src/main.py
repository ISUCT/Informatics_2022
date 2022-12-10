import math


def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def taskA(x_start=0.11, x_end=0.36, delta_x=0.05, x=0.11):
    answer: list = []
    while x_start <= x <= x_end:
        answer.append(schet(x))
        x += delta_x

    return answer


def taskB(x: list = [0.2, 0.3, 0.38, 0.43, 0.57]):
    answer: list = []
    for i in range(len(x)):
        answer.append(schet(x))
    return answer


def schet(x):
    return ((math.sin(x) ** 3) + (math.cos(x) ** 3)) * math.log10(x)
