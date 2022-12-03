# Task1
def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"


even_or_odd(25)


# Task2
def count_sheeps(array_of_sheep):
    return array_of_sheep.count(True)


# Task3
def monkey_count(c):
    result = []
    for i in range(0, c):
        result.append(i)
    return result


monkey_count(25)


# Task4
def papework(a, b):
    if a > 0 and b > 0:
        return a*b
    else:
        return 0


papework(25, 11)


# Task5
def hero(x, y):
    bulletsneeded = y*2
    if x >= bulletsneeded:
        return "True"
    else:
        return "False"


hero(25, 11)
