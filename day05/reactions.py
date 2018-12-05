def reactive(unit, unit2):
    return (unit.lower() == unit2.lower()) and (unit != unit2)

def polymer(line):
    polymers = []
    for i in line:
        if polymers and reactive(i, polymers[-1]):
            polymers.pop()
        else:
            polymers += [i]
    return polymers

def shortest_removal(line):
    value = -1
    for i in range(26):
        upper, lower = chr(ord('A') + i), chr(ord('a') + i)
        l = line.replace(upper, '').replace(lower, '')
        value = len(polymer(l)) if value == -1 or len(polymer(l)) < value else value
    return value

line = [l for l in open("input.txt")][0].rstrip()

print(len(polymer(line))) # part 1

print(shortest_removal(line))
