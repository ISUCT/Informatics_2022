def even_or_odd(number: int):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"


def count_sheeps(sheep):
    sheep == list
    return sheep.count(True)


def monkey_count(monkey):
    c = list()
    for i in range(1, monkey + 1):
        c.append(i)
    return c


def paperwork(n: int, m: int):
    if n <= 0 or m <= 0:
        return 0
    else:
        return (m * n)


def hero(bullets: int, dragons: int):
    if (bullets // dragons >= 2):
        return True
    else:
        return False


print('1)', even_or_odd(4))
print('2)',count_sheeps([True, True, True, False, True, True, True, True]))
print('3)',monkey_count(10))
print('4)',paperwork(10, 3))
print('5)',hero(10, 3))
