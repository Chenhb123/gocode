# 函数参数组合
def f2(a, b, c=1, *args, **kw):
    print(a, b, c, args, kw)


# def f2(a, b, c=1, *, d, e, **kw):
#     print(a, b, c, d, e, kw)


args1 = (1, 2, 3, 4)
kw1 = {'A': 1, 'B': 2, 'C': 3}
f2(*args1, **kw1)
# f1(1, 2)
# f1(1, 2, 3)
# f2(1, 2, 3, d=4, e=5)


# def f1(a, b, c, d=0, *args, **kw):
#     print('a =', a, 'b =', b, 'c =', c, 'd=', d, 'args =', args, 'kw =', kw)


# args = (1, 2, 3)
# kw0 = {'d': 99}
# kw1 = {'dd': 99}
# f1(*args, **kw0)
# f1(*args, **kw1)
