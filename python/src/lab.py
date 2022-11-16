# 3var
import math
a = 2.0
b = 0.95
x = 1.25
print('Задание А:')


def y(x):
    y = (1 + math.pow(math.log10(x / a), 2)) / (b - math.exp(x / a))
    return y


while x < 2.75:
    print(y(x))
    x += 0.3


# for x in task_a(1.25, 2.75, 0.3):
#     print(y(x))


print('Задание B:')
xlist = [2.2, 3.78, 4.51, 6.58, 1.2]
xlen = len(xlist)
for i in range(0, xlen):
    x = xlist[i]
    print(y(x))
