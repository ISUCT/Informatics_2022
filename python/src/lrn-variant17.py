import math

a = 0.1
b = 0.5
x = 0.15

print('  ')
print('Задание А:')
print('  ') 

while x <= 1.37:
    print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))
    x += 0.25

print('  ')
print('Задание B:')
print('  ')

x = 0.2
print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))

x = 0.3
print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))
x = 0.44
print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))

x = 0.6
print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))
x = 0.56
print(( a + (math.tan (b * x)) ** 2) / (b + (1 / math.tan (a * x)) ** 2))
