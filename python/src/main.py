def frange(start, stop, step):
    i = start
    while i < stop:
        yield i
        i += step
import math 
def y(x):
    y = (b**3 + (math.sin(a*x))**2) / (math.acos(x*b*x) + e**(-x/2))
    return y
for x in frange(0.7,2.2,0.3):
    a = 1.2
    b = 0.48
    e = 2.7
    print(y(x))


xrange = [0.25, 0.36, 0.56, 0.94, 1.28]
xlen = len(xrange)
for  i in range(0,xlen):
    x = xrange[i]
    a = 1.2
    b = 0.48
    e = 2.7
    print(y(x))


