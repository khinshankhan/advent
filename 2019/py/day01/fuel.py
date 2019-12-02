def fuel(sum1, sum2):
    with f as open("input.txt"):
        for l in f:
            t = int(l)
            sum1 += t // 3 - 2
            while (t // 3 - 2) > 0:
                sum2 += t // 3 - 2
                t = t // 3 - 2
    return (sum1, sum2)

print(fuel(0, 0))
