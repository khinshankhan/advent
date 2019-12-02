def fuel(f, sum1, sum2):
    for l in f:
        t = int(l)
        sum1 += t // 3 - 2
        while (t // 3 - 2) > 0:
            sum2 += t // 3 - 2
            t = t // 3 - 2
    return (sum1, sum2)


def main():
    f = open("input.txt")
    print(fuel(f, 0, 0))
    f.close()


if __name__ == "__main__":
    main()
