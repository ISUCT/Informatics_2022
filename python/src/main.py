from codewars import find_all, even_or_odd, correct_polish_letters, count_sheeps, monkey_count, hero, paperwork, sum_of_minimums
import math


def solution(x, a, b):
    first = b ** 3 + (math.sin(a * x) ** 2)
    second = math.acos(b) + math.e ** ((-x) / 2)
    third = first / second
    print(third)


def task_a():
    a = 1.2
    b = 0.48
    xn = 0.7
    xk = 2.2
    while xn <= xk:
        solution(xn, a, b)
        xn += 0.3
    return solution(xn, a, b)


def task_b():
    a = 0.25
    b = 0.36
    xn = 0.56
    xk = 0.94
    while xn <= xk:
        solution(xn, a, b)
        xn += 1.28
    return solution(xn, a, b)


task_a()
task_b()


print(even_or_odd(2))
print(count_sheeps([True, True, True, False,
                    True, True, True, True,
                    True, False, True, False,
                    True, False, False, True,
                    True, True, True, True,
                    False, False, True, True]))
print(monkey_count(5))
print(paperwork(10, 10))
print(hero(20, 5))
print(correct_polish_letters("Jędrzej Błądziński"))
print(sum_of_minimums([[7, 9, 8, 6, 2], [6, 3, 5, 4, 3], [5, 8, 7, 4, 5]]))
print(find_all([6, 9, 3, 4, 3, 82, 11], 3))
