import math


def calculate_function(a: float, b: float, x: float) -> float:
    try:
        y = (a * (x ** 0.5) - b * math.log(x, 5)) / math.log10(abs(x - 1))
    except:
        y = None
    return y


def lb3_values(a: float, b: float, values_list: list) -> list:
    if values_list:
        return [calculate_function(a, b, value) for value in values_list]
    else:
        raise Exception("Values list`s empty")


def lb3_start_end(a: float, b: float, x_start: float, x_end: float, x_step: float) -> list:
    data_arr = []
    if x_start == x_end == x_step:
        return [calculate_function(a, b, x_start)]
    while x_start <= x_end:
        data_arr.append(calculate_function(x_start, a, b))
        x_start += x_step
    return data_arr
