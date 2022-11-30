
def even_or_odd(number: int):
    return 'Even' if number % 2 == 0 else 'Odd'


def count_sheeps(sheep: list):
    return sheep.count(True)


def monkey_count(n: int):
    return list(i for i in range(1, n + 1))


def paperwork(n, m):
    if n > 0 and m > 0:
        return n * m
    return 0


def hero(bullets, dragons):
    return bullets // dragons >= 2 if dragons != 0 else True


def correct_polish_letters(st):
    d = {
        'ą': "a",
        'ć': 'c',
        'ę': 'e',
        'ł': 'l',
        'ń': 'n',
        'ó': 'o',
        'ś': 's',
        'ź': 'z',
        'ż': 'z'
    }
    for i in st:
        try:
            st = st.replace(i, d[f'{i}'])
        except:
            ...
    return st


def find_all(array, n):
    return [i for i in range(0, len(array)) if array[i] == n]


def sum_of_minimums(numbers):
    return sum([min(i) for i in numbers])


if __name__ == "__main__":
    print(even_or_odd(12))
    print(count_sheeps([True, True, False, True]))
    print(monkey_count(10))
    print(paperwork(5, 5))
    print(hero(78, 5))
    print(correct_polish_letters("Lech Wałęsa"))
    print(find_all([6, 9, 3, 4, 3, 82, 11], 3))
    print(sum_of_minimums([[1, 2, 3, 4, 5], [5, 6, 7, 8, 9], [20, 21, 34, 56, 100]]))
