def summation():
    sum = 0
    for l in open("input.txt"):
        sum += int(l)
    return sum

print(summation())

def summation_repeat(sum, sums):
    for l in open("input.txt"):
        sum += int(l)
        if(sum in sums):
            return sum
        sums += [sum]
    return summation_repeat(sum, sums)

print(summation_repeat(0, [0]))
