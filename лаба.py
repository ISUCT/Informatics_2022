import math
#A
x = 0.11
while 0.11<= x <=0.36:
    y= ((math.sin(x)**3)+(math.cos(x)**3))*math.log10(x)
    x += 0.05
    print('A)', y)
#B
n= [0.2,0.3,0.38,0.43,0.57]
for i in range(len(n)):
    x=n[i]
    y=((math.sin(x)**3)+(math.cos(x)**3))*math.log10(x)
    print('B)',y)
