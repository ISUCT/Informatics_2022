import math
a = 2.0
b = 3.0
x = 0.11


def y(x):
    y = (math.asin(x ** a)) + math.acos(x ** b)
    return y


print("Задание А")


while x < 0.36:
    print(y(x))
    x += 0.05


print("Задание Б")
m = [0.08, 0.25, 0.35, 0.41, 0.53]
u = len(m)
for i in range(0, u):
    x = m[i]
    print(y(x))
