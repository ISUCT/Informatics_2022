import math

i = 0


def frange(start, stop, step):
    i = start
    while i < stop:
        yield i
        i += step


def y(x):
    y = ((math.log(a + x)) / ((a + x) ** 2))
    return y


for x in frange(1.2, 4.2, 0.6):
    a = 2.0
    print(y(x))
print()
xrange = [1.16, 1.32, 1.47, 1.65, 1.93]
xlen = len(xrange)
for i in range(0, xlen):
    x = xrange[i]
    a = 2.0
    print(y(x))