import math
a = 2


def y(x):
    y = ((math.log10(a + x)) ** 2) / (a + x) ** 2
    return y


x = 1.2
while x < 4.2:
    print('Task A', y(x))
    x += 0.6


xl = [1.16, 1.32, 1.47, 1.65, 1.93]
lenx = len(xl)
for i in range(0, lenx):
    x = xl[i]
    print('Task B', y(x))
