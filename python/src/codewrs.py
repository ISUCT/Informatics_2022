def even_or_odd(number):
    if number % 2 == 0:
        return "even"
    else:
        return "odd"


def count_sheep(sheep):
    count = 0
    for i in sheep:
        if i == True:
            count += 1
            return count


def count_monkey(n):
    res = []
    for i1 in range(1, n + 1):
        res.append(i1)
        return res


def paper(n, m):
    if n > 0 and m > 0:
        return n*m
    else:
        return 0


def hero(dragon, bullet):
    if bullet >= dragon*2:
        return True
    else:
        return False