#task1
def polish_letters(s):
    polish = {"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z", "ż": "z"}
    return "".join([polish[c] if c in polish else c for c in s])

#task2
def sum_of_minimums(numbers):
    result = 0
    for i in numbers:
        result += min(i)
    return result
#or
def sof(n):
    return(sum(map(min, n)))
print(sof([[4,5,2],[8,33,1],[87,323,16]]))


#task3
list = []


def find_all(list, x):
    l = len(list)
    result = []
    for i in range(0,l):
        if x == list[i]:
            result.append(i)
    return result
print(find_all([3,8,3,5,7,8],3))