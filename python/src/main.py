import math


def summ(a: int, b: int) -> int:
    return a + b


def function1(t):
    return math.asin(t ** 2) + math.acos(t ** 3)


def task1(xst, xend, xstep):
    answer = []
    while xst <= xend:
        answer.append(function1(xst))
        xst += xstep
    return answer


def task2(xlist):
    answer = []
    for n in xlist:
        answer.append(function1(n))
    return answer


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))
    print(task1(0.11, 0.36, 0.05))
    print(task2([0.08, 0.26, 0.35, 0.41, 0.53]))
