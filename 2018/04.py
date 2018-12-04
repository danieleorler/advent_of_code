import re, sys
from datetime import datetime
patternTs = re.compile("\[(?P<date>[0-9]{4}-[0-9]{2}-[0-9]{2}\s[0-9]{2}:[0-9]{2})\].*")
patternHeader = re.compile(".*Guard\s#(?P<id>[0-9]+).*")
patternStatus = re.compile(".*(?P<status>falls|wakes).*")

guards = {}

def parseHeader(line):
    match = patternHeader.match(line)
    return int(match.group("id"))

def parseStatus(line):
    match = patternStatus.match(line)
    return match.group("status")

def analyzeShift(id, changes):
    if not id in guards:
        guards[id] = [0 for i in range(60)]
    for i in range(1, len(changes), 2):
       for j in range(changes[i-1][0].minute, changes[i][0].minute):
           guards[id][j] = guards[id][j] + 1

def findMostSleptMinute(freq):
    max = -1
    id = -1
    for i in range(0, len(freq)):
        if freq[i] > max:
            max = freq[i]
            id = i
    return (id, max)

def calculateFreq():
    records = []
    for line in sys.stdin:
        match = patternTs.match(line)
        ts = datetime.strptime(match.group("date"),"%Y-%m-%d %H:%M")
        if patternHeader.match(line):
            records.append((ts, parseHeader(line)))
        if patternStatus.match(line):
            records.append((ts, parseStatus(line)))

    records.sort(key=lambda x: x[0])
    changes = []
    currentId = 0
    for record in records:
        if type(record[1]) is int:
            analyzeShift(currentId, changes)
            currentId = record[1]
            changes = []
        else:
            changes.append(record)

def part_one():
    calculateFreq()
    minutesSleeping = -1
    mostSleptMinute = -1
    guard = -1
    for key in guards:
        ms = sum(guards[key])
        if(minutesSleeping < ms):
            minutesSleeping = ms
            mostSleptMinute = findMostSleptMinute(guards[key])[0]
            guard = key

    print guard, mostSleptMinute, minutesSleeping, guard * mostSleptMinute

def part_two():
    calculateFreq()
    rank = []
    for key in guards:
        m = findMostSleptMinute(guards[key])
        rank.append((key, m[0], m[1], key*m[0]))

    rank.sort(key=lambda x: x[1])
    for entry in rank:
        print entry

part_one()
print "------"
part_two()