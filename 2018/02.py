import sys
from collections import Counter

def categorize(id):
    frequencies = Counter(id)
    result = [0, 0]
    for count in frequencies.values():
        if count == 2:
            result[0] = 1
        if count == 3:
            result[1] = 1
    return result

def part_one():
    two = 0
    three = 0
    for id in sys.stdin:
        cat = categorize(id)
        two = two + cat[0]
        three = three + cat[1]
    print two * three

def diff(a, b):
    if len(a) != len(b):
        return False
    differences = 0
    for i in range(len(a)):
        if a[i] != b[i]:
            differences = differences + 1
    if differences > 1:
        return False
    result = ""
    for i in range(len(a)):
        if a[i] == b[i]:
            result = result + a[i]

    print result

def part_two():
    ids = []
    for id in sys.stdin:
        ids.append(id.rstrip())

    for i in range(len(ids)):
        for j in range(i+1, len(ids)):
            diff(ids[i], ids[j])

part_two()