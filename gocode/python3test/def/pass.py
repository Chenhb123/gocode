# 命名关键字
def person(name, age, **kw):
    if 'city' in kw:
        pass
    if 'job' in kw:
        pass
    print('name:', name, 'age:', age, 'other:', kw)


person('Jack', 24, city='beijing', job='teacher', zipcode='123456')

# 标准写法
def person2(name, age, *, city, job):
    print('name:', name, 'age:', age, city, job)


person2('Tom', 20, city='HangZhou', job='student')
# person2('Tom', 20, 'WH', 'stu')

#跟在可变参数后面
def person3(name, age, *args, city, job):
    print(name, age, args, city, job)


person3('Tom', 20, city='HangZhou', job='student')

#命名关键字默认值
def person4(name, age, *, city='北京', job='店长'):
    print(name, age, city, job)

# 关键字参数
# def person(name, age, **kw):
#     print('name:', name, 'age:', age, 'other:', kw)
#     for w in kw:
#         print(kw[w])


# person('tom', 12)
# person('andy', 15, city='北京')
# kw = {'city': 1, 'person': 2}
# person('mall', 16, **kw)

# 可变参数
# def calc(*numbers):
#     sum = 0
#     for n in numbers:
#         sum += n * n
#     return sum


# nums = [1, 2, 3]
# print(calc(*nums))
# print(calc(1, 2, 3, 4, 5))

# def add_end(l=[]):
#     l.append('END')
#     print(l)


# def add_end(l=None):
#     if l is None:
#         l = []
#     l.append('END')
#     print(l)


# add_end()
# add_end()
# add_end()

# 默认参数
# def enroll(name, gender, age=6, city='北京'):
#     print(name)
#     print(gender)
#     print(age, city)


# enroll('小米', "f")

# def power(x, n=2):
#     sum = 1
#     while n > 0:
#         n -= 1
#         sum *= x
#     return sum


# sum = power(3)
# print(sum)

# import math


# def move(x, y, step, angle=0):
#     nx = x+step*math.cos(angle)
#     ny = y-step*math.sin(angle)
#     return nx, ny


# nx, ny = move(100, 100, 60, math.pi/6)
# print(nx, ny)
# xy = move(100, 100, 60, math.pi/6)
# print(xy)

# def nop():
#     pass
# print("111")


# def my_abs(x):
#     if not isinstance(x, (int, float)):
#         raise TypeError("bad operand type")
#     if x >= 0:
#         return x
#     else:
#         return -x


# my_abs('A')
