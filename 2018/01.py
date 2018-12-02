import sys

def part_one():
    freq = 0
    for line in sys.stdin:
        freq = freq + int(line)
    print freq

def part_two():
    freq = 0
    frequencies = set([0])

    changes = []
    for line in sys.stdin:
        changes.append(int(line))

    while True:
        print "looping..."
        for change in changes:
            freq = freq + change
            if freq in frequencies:
                return freq
            else:
                frequencies.add(freq)
    print "found " + str(freq)

print part_two()