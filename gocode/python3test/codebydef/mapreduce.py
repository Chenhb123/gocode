# str to int
import functools

# lambda

DIGITS = {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
          '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}


def char2num(s):
    return DIGITS[s]


def str2int(s):
    return functools.reduce(lambda x, y: x * 10 + y, map(char2num, s))


print(str2int('01234'))

# DIGITS = {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
#           '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}


# def str2int(s):
#     def fn(x, y):
#         return x * 10 + y

#     def char2num(s):
#         return DIGITS[s]
#     return functools.reduce(fn, map(char2num, s))

# print(str2int('12345'))

# def fn(x, y):
#     return x * 10 + y


# def char2num(s):
#     digits = {'0': 0, '1': 1, '2': 2, '3': 3, '4': 4,
#               '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
#     return digits[s]

# print(functools.reduce(fn, map(char2num, '13579')))
