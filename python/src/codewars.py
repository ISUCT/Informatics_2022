def even_or_odd(number):
    if number % 2 == 0:
        number = "Even"
    else:
        number = "Odd"
    return number


def count_sheeps(sheep):
    return sheep.count(True)
    pass


def monkey_count(n):
    k = 0
    c: list = []
    while k < n:
        k += 1
        c.append(k)
    return c


def paperwork(n, m):
    if n > 0 and m > 0:
        return n * m
    return 0


def hero(bullets, dragons):
    if bullets // dragons >= 2:
        return True
    return False
