from numpy import arange


# A


def y(x):
    y = ((((x - a) ** 2) ** (1.0 / 3)) + (abs(x + b) ** (1.0 / 5))) / ((x ** 2 - (a + b) ** 2) ** (1.0 / 9))
    return y


for x in arange(1.23, 7.23, 1.2):
    a = 0.8
    b = 0.4
    print(y(x))

# B
xrange = [1.88, 2.26, 3.84, 4.55, -6.21]
xlen = len(xrange)
for i in range(0, xlen):
    x = xrange[i]
    a = 0.8
    b = 0.4
    print(y(x))
