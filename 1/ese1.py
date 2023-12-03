import os
import sys
import re

# read from file every line


def find_all_digit_in_string(string):
    return re.findall("\d", string)


with open("./input_1.txt", "r") as file:
    lines = file.readlines()

    numbers = []
    for line in lines:
        digits = find_all_digit_in_string(line)
        first_digit = digits[0]
        last_digit = digits[-1]

        final_number = int(first_digit + last_digit)
        print(first_digit, last_digit, final_number)
        numbers.append(final_number)

    print(sum(numbers))
