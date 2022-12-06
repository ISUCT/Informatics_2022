import math
#task1
def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"



#task2
def count_sheeps(sheep):
    return sheep.count(True)


#task3
def monkey_count(n):
    result = []
    for num in range(1, n + 1):
        result.append(num)
    return result



#task4
def paperwork(n, m):
    return n * m if n > 0 and m > 0 else 0 



#task5
def hero(b, d):
  bulletsneeded = d * 2
  if b >= bulletsneeded:
    return True
  else:
    return False


if even_or_odd == "main":
    even_or_odd(23,33,22)
if count_sheeps == "main":
    count_sheeps([True,False,True,False])
if monkey_count == "main":
    monkey_count(33)
if paperwork == "main":
    paperwork(33,12)
if hero == "main":
    hero(12,89)