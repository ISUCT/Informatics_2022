# https://www.codewars.com/kata/59a9919107157a45220000e1

list = []


def find_all(list, x):
    l = len(list)
    result = []
    for i in range(0, l):
        if x == list[i]:
            result.append(i)
    return result


print(find_all([6, 9, 3, 4, 3, 82, 11], 3))
