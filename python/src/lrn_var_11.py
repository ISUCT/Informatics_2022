import math

print('Part A')

a = 2
b = 3
x = 0.11

while x <= 0.36:
    print(math.asin(x ** a) + math.acos(x ** b))
    x += 0.05

print('Part B')

x = 0.08
print(math.asin(x ** a) + math.acos(x ** b))

x = 0.26
print(math.asin(x ** a) + math.acos(x ** b))

x = 0.35
print(math.asin(x ** a) + math.acos(x ** b))

x = 0.41
print(math.asin(x ** a) + math.acos(x ** b))

x = 0.53
print(math.asin(x ** a) + math.acos(x ** b))