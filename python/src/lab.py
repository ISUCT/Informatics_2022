import math

def y(a,b,x): 
    a = 0.1
    b = 0.5
    y=(a + (math.tan(b * x)) ** 2) / (b + (1/math.tan(a * x)) ** 2)
    return y


#  task №17 A
print("task №17 A")
def task_a(a,b,x_first,x_last,step):
    x_first = 0.15
    x_last = 1.37
    step = 0.25
    x = x_first
    while x < x_last:
        print(f"x = {x} y = {y(x)}")
        x += step


#  task №17 B
print("task №17 B")
def task_b(a, b, array_x = []):
    array_x = [0.2, 0.3, 0.44, 0.6, 0.56]
    for x in array_x:
        print(f"x = {x} y = {y(x)}")

if task_a == "main":
    task_a(12,46,28,75,3)
if task_b == "main":
    task_b(34,654,2)

    
