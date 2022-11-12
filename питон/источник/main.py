import math
def  frange(start,stop,step):

    i=start
    while i < stop:
        yield i 
        i += step
def y(x):
    y=(a**(x**2)-1)-math.log10((x**2)-1)+pow((x**2)-1,1/3)
    return y
for x in frange (1.2,3.7,0.5):
    a =1.6
    print(y(x))



xrange=[1.28,1.36,2.47,3.68,4.56]
xlen =len(xrange)
for i in range(0,xlen):
    x=xrange[i]
    a = 1.6
    print(y(x))
