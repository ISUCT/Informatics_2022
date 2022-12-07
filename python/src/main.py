
import math
x_start = 1.25
x_end = 3.25
x_step = 0.4

while x_start <= x_end:
    y = abs(x_start ** 2 - 2.5) ** (1/4) + (math.log10(x_start*x_start)) ** (1/3)
    x_start += x_step
    print(y)




