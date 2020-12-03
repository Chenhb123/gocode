import collections
print(isinstance(123, collections.Iterable))
d = {'a': 1, 'b': 2, 'c': 3}
print(isinstance(d, collections.Iterable))
l = list(range(10))
for k, v in enumerate(l):
    print(k, v)
# for k in d:
#     print(k)
# for v in d.values():
#     print(v)
# for k, v in d.items():
#     print(k, v)
