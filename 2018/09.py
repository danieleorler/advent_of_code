import re, sys, math
pattern = re.compile("(?P<players>[0-9]+)\splayers;.*worth\s(?P<value>[0-9]+).*")

def read_input():
    i = 0
    input = {}
    for line in sys.stdin:
        match = pattern.match(line)
        input[i] = (int(match.group("players")), int(match.group("value")))
        i = i + 1
    return input

def cockpit(circle, current, player):
    print "[{}] ".format(player),
    for i in range(len(circle)):
        if i == current:
            print "{}{}{} [{}]".format("\033[92m", str(circle[i]).rjust(2, " "), "\033[0m", str(i).rjust(2, " ")),
        elif circle[i] == 1:
            print "{}{}{} [{}]".format("\033[93m", str(circle[i]).rjust(2, " "), "\033[0m", str(i).rjust(2, " ")),
        elif circle[i] == 2:
            print "{}{}{} [{}]".format("\033[1;36m", str(circle[i]).rjust(2, " "), "\033[0m", str(i).rjust(2, " ")),
        elif circle[i] == 3:
            print "{}{}{} [{}]".format("\033[1;34m", str(circle[i]).rjust(2, " "), "\033[0m", str(i).rjust(2, " ")),
        else:
            print "{} [{}]".format(str(circle[i]).rjust(2, " "), str(i).rjust(2, " ")),
    print

def part_one(input):
    score = [0 for p in range(input[0])]
    circle = [0]
    currentIndex = 0
    prizeFound = 0
    newline = 0
    for i in range(1,input[1]+1):
        if i < 2:
            circle.append(i)
            currentIndex = len(circle) - 1
        else:
            if i % 23 == 0:
                otherPos = (currentIndex - 7) % len(circle)
                score[i%input[0]] += i + circle[otherPos]
                del circle[otherPos]
                currentIndex = otherPos
                prizeFound = prizeFound + 1
            else:
                if currentIndex == len(circle) - 1:
                    circle.insert(1, i)
                    currentIndex = 1
                    prizeFound = 0
                    newline = newline + 1
                else:
                    circle.insert(currentIndex+2, i)
                    currentIndex = currentIndex+2
        
        #cockpit(circle, currentIndex, i%input[0])
    return score

#######################################
### \/FAILED ATTEMPT TO BE SMART \/ ###
#######################################

def part_two(input):
    for i in range(1,input[1]+1):
        if i % 23 == 0:
            print i%input[0]
input = read_input()

for k in input.keys():
    print "[{}] ==> {}".format(k, max(part_one(input[k])))

def pos(n):
    if n == 0:
        return 0

    if n % 23 == 0:
        return -999

    anomalies = n/23
    shiftedN = n - (4 * 1 if anomalies else 0)
    if shiftedN == pow(2, int(math.log(shiftedN,2))):
        return 1
    return ((shiftedN - pow(2, int(math.log(shiftedN,2)))) * 2 + 1) - (1 if anomalies else 0)

def drift(a,b):
    if a > b:
        return pos(a)
    current_pos = pos(a)
    anomalies = b/23 - a/23
    for i in range(a+1,b+1):
        if pos(i) <= current_pos:
            current_pos = current_pos + 1
    return current_pos - anomalies*2
