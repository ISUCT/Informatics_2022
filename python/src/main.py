def summ(a: int, b:int) -> int:
    return a + b


if __name__ == "_main_":
    print("Hello World")
    print(summ(3, 4))


import math
# задача под А

def y(x):
    y = ((math.log(a + x))/((a + x) ** 2))
    return y


x = 1.2
X = 4.2
while x < X:
    a = 2.0
    print(y(x))
    x += 0.6

# задача под B
A = [1.16, 1.32, 1.47, 1.65, 1.93]
a = 2.0
for x in A:
    print (y(x))