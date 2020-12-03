import collections
print(isinstance([], collections.Iterator))
print(isinstance(iter([]), collections.Iterator))
print(isinstance((k for k in range(3)), collections.Iterator))

# import collections
# print(isinstance([], collections.Iterable))
# print(isinstance({}, collections.Iterable))
# print(isinstance((), collections.Iterable))
# print(isinstance(set([1, 2, 1, 2]), collections.Iterable))
# print(isinstance([k for k in range(10)], collections.Iterable))
# print(isinstance((k for k in range(5)), collections.Iterable))
# print(isinstance('abc', collections.Iterable))
# print(isinstance(1234, collections.Iterable))
