import math
a=0.8
b=0.4
x=1.23
while x<=7.23:
    y = (math.cbrt((x-a)**2)+(abs(x+b))**0.2)/(x**2-(a+b)**2)**(1/9)
    x+=1.2
    print(y)