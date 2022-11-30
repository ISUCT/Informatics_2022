def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def frange(start, stop, step):
    i = start
    while i < stop:
        yield i
        i += step


def y(x):
    y = ((((x - a) ** 2) ** (1.0 / 3)) + (abs(x + b) ** (1.0 / 5))) / ((x ** 2 - (a + b) ** 2) ** (1.0 / 9))
    return y


for x in frange(1.23, 7.23, 1.2):
    a = 0.8
    b = 0.4
    print(y(x))
# B
xrange = [1.88, 2.26, 3.84, 4.55, -6.21]
lenx = len(xrange)
for i in range(0, lenx):
    x = xrange[i]
    a = 0.8
    b = 0.4
    print(y(x))
