def checksum():
    repeated = []
    for i in range(50):
        repeated += [0]
    d = {}
    for l in open("input.txt"):
        line = str(l)
        d = {}
        for i in line:
            if i in d:
                d[i] += 1
            else:
                d[i] = 1
        d = dict((k, v) for k, v in d.items() if v > 1)
        s = set( val for val_orig in d for val in d.values())
        for i in s:
            repeated[i] += 1
    product = 1
    for i in repeated:
        if(i > 1):
            product *= i
            
    return product

print (checksum())

def difference():
    ans = []
    words = []
    for l in open("input.txt"):
        line = l.rstrip()
        for i in words:
            same = ""
            zipped = zip(i, line)
            for j,k in zipped:
                if(j == k):
                    same += j
                else:
                    same += '*'
            #print same
            occurences = 0
            if same != "":
                for i in same:
                    if i == '*':
                        occurences += 1
            if ans == [] or ans[1] > occurences:
                ans = [same, occurences]
        words += [line]
    return ans

print(difference())
