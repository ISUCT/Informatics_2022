
def even_or_odd(number: int):
    return 'Even' if number % 2 == 0 else 'Odd'


def count_sheeps(sheep: list):
    return sheep.count(True)


def monkey_count(n: int):
    return list(i for i in range(1, n + 1))


def paperwork(n, m):
    return n * m if n > 0 and m > 0 else 0


def hero(bullets, dragons):
    return bullets // dragons >= 2 if dragons != 0 else True


if __name__ == "__main__":
    print(even_or_odd(12))
    print(count_sheeps([True, True, False, True]))
    print(monkey_count(10))
    print(paperwork(5, 5))
    print(hero(78, 5))
