import math

a = 4.1
b = 2.7

def Y(x):
    global a, b
    
    return (a * (x ** 1/3) - b * math.log(x, 5)) / (math.log(x - 1, 10) ** 3)
    

#  task №23 A
print("task №23 A")

x_first = 1.5
x_last = 3.5
step = 0.4

x = x_first
while x < x_last:
    print(f"y({x}) = {Y(x)}")
    x+=step


#  task №23 B
print("task №23 B")

array_x = [1.9, 2.15, 2.34, 2.74, 3.16]
for x in array_x:
    print(f"y({x}) = {Y(x)}")
