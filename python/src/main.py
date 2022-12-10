def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))

def testA(x_start, x_finish, x_delta, x, a, b):
    
    # 1 zadanie
    def EvenOrOdd(n):
        if n % 2 == 0:
            return 'Even'
    return 'Odd'

    # 2 zadanie
def count_sheeps(sheep):
    a = 0
    for i in sheep:
        if i == True:
            a += 1

    # 3 zadanie
def monkey_count(n):
    return [x for x in range(1, n+1)]

#4 zadanie
def paperwork(n, m):
    return n * m if n > 0 and m > 0 else 0

#5 zadanie
    bulletsNeeded = dragons * 2;
    if bullets >= bulletsNeeded:
        return True;
    else:
        return False;   

#6 zadanie
def correct_polish_letters(st):
    return st.translate(str.maketrans("ąćęłńóśźż", "acelnoszz"))

#7 zadanie
def find_all(array, n):
    res = []
    for i in range(len(array)):
        if array[i] == n:
            res.append(i)   
    return res

#8 zadanie
def sum_of_minimums(numbers):
    return sum([min(x) for x in numbers])
    return a