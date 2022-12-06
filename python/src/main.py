import math


def Problem(a: float, b: float, x: float):
    result = math.log(x * x - 1) / Log(5, (a * x * x - b))
    message = "a={a}, b={b} x = {x}, result={result}"
    print(message.format(a=a, b=b, x=x, result=result))


def Log(base: float, x: float):
    return math.log(x) / math.log(base)


# Sum of Minimums
def sum_of_minimums(numbers):
    result = 0
    for i in range(len(numbers)):
        minimum = None
        for j in range(len(numbers[0])):
            if (minimum is None or minimum > numbers[i][j]):
                minimum = numbers[i][j]
        result += minimum
    return result


# Find all occurrences of an element in an array
def find_all(array, n):
    result = []
    for idx, element in enumerate(array):
        if element == n:
            result.append(idx)
    return result


# Polish alphabet
def correct_polish_letters(st):
    st = st.replace('ą', 'a')
    st = st.replace('ć', 'c')
    st = st.replace('ę', 'e')
    st = st.replace('ł', 'l')
    st = st.replace('ń', 'n')
    st = st.replace('ó', 'o')
    st = st.replace('ś', 's')
    st = st.replace('ź', 'z')
    st = st.replace('ż', 'z')
    return st


# Beginner Series #1 School Paperwork
def paperwork(n, m):
    if n < 0 or m < 0:
        return 0
    else:
        return n * m


# Count the Monkeys!
def monkey_count(n):
    count = []
    for i in range(n):
        count.append(i + 1)
    return count


# Is he gonna survive?
def hero(bullets, dragons):
    if bullets < dragons * 2:
        return False
    else:
        return True


# Even or Odd
def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"


# Counting sheep
def count_sheeps(sheep):
    if sheep is not None:
        answer = 0
        for item in sheep:
            if item:
                answer += 1
        return answer
    return 0


if __name__ == "__main__":
    a = 1.1
    b = 0.09

    print("Задача А\n")
    i = 1.2
    while i < 2.2:
        Problem(a, b, i)
        i += 0.2

    print("-------------------------------\n")

    print("Задача B\n")
    Problem(a, b, 1.21)
    Problem(a, b, 1.76)
    Problem(a, b, 2.53)
    Problem(a, b, 3.48)
    Problem(a, b, 4.52)
