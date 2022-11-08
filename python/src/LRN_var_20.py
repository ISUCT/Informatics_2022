import math
a = 0.8
b = 0.4
x = 1.23
print('Task A')
while x<=7.23:
    y = (math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9)
    x += 1.2
    print(y)

print('___________')
print('Task B')
x = 1.88
print((math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9))

x = 2.26
print((math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9))

x = 3.84
print((math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9))

x = 4.55
print((math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9))

x = -6.21
print((math.cbrt((x - a) ** 2) + (abs(x + b)) ** 0.2) / (x ** 2 - (a + b) ** 2) ** (1 / 9))
