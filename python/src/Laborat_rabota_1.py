import math




def y(x):
    k: float = (math.asin(x) ** 2 + math.acos(x) ** 4) ** 3
    return k


x = 0.26
while x < 0.66:
    print(y(x))
    x += 0.08


xrange = [0.1, 0.35, 0.4, 0.55, 0.6]
xlen = len(xrange)
for s in range(0, xlen):
    x = xrange[s]
    print(y(x))
