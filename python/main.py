import math

def f(x, a=0.05, b=0.06):
    y1 = x**2 - b**2
    y2 = x**2 - a**2
    if (y1<-1 or 1<y1) or (y2<-1 or 1<y2):
        return None
    else:
        y1 = math.acos(y1)
        y2 = math.asin(y2)
        if y2 == 0:
            return None
        else:
            return y1/y2

def a(xh, xk, kd, a=0.05, b=0.06):
    q = []
    while xh <= xk:
        xh += kd
        q.append(f(xd))
    return q
    
def b(x1, x2, x3, x4, x5, a=0.05, b=0.06):
#    q = []
#    q.append(f(x1))
#    q.append(f(x2))
#    q.append(f(x3))
#    q.append(f(x4))
#    q.append(f(x5))
    return [f(x1), f(x2), f(x3), f(x4), f(x5)]


#for i in range(-1000, 1000):
#    print(i, f(i))
#for i in range(-1000, 1000):
#    i /= 10
#    print(i, f(i))
#for i in range(-1000, 1000):
#    i /= 100
#    print(i, f(i))
#for i in range(-1000, 1000):
#    i /= 1000
#    print(i, f(i))

#def f(x, a=0.05, b=0.06):
#    y = math.acos(x**2 - b**2) / math.asin(x**2 - a**2)
#    return y

#for i in range(-100, 100):
#    i /= 100
#    print(i, math.acos(i))
