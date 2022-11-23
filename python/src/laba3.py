import math
from itertools import islice, count
from typing import Union


def calculate_function(values_list: list, a: float, b: float) -> list:
    result_list = []
    for value in values_list:
        if value - 1 != 0 and value - 1 != 1 and value > 0:
            result_list.append(round(((a * (value ** 0.5) - b * math.log(value, 5)) / math.log10(abs(value - 1))), 2))
        else:
            result_list.append(f"Incorrect value when x = {value}")
    return result_list


def lb3_main(a: float, b: float, x_start: float = None, x_end: float = None, x_step: float = None,
             values_list: list = None) -> Union[list, str]:
    if x_start and x_end and x_step:
        if x_end <= x_step:
            return calculate_function([x_start], a, b)
        return calculate_function(list(islice(count(x_start, x_step),
                                              int(abs(x_end - x_start) / x_step))), a, b)
    elif values_list:
        return calculate_function(values_list, a, b)
    return "Optional parameters are missing"
