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
    pattern = "|".join(possible_string_digit.keys()) + "|\\d"
    all_digits = re.findall(pattern, string)
    mapped_digits = map(transform_string_to_digit, all_digits)

    return list(mapped_digits)


def transform_string_to_digit(string):
    if string.isdigit():
        return string
    return possible_string_digit[string]


def main():
    with open("./input_1.txt", "r") as file:
        lines = file.readlines()

        numbers = []
        for line in lines:
            digits = find_all_digit_in_string(line)

            first_digit = digits[0]
            last_digit = digits[-1]

            final_digit = int(first_digit + last_digit)
            print(first_digit, last_digit, final_digit)
            numbers.append(final_digit)
        print(sum(numbers))

        with open("./output_1.txt", "w") as output_file:
            for number in numbers:
                output_file.write(str(number) + "\n")


def test():
    test_string = "eightfourzkdxgqn8"
    expected = ["8", "4", "8"]
    actual = find_all_digit_in_string(test_string)
    print(actual)
    print(expected)


if __name__ == "__main__":
    main()
    # test()
