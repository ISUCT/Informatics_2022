import math

x = 1.14
print('Задание А:')


# Задание А

def A(x, a=1.35, b=0.98):
    y = (((a * x + b) ** (1 / 3)) / ((math.log(x, math.e)) ** 2))
    return y


if __name__ == '__main__':
    while x < 4.24:
        print(A(x))
        x += 0.62


# Задание B


def B(x):
    y = (((1.35 * x + 0.98) ** (1 / 3)) / ((math.log(x, math.e)) ** 2))
    return y


print('Задание B:')
xlist = [0.35, 1.28, 3.51, 5.21, 4.16]
xlen = len(xlist)
for i in range(0, xlen):
    x = xlist[i]
    print(B(x))

