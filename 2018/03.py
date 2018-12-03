import re, sys
pattern = re.compile("#(?P<id>[0-9]+)\s@\s(?P<left>[0-9]+),(?P<bottom>[0-9]+):\s(?P<width>[0-9]+)x(?P<height>[0-9]+)")

class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return "({},{})".format(self.x, self.y)

class Claim:
    def __init__(self, bl, tr, id):
        self.bl = bl
        self.tr = tr
        self.id = id

def parse(input):
    match = pattern.match(input)
    id = int(match.group("id"))
    left = int(match.group("left"))
    bottom = int(match.group("bottom"))
    width = int(match.group("width"))
    height = int(match.group("height"))
    return Claim(Point(x=left, y=bottom), Point(x=left+width, y=bottom+height), id)

def overlap(a, b):
    if a.bl.y >= b.tr.y or b.bl.y >= a.tr.y or a.bl.x >= b.tr.x or b.bl.x >= a.tr.x:
        return []
    else:
        vbl = Point(max(a.bl.x, b.bl.x), max(a.bl.y, b.bl.y))
        vtr = Point(min(a.tr.x, b.tr.x), min(a.tr.y, b.tr.y))
        result = []
        for y in range(vbl.y,vtr.y):
            for x in range(vbl.x, vtr.x):
                result.append("[{},{}]".format(x,y))
        return result

def overlapId(a,b):
    if a.bl.y >= b.tr.y or b.bl.y >= a.tr.y or a.bl.x >= b.tr.x or b.bl.x >= a.tr.x:
        return None
    else:
        return [a.id, b.id]

def part_one():
    claims = []
    for line in sys.stdin:
        claims.append(parse(line))
    
    overlapping = set()
    for i in range(len(claims)):
        for j in range(i+1, len(claims)):
            overlapping.update(overlap(claims[i], claims[j]))
            
    print len(overlapping)

def part_two():
    claims = []
    nonOverlapping = set()
    for line in sys.stdin:
        claim = parse(line)
        claims.append(claim)
        nonOverlapping.add(claim.id)
    
    for i in range(len(claims)):
        for j in range(i+1, len(claims)):
            claimIds = overlapId(claims[i], claims[j])
            if claimIds is not None:
                if claimIds[0] in nonOverlapping:
                    nonOverlapping.remove(claimIds[0])
                if claimIds[1] in nonOverlapping:
                    nonOverlapping.remove(claimIds[1])
            
    print nonOverlapping  

part_two()