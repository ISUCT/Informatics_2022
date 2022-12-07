import math


def summ(a: int, b: int) -> int:
    return a + b


# lab 1
def function1(t):
    return math.asin(t ** 2) + math.acos(t ** 3)


def task1(xst, xend, xstep):
    answer = []
    while xst <= xend:
        answer.append(function1(xst))
        xst += xstep
    return answer


def task2(xlist):
    answer = []
    for n in xlist:
        answer.append(function1(n))
    return answer


# codewars
def even_or_odd(number):
    if abs(number) % 2 == 0:
        return "Even"
    else:
        return "Odd"


def count_sheeps(sheep):
    count = 0
    for n in range(len(sheep)):
        if sheep[n] is True:
            count += 1
    return count


def monkey_count(n):
    answer = []
    while n != 0:
        answer.append(n)
        n -= 1
    return answer[::-1]


def paperwork(n, m):
    if n < 0 or m < 0:
        return 0
    else:
        return n * m


def hero(bullets, dragons):
    if 2 * dragons <= bullets:
        return True
    else:
        return False
    

def correct_polish_letters(st):
    polish = 'ąćęłńóśźż'
    normal = 'acelnoszz'
    for n in range(len(polish)):
        st = st.replace(polish[n], normal[n])
    return st


def find_all(array, n):
    answer = []
    for x in range(len(array)):
        if array[x] == n:
            answer.append(x)
    return answer


def sum_of_minimums(numbers):
    answer = 0
    for n in numbers:
        answer += min(n)
    return answer


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))
    print(task1(0.11, 0.36, 0.05))
    print(task2([0.08, 0.26, 0.35, 0.41, 0.53]))
