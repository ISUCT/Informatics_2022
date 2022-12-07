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
        return m * n


def hero(bullets: int, dragons: int):
    if bullets // dragons >= 2:
        return True
    else:
        return False


def correct_polish_letters(st):
    st = st.translate(str.maketrans('ąćęłńóśźż', 'acelnoszz'))
    return (st)


def find_all(array: list, n: int):
    c = []
    for i in range(len(array)):
        if array[i] == n:
            c.append(i)
    return c


def sum_of_minimums(numbers):
    c = sum(min(x) for x in numbers)
    return c


print('1)', even_or_odd(4))
print('2)', count_sheeps([True, True, True, False, True, True, True, True]))
print('3)', monkey_count(10))
print('4)', paperwork(10, 3))
print('5)', hero(10, 3))
print('6)', correct_polish_letters("Lech Wałęsa"))
print('7)', find_all([6, 9, 3, 4, 3, 82, 11], 3))
print('8)', sum_of_minimums([[7, 9, 8, 6, 2], [6, 3, 5, 4, 3], [5, 8, 7, 4, 5]]))