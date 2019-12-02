def set(pos, listo):
    if listo[pos] == 1:
        listo[listo[pos + 3]] = listo[listo[pos + 1]] + listo[listo[pos + 2]]
        return set(pos + 4, listo)
    elif listo[pos] == 2:
        listo[listo[pos + 3]] = listo[listo[pos + 1]] * listo[listo[pos + 2]]
        return set(pos + 4, listo)
    else:
        return listo

def gravity(f, val1, val2):
    for l in f:
        l = [int(i) for i in l.strip().split(",")]
        ll = [x for x in l]
        l[1] = 12
        l[2] = 2
        l = set(0, l)
        val1 = l[0]

        for i in range(100):
            for j in range(100):
                lll = [x for x in ll]
                lll[1] = i
                lll[2] = j
                lll = set(0, lll)
                if lll[0] == 19690720:
                    val2 = (100 * i) + j
                    break
    return (val1, val2)


def main():
    f = open("input.txt")
    print(gravity(f, 0, 0))
    f.close()


if __name__ == "__main__":
    main()
