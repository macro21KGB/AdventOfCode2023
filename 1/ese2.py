import os
import sys
import re

# read from file every line

possible_string_digit = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def find_all_digit_in_string(string):
    pattern = "one|two|three|four|five|six|seven|eight|nine|\d"
    all_digits = re.findall(pattern, string)

    return list(all_digits)


def transform_string_to_digit(string: str) -> str:
    if string.isdigit():
        return string
    return possible_string_digit[string]


def main():
    with open("./input_1.txt", "r") as file:
        lines = file.readlines()

        number_sum = 0
        for line in lines:
            digits = find_all_digit_in_string(line)
            first_digit = transform_string_to_digit(digits[0])
            last_digit = transform_string_to_digit(digits[-1])

            print(digits)
            print(first_digit, last_digit)
            joined_number = int(first_digit + last_digit)
            number_sum = joined_number + number_sum

    print(number_sum)




def test():
    test_string = "eightfourzkdxgqn8"
    expected = ["8", "4", "8"]
    actual = find_all_digit_in_string(test_string)
    print(actual)
    print(expected)


if __name__ == "__main__":
    main()
    # test()
