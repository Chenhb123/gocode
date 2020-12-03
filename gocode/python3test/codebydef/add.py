# reduce
import functools


def add(x, y):
    return x+y


def cal(x, y):
    return x*10+y


l = list(range(5))
r = functools.reduce(add, l)
print(r)
r = functools.reduce(cal, l)
print(r)

# map
# def ff(x):
#     return x*x
# l = [x for x in range(10)]
# r = map(ff, l)
# print(list(r))
# s = map(str, l)
# print(list(s))
# def add(x, y, f):
#     return f(x) + f(y)


# print(add(-5, 6, abs))
