# 列表生成器
import os
l = [x * x for x in range(1, 11)]
print(l)
l = [x+x for x in range(1, 11) if x % 2 == 0]
print(l)

qe = [m+n for m in 'ABC' for n in 'DEF']
print(qe)

dir = [d for d in os.listdir('.')]
print(dir)

kw = {'A': 1, 'B': 2, 'C': 3}
kv = [k + '=' + str(v) for k, v in kw.items()]
print(kv)

low = ['hello', 'tea', 'coffee']
great = [x.upper() for x in low]
print(great)
