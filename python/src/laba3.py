import math
from typing import Union, Any


def calculate_function(a: float, b: float, values_list: list) -> Union[str, list[Union[str, Any]]]:
    result_list = []
    if not values_list:
        return "Optional params are missing"
    for value in values_list:
        try:
            result_list.append(((a * (value ** 0.5) - b * math.log(value, 5)) / math.log10(abs(value - 1))))
        except:
            result_list.append(f"Incorrect value when x = {value}")
    return result_list


def lb3_value(a: float, b: float, value_list: list):
    return calculate_function(a, b, values_list=value_list)


def lb3_gap(a: float, b: float, x_start: float, x_end: float, x_step: float):
    data_arr = []
    if x_start == x_end == x_step:
        return calculate_function(a, b, values_list=[x_start])
    while x_start <= x_end:
        data_arr.append(x_start)
        x_start += x_step
    return calculate_function(a, b, values_list=data_arr)
