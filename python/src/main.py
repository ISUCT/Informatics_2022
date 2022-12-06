import math
from codewars import find_all, even_or_odd, correct_polish_letters, count_sheeps, monkey_count, hero, paperwork, sum_of_minimums
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


def frange(start, stop, step):

    i = start
    while i < stop:
        yield i
        i += step


def y(x):
    y = (a**(x**2) - 1) - math.log10((x**2) - 1) + pow((x ** 2) - 1, 1 / 3)
    return y


for x in frange(1.2, 3.7, 0.5):
    a = 1.6
    print(y(x))


xrange = [1.28, 1.36, 2.47, 3.68, 4.56]
xlen = len(xrange)
for i in range(0, xlen):
    x = xrange[i]
    a = 1.6
    print(y(x))
