import math

a = 1.6
x = 1.2
print('Задание А:')


def y(x):
    y = (a ** ((x ** 2) - 1)) \
        - (math.log10((x ** 2) - 1)) \
        + (((x ** 2) - 1) ** 1 / 3)
    return y


while x < 3.7:
    print(y(x))
    x += 0.5

print('Задание B:')
xlist = [1.28, 1.36, 2.47, 3.68, 4.56]
xlen = len(xlist)
for i in range(0, xlen):
    print(y(x))
    x = xlist[i]
