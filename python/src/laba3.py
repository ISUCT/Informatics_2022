import math


def calculate_function(a: float, b: float, x: float) -> float:
    return (a * (x ** 0.5) - b * math.log(x, 5)) / math.log10(abs(x - 1))


def lb3_value(a: float, b: float, x: float):
    return calculate_function(a, b, x)


def lb3_values(a: float, b: float, value_list: list):
    return main(a, b, values_list=value_list) if value_list else None


def lb3_gap(a: float, b: float, x_start: float, x_end: float, x_step: float):
    data_arr = []
    if x_start == x_end == x_step:
        return main(a, b, values_list=[x_start])
    while x_start <= x_end:
        data_arr.append(x_start)
        x_start += x_step
    return main(a, b, values_list=data_arr)


def main(a: float, b: float, values_list: list) -> list:
    result_list = []
    for value in values_list:
        try:
            result_list.append(calculate_function(a, b, value))
        except Exception as Error:
            print(Error)
    return result_list
