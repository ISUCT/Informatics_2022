
# задача A


def y(x):
    a = 0.8
    b = 0.4
    y = ((((x - a) ** 2) ** (1.0 / 3)) + (abs(x + b) ** (1.0 / 5))) / ((x ** 2 - (a + b) ** 2) ** (1.0 / 9))
    return y


x = 1.23
while x < 7.23:
    print(y(x))
    x += 1.2


# задача B
xrange = [1.88, 2.26, 3.84, 4.55, -6.21]
xlen = len(xrange)
for i in range(0, xlen):
    x = xrange[i]
    print(y(x))