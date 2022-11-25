import math
from itertools import islice, count
from typing import Union


def calculate_function(values_list: list, a: float, b: float) -> list:
    result_list = []
    for value in values_list:
        try:
            result_list.append(round(((a * (value ** 0.5) - b * math.log(value, 5)) / math.log10(abs(value - 1))), 2))
        except:
            result_list.append(f"Incorrect value when x = {value}")
    return result_list


def lb3_main(a: float, b: float, values_list: list, gap: bool = False) -> Union[list, str]:
    if gap:
        if values_list[1] <= values_list[2]:
            return calculate_function([values_list[0]], a, b)
        return calculate_function(list(islice(count(values_list[0], values_list[2]),
                                              int(abs(values_list[1] - values_list[0]) / values_list[2]))), a, b)
    elif not gap:
        return calculate_function(values_list, a, b)
    return "Optional parameters are missing"
