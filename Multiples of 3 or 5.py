def solution(number):
    return sum(x for x in range(1, number) if not x % 5 or not x % 3)
