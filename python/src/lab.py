# 3var
import math
a = 2.0
b = 0.95
xlist = [2.2, 3.78, 4.51, 6.58, 1.2]
xlen = len(xlist)
print('Задание А:')


def frange(start, stop, step):
    st = start
    while st < stop:
        yield st
        st += step


def y(x):
    y = (1 + math.pow(math.log10(x / a), 2)) / (b - math.exp(x / a))
    return y


for x in frange(1.25, 2.75, 0.3):
    print(y(x))


print('Задание B:')
for i in range(0, xlen):
    x = xlist[i]
    print(y(x))
