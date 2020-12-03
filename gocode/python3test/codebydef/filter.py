def is_odd(n):
    return n % 2 == 1


print(list(filter(is_odd, [1, 2, 3, 4, 5, 6, 7, 8, 9])))


def not_empty(s):
    return s and s.strip()


print(list(filter(not_empty, ['', "", None, 'A', '  ', 'B', '\n', 'C'])))
