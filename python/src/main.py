if __name__ == "__main__":
    import math
def f(x):
    b = 2.5
    return (1 + (math.sin(b ** 3 + x ** 3)) ** 2) / ((b ** 3 + x ** 3) ** 1/3)
xn = 1.28
while xn <= 3.28:
    print (f(xn))
    xn += 0.4
print (f(1.1))
print (f(2.4))
print (f(3.6))
print (f(1.7))
print (f(3.9))
