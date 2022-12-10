def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


def EvenOrOdd(number):
    x = number
    if x % 2 == 0:
        return 'Even'
    else:
        return 'Odd'


def count_sheeps(sheep):
    a = 0
    for i in sheep:
        if i is True:
            a += 1


def monkey_count(n):
    return [x for x in range(1, n + 1)]


def paperwork(n, m):
    if n > 0 and m > 0:
        return n * m
    return 0


def hero(bullets, dragons):
    return bullets // dragons >= 2 if dragons != 0 else True


def correct_polish_letters(st):
    return st.translate(str.maketrans("ąćęłńóśźż", "acelnoszz"))


def find_all(array, n):
    res = []
    for i in range(len(array)):
        if array[i] == n:
            res.append(i)
    return res


def sum_of_minimums(numbers):
    return sum([min(x) for x in numbers])
