def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))

# лабораторная работа на вычисление функций
import math

# задача 1
print('Задача 1:')
x = 0.11
while x <= 0.36:
    print(math.asin(x ** 2) + math.acos(x ** 3))
    x += 0.05

# задача 2
print('Задача 2:')
x2 = [0.08, 0.26, 0.35, 0.41, 0.53]
for n in range(len(x2)):
    print(math.asin(x2[n] ** 2) + math.acos(x2[n] ** 3))
