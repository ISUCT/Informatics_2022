import math
xn = 1.25
a = 2
b = 0.95
print ("задача А")
while xn <= 2.75:
    xn = xn+0.3
    y = ((1+(math.log10(xn/a))**2)/(b-(math.e**xn/a)))
    print(y)
print ("задача В")
x1=[2.2,3.78,4.51,6.58,1.2]
for n in range(len(x1)):
    y2 = ((1 + (math.log10(x1[n] / a)) ** 2) / (b - (math.e ** x1[n] / a)))
    print (y2)