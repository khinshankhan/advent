def manhattan_distance(r, c, rr, cc):
  return abs(rr - r) + abs(cc - c)

def coords():
    arr = [l.rstrip().split(',') for l in open("input.txt")]
    return [map(int, x) for x in arr]

c = coords()
mapa = [[-1] * 500 for x in xrange(500)]
count = 0
for i in c:
    mapa[i[1]][i[0]] = count
    count += 1

for j in range(500):
    for k in range(500):
        letter, val, count = 9000, 9000, 0
        for i in c:
            if mapa[j][k] == -1:
                if manhattan_distance(j, k, i[1], i[0]) == val:
                    letter = -1
                if manhattan_distance(j, k, i[1], i[0]) < val:
                    letter, val = count, manhattan_distance(j, k, i[1], i[0])
                count += 1
            else:
                letter = mapa[j][k]
        mapa[j][k] = letter
edges = []
for i in range(500):
    edges.append(mapa[0][i])
    edges.append(mapa[i][0])
    edges.append(mapa[499][i])
    edges.append(mapa[i][499])
for i in range(500):
    for j in range(500):
        if mapa[i][j] in set(edges):
            mapa[i][j] = -1
'''
for row in mapa:
    line = ""
    for val in row:
        line += str(val)
    print line
'''
flat_list = [item for sublist in mapa for item in sublist]
s = filter(lambda x: x != -1, flat_list)
most = (max(set(s), key = s.count))
occur = 0
for i in flat_list:
    if i == most:
        occur += 1
print occur

points = []
for i in range(500):
    for j in range(500):
        value = 0
        for k in c:
            value += manhattan_distance(j, i, k[1], k[0])
        if value < 10000:
            points.append(mapa[i][j])
print len(points)
