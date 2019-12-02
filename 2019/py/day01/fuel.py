def fuel(sum1, sum2):
    for l in open("input.txt"):
        t = int(l)
        sum1 += t // 3 - 2
        while (t // 3 - 2) > 0:
            sum2 += t // 3 - 2
            t = t // 3 - 2
    return (sum1, sum2)

print(fuel(0, 0))
