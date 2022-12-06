# Evenor Odd
def even_or_odd(number):
    if (number % 2) != 0:
        return "Odd"
    return "Even"


# Counting sheep...
def count_sheeps(sheep):
    return sheep.count(True)


# Count the Monkeys!
def monkey_count(n):
    result = []
    for num in range(1, n + 1):
        result.append(num)
    return result


# Beginner Series #1 School Paperwork
def paperwork(n, m):
    if n < 0 or m < 0:
        return 0
    else:
        return n * m


# Is he gonna survive?
def hero(bullets, dragons):
    return bullets >= (2 * dragons)


# Polish alphabet
def correct_polish_letters(st):
    pol = pol = {"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z", "ż": "z"}
    return "".join([pol[c] if c in pol else c for c in st])


# Sum of Minimums!
def sum_of_minimums(numbers):
    return sum(map(min, numbers))


# Find all occurrences of an element in an array
def find_all(array, n):
    indices = []
    for i in range(len(array)):
        if array[i] == n:
            indices.append(i)
    return indices
