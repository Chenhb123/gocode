# print("Hello World")
# a = "xyz"
# b = a
# a = "qwe"
# print(a, b)
# aa = 1
# bb = aa
# aa = 2
# print(aa, bb)
# print('包含中文的str')
#
# print('hello %s, you sex is %s, and in fact is %s' % ('tom', 'man', True))
per = ["tom", "marry", "jack"]
print(per[0], per[1], per[2])
print(per[-1])
print(per[len(per)-1])
per.append("andy")
print(per)
per.insert(1, "monky")
print(per)
per.pop()
print(per)
per.pop(1)
print(per)
per[0]='tom2'
print(per)
