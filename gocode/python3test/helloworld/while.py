# sum = 0
# n = 99
# while n > 0:
#     sum += n
#     print(n)
#     n -= 2
# print(sum)

# n = 100
# while n > 0:
#     print(n)
#     n -= 1
#     if n < 90:
#         break
# print('END')

n = 10
while n > 0:
    n -= 1
    if n % 2 == 0:
        continue
    print(n)
