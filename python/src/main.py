import math


def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def taskA(x_start=0.11, x_end=0.36, delta_x=0.05, x=0.11):
    answer: list = []
    while x_start <= x <= x_end:
<<<<<<< HEAD
        answer.append(schet(x))
=======
        y = ((math.sin(x) ** 3) + (math.cos(x) ** 3)) * math.log10(x)
        answer.append(y)
>>>>>>> dac6f410932518146ec1a0cf5b8c422fc5aff3a6
        x += delta_x

    return answer


<<<<<<< HEAD
def taskB(x: list = [0.2, 0.3, 0.38, 0.43, 0.57]):
    answer: list = []
    for i in range(len(x)):
        answer.append(schet(x))
    return answer


def schet(x):
    return ((math.sin(x) ** 3) + (math.cos(x) ** 3)) * math.log10(x)
=======
def taskB(arguments: list = [0.2, 0.3, 0.38, 0.43, 0.57]):
    answer: list = []
    for i in range(len(arguments)):
        x = arguments[i]
        y = ((math.sin(x) ** 3) + (math.cos(x) ** 3)) * math.log10(x)
        answer.append(y)
    return answer
>>>>>>> dac6f410932518146ec1a0cf5b8c422fc5aff3a6
