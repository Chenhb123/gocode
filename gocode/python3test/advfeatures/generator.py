def fib(max):
    n, a, b = 0, 0, 1
    while n < max:
        yield b
        a, b = b, a+b
        n += 1
    return 'Done'

fi = fib(6)
while True:
    try:
        x = next(fi)
        print('g:', x)
    except StopIteration as e:
        print('Gengerator return value:', e.value)
        break
# fi = fib(10)
# for f in fi:
#     print(f)
# g = (x + x for x in range(1, 6))
# for k in g:
#     print(k)
