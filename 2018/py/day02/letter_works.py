def checksum():
    repeated = [0] * 50
    for l in open("input.txt"):
        line, d = str(l), {}
        for i in line:
            d[i] = d[i] + 1 if (i in d) else 1
        d = dict((k, v) for k, v in d.items() if v > 1)
        s = set(val for val_orig in d for val in d.values())
        for i in s:
            repeated[i] += 1
    s = filter(lambda x: x > 1, repeated)
    return reduce(lambda x,y: x*y, s)    

print (checksum())

def difference():
    ans, words = [], []
    for l in open("input.txt"):
        line = l.rstrip()
        for i in words:
            zipped, same = zip(i, line), ""
            for j,k in zipped:
                same += j if (j == k) else '*'
            occurences = [1 for i in same if i == '*']
            occurs = sum(occurences)
            if ans == [] or ans[1] > occurs:
                ans = [same, occurs]
        words += [line]
    return ans

print(difference())
