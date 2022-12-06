import math

#task1
def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"
even_or_odd(10)


#task2
def count_sheeps(sheep):
    return sheep.count(True)


#task3
def monkey_count(n):
    result = []
    for num in range(1, n + 1):
        result.append(num)
    return result
monkey_count(10)


#task4
def paperwork(n, m):
    return n * m if n > 0 and m > 0 else 0
paperwork(10,16)


#task5
def hero(b, d):
  bulletsneeded = d * 2
  if b >= bulletsneeded:
    return True
  else:
    return False
hero(30,10)

#task6
def correct_polish_letters(st):
    pol=pol = {"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z", "ż": "z"}
    return "".join([pol[c] if c in pol else c for c in st])

#task7
def find_all(array, n):
    arr = []
    for el in range(len(array)-1):
        if array[el] == n:
            arr.append(el)
    return arr

#task8
def sum_of_minimums(numbers):
    return sum(map(min, numbers))