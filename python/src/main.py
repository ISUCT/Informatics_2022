import math
x = 0.08
a = float(2.0)
b = float(1.1)
while x <= 1.08:
  y = (math.log1p(a) * (abs(b ** 2 - x ** 2))) / (x ** 2 - a ** 2)**(1 / 5)
  x += 0.2
  print('A) ', y)

n = [0.1,0.3,0.4,0.45,0.65]
for i in range(len(n)):
  x = n[i]
  y = (math.log1p(a) * (abs(b ** 2 - x ** 2))) / (x ** 2 - a ** 2)**(1 / 5)
  print('B) ', y)