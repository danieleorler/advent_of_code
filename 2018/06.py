import sys, operator

def distance(a,b):
    return abs(a[0] - b[0]) + abs(a[1] - b[1])

def printMatrix(matrix, best = -100):
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if len(matrix[y][x][1]) == 1 and matrix[y][x][1][0] == best and matrix[y][x][0] == -1:
                color = "\033[92m"
                endColor = "\033[0m"
            elif len(matrix[y][x][1]) == 1 and matrix[y][x][1][0] == best and matrix[y][x][0] > -1:
                color = "\033[93m"
                endColor = "\033[0m"
            elif len(matrix[y][x][1]) == 1 and matrix[y][x][1][0] != best and matrix[y][x][0] == -1:
                color = "\033[92m"
                endColor = "\033[0m"
            else:
                color = ""
                endColor = ""

            if len(matrix[y][x][1]) == 1:
                if(0 <= matrix[y][x][1][0] < 10):
                    print "{} {}{}".format(color, matrix[y][x][1][0], endColor),
                else:
                    print "{}{}{}".format(color, matrix[y][x][1][0], endColor),
            else:
                print " *",
        print "[{}]".format(y)

def printMatrixSimple(matrix):
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if(len(matrix[y][x]) == 1):
                print " {}".format(matrix[y][x]),
            else:
                print "{}".format(matrix[y][x]),
        print "[{}]".format(y)

def parseCoordinates():
    c = []
    counter = 0
    for line in sys.stdin:
        c.append((int(line.rstrip().split(', ')[0]), int(line.rstrip().split(', ')[1]), counter))
        counter = counter + 1
    return c

coordinates = parseCoordinates()
minx = min(coordinates, key=operator.itemgetter(0))[0]
maxx = max(coordinates, key=operator.itemgetter(0))[0]
miny = min(coordinates, key=operator.itemgetter(1))[1]
maxy = max(coordinates, key=operator.itemgetter(1))[1]
coordinates = [(coordinate[0]-minx, coordinate[1]-miny, coordinate[2]) for coordinate in coordinates]

def part_one():
    matrix = [[(0,[-1]) for x in range(minx, maxx+1)] for y in range(miny, maxy+1)]

    for c in coordinates:
        try:
            matrix[c[1]][c[0]] = (-1, [c[2]])
        except IndexError:
            print "error at ({},{})".format(c[1], c[0])

    for c in coordinates:
        for y in range(len(matrix)):
            for x in range(len(matrix[y])):
                d = distance((c[0], c[1]), (x,y))
                if matrix[y][x][0] < 0:
                    pass
                elif matrix[y][x][0] == d:
                    matrix[y][x][1].append(c[2])
                elif matrix[y][x][0] == 0 or matrix[y][x][0] > d:
                    matrix[y][x] = (d, [c[2]])
                else:
                    pass

    areas = {}
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if len(matrix[y][x][1]) == 1:
                if not matrix[y][x][1][0] in areas:
                    areas[matrix[y][x][1][0]] = 0
                if areas[matrix[y][x][1][0]] >= 0:
                    if minx < x+minx < maxx and miny < y+miny < maxy:
                        areas[matrix[y][x][1][0]] = areas[matrix[y][x][1][0]] + 1
                    else:
                        areas[matrix[y][x][1][0]] = -1

    biggestArea = 0
    bestC = -1
    for area in areas:
        if areas[area] > biggestArea:
            biggestArea = areas[area]
            bestC = area
            
    print biggestArea

def totalDistance(location, coordinates):
    total = 0
    for coordinate in coordinates:
        total = total + distance(location, coordinate)
    return total

def part_two():    
    matrix = [["*" for x in range(minx, maxx+1)] for y in range(miny, maxy+1)]
    for c in coordinates:
        try:
            matrix[c[1]][c[0]] = str(c[2])
        except IndexError:
            print "error at ({},{})".format(c[1], c[0])

    totalArea = 0
    for y in range(len(matrix)):
        for x in range(len(matrix[y])):
            if totalDistance((x,y), coordinates) < 10000:
                matrix[y][x] = "#"
                totalArea = totalArea + 1

    print totalArea

part_one()
print "----"
part_two()
        