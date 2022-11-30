import math
a = 2.25
x = 1.2

print('Задание А')


def y(x):
    y = (a ** ((x ** 2) - 1)) - (math.log10((x ** 2) - 1)) + (((x ** 2) - 1) ** 1/3)
    return y


while x < 2.7:
    print(y(x))
    x += 0.3


print('Задание B')
xlist = [1.31, 1.39, 1.44, 1.56, 1.92]
xlen = len(xlist)
for i in range(0, xlen):
    x = xlist[i]
    print(y(x))
