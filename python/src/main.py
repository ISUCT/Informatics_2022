from codewar import *
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


print(even_or_odd(2))
print(count_sheeps([True, True, True, False,
                    True, True, True, True,
                    True, False, True, False,
                    True, False, False, True,
                    True, True, True, True,
                    False, False, True, True]))
print(monkey_count(5))
print(paperwork(10, 10))
print(hero(20, 5))
print(correct_polish_letters("Jędrzej Błądziński"))
print(sum_of_minimums([[7, 9, 8, 6, 2], [6, 3, 5, 4, 3], [5, 8, 7, 4, 5]]))
print(find_all([6, 9, 3, 4, 3, 82, 11], 3))
