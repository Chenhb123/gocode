# 尾递归
def fact(n):
    return fact_rec(n, 1)


def fact_rec(num, product):
    if num == 1:
        return product
    return fact_rec(num-1, num*product)


print(fact(5))
# def fact(n):
#     if n == 1:
#         return n
#     return n * fact(n-1)


# print(fact(5))
