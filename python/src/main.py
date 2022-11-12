import math

def summ(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    print("Hello world")
    print(summ(3, 4))


x = float(1.14)
a = float(1.35)
b = float(0.98)
while 1.14 <= x <= 4.24:
    z = (math.log(x, 10))**2
    y = (( a * x + b) ** ( 1 / 3 ) ) / z
    x += 0.62
    print('A) ', y)

n = [0.35, 1.28, 3.51, 5.21, 4.16]
a = float(1.35)
b = float(0.98)
for i in range(len(n)):
    x = n[i]
    z = (math.log(x,10))**2
    y = (( a * x + b) ** ( 1 / 3 ) ) / z
    print('B) ', y)