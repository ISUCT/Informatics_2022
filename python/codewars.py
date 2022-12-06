# task 1
def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"


# task 2
def count_sheeps(sheep):
    return sheep.count(True)


# task 3
def monkey_count(n):
    result = []
    for num in range(1, n + 1):
        result.append(num)
    return result


# task 4
def paperwork(n, m):
    return n * m if n > 0 and m > 0 else 0


# task 5
def hero(b, d):
    bulletsneeded = d * 2
    if b >= bulletsneeded:
        return True
    else:
        return False


# task 1
def correct_polish_letters(st):
    pol = {"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z", "ż": "z"}
    return "".join([pol[i] if i in pol else i for i in st])


# task 2
def find_all(array, n):
    arr = []
    for i in range(len(array)):
        if array[i] == n:
            arr.append(i)
    return arr


# task 3
def sum_of_minimums(numbers):
    return sum(map(min, numbers))
