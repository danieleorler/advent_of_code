import re, sys
pattern = re.compile("Step\s(?P<from>[A-Z]+)\smust.*step\s(?P<to>[A-Z]+).*")

def allIn(them, s):
    result = True
    for it in them:
        result = result and it in s
    return result

def schedule(q, g, i):
    if i in g:
        q.extend(g[i])
        q.sort()

graph = {}
for line in sys.stdin:
    match = pattern.match(line)
    if not match.group("from") in graph:
        graph[match.group("from")] = []
    graph[match.group("from")].append(match.group("to"))

allDestinations = set()
for v in graph:
    allDestinations.update(graph[v])

# dependencies graph
dependencies = {}
dv = set()
for v in graph:
    dv.update(graph[v])
for uv in dv:
    if uv not in dependencies:
        dependencies[uv] = []
    for v in graph:
        if uv in graph[v]:
            dependencies[uv].append(v)

def part_one():
    queue = []
    for v in graph:
        if v not in allDestinations:
            queue.append(v)
    queue.sort()
    executionSequence = ""
    while len(queue) > 0:
        next = queue.pop(0)
        if next in executionSequence: #already executed
            continue
        if not next in dependencies: #doesn't need any dependency
            executionSequence = executionSequence + next
            schedule(queue, graph, next)
            continue

        if allIn(dependencies[next], executionSequence): #all dependencies resolved
            executionSequence = executionSequence + next
            schedule(queue, graph, next)
        else: #missing dependencies
            queue.append(next)

    print executionSequence

def scheduleFancy(queue, graph, task, step, t):
    if task in graph:
        for dep in graph[task]:
            candidate = (dep, step + ord(dep) - ord('A') + 1)
            if not candidate in queue:
                queue.append(candidate)

def cockpit(processing, workers, time):
    for task in processing:
        print " {} ".format('%02d' % processing[task]),

    if len(processing.keys()) < workers:
        for i in range(workers - len(processing.keys())):
            print " {} ".format('%02d' % 0),

    print time

def part_two(workers=5, step=60):
    processing = {}
    executionSequence = ""
    for v in graph:
        if v not in allDestinations:
            processing[v] = step + ord(v) - ord('A') + 1
    executionTime = 0
    queue = []
    while len(processing.keys()) > 0:
        completed = set()
        # processing
        for key in processing:
            processing[key] = processing[key]-1
            if processing[key] == 0:
                executionSequence = executionSequence + key
                completed.add(key)
        
        # check if any task has finished
        for key in completed:
            completedTask = processing[key]
            del processing[key]
            if key in graph:
                scheduleFancy(queue, graph, key, step, executionTime)
        queue.sort(key=lambda x: x[0])

        # schedule new tasks
        stop = 0
        while len(processing.keys()) < workers:
            try:
                candidate = queue.pop(0)
                if allIn(dependencies[candidate[0]], executionSequence):
                    processing[candidate[0]] = candidate[1]
                else:
                    queue.append(candidate)
                stop = stop +1
                if stop > len(queue):
                    break
            except IndexError as e:
                break

        cockpit(processing, workers, executionTime)
        executionTime = executionTime + 1

    print executionTime - 1
    print executionSequence

part_one()
print "-----"
part_two(5, 60)
