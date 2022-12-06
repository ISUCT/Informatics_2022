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



#task3
list = []


def find_all(list, x):
    l = len(list)
    result = []
    for i in range(0,l):
        if x == list[i]:
            result.append(i)
    return result

if polish_letters == "main":
    polish_letters()
if sum_of_minimums == "main":
    sum_of_minimums([5,46,35,3],[2,3,754],[234,432,4567,6543])
if find_all == "main":
    find_all([234,5432,24,34,24],24)
