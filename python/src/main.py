def summ(a: int, b: int) -> int:
    return a + b


def sum_of_minimums(numbers):
    result = 0
    for i in range(len(numbers)):
        minimum = None
        for j in range(len(numbers[0])):
            if (minimum is None or minimum > numbers[i][j]):
                minimum = numbers[i][j]
        result += minimum
    return result


def find_all(array, n):
    result = []
    for idx, element in enumerate(array):
        if element == n:
            result.append(idx)
    return result


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


def paperwork(n, m):
    if n < 0 or m < 0:
        return 0
    else:
        return n * m


def monkey_count(n):
    count = []
    for i in range(n):
        count.append(i + 1)
    return count


def hero(bullets, dragons):
    if bullets < dragons * 2:
        return False
    else:
        return True


def even_or_odd(number):
    if number % 2 == 0:
        return "Even"
    else:
        return "Odd"


def count_sheeps(sheep):
    if sheep is not None:
        answer = 0
        for item in sheep:
            if item:
                answer += 1
        return answer
    return 0


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))
