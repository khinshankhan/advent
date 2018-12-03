def line_analysis(line):
    vals = line.rstrip().split(' ')
    first =  vals[2].split(':')[0].split(',')
    second = vals[3].split('x')
    id = int(vals[0].split('#')[1])
    return (id, int(first[0]), int(first[1]), int(second[0]), int(second[1]))

def claim():
    repeated = [[0] * 2000 for x in xrange(2000)]
    for l in open("input.txt"):
        id, left, top, wide, tall = line_analysis(l)
        for i in range(wide):
            for j in range(tall):
                if repeated[top + j][left + i] == 0:
                    repeated[top + j][left + i] = id
                else:
                    repeated[top + j][left + i] = 'X'
    return repeated

def overlapping(arr):
    occurences = 0
    for row in arr:
        for val in row:
            if val == 'X':
                occurences += 1
    return occurences

def not_overlapped(arr):
    for l in open("input.txt"):
        id, left, top, wide, tall = line_analysis(l)
        same = 0
        for i in range(wide):
            for j in range(tall):
                if arr[top + j][left + i] != id:
                    same = 1
        if (same == 0):
            return id
    return -1

print(overlapping(claim()))
print(not_overlapped(claim()))
