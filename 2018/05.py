import re, sys

def findToReplace(s):
    to_replace = set()
    for i in range(1, len(s)):
        if abs(ord(s[i-1]) - ord(s[i])) == 32:
            to_replace.add(s[i-1]+s[i])
    return to_replace

def reduce(polymer):
    to_replace = findToReplace(polymer)
    while len(to_replace) > 0:
        for replacement in to_replace:
            polymer = polymer.replace(replacement, "")
        to_replace = findToReplace(polymer)

    return polymer

def part_one():
    print len(reduce(sys.stdin.read()))

def part_two():
    polymer = sys.stdin.read()
    result = []
    for i in range(ord("a"),ord("z")+1):
        tmp = polymer.replace(chr(i), "")
        tmp = tmp.replace(chr(i-32), "")
        l = len(reduce(tmp))
        result.append(l)

    print "result", min(result)

part_one()
print "----"
part_two()
