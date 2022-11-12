import math

b = 2.5
x = 1.28
while x <= 3.28:
    y = 1 + ( math.sin( b ** 3 + x ** 3 ) ) ** 2 / ( b **3 + x ** 3) ** 1/3
    print(y)
    x += 0.4
x = 1.1
print (( 1 + (math.sin (b ** 3 + x ** 3)) ** 2) / ((b ** 3 + x ** 3) ** (1 / 3)))

x = 2.4
print ( 1+ (math.sin (b ** 3 + x ** 3)) ** 2 / ((b ** 3 + x ** 3) ** (1 / 3)))
x = 3.6
print (( 1 + (math.sin (b ** 3 + x ** 3)) ** 2) / ((b ** 3 + x ** 3) ** (1 / 3)))
x = 1.7
print ( 1 + (math.sin (b ** 3 + x ** 3)) ** 2) / ((b ** 3 + x ** 3) ** (1 / 3))
x = 3.9
print ( 1 + (math.sin (b ** 3 + x ** 3)) ** 2) / ((b ** 3 + x ** 3) ** (1 / 3))

