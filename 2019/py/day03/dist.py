DIRS = {
    # we flip map upside down to start at (0,0)
    'U':(-1, 0),
    'D':(1, 0),
    'R': (0, 1),
    'L': (0, -1)
}


def get_steps(wire, steps):
    d, x, y = {}, 0, 0

    # dir is a keyword?
    for DIR, n in wire:
        # annoying but needed for dict
        for _ in range(n):
            dx, dy = DIRS[DIR]
            x += dx
            y += dy
            steps += 1
            d[(x, y)] = steps
    return d


def dist(f, val1, val2):
    listo = [l.strip().split(",") for l in f]
    listo[0] = [(i[0], int(i[1:])) for i in listo[0]]
    listo[1] = [(i[0], int(i[1:])) for i in listo[1]]

    d1 = get_steps(listo[0], 0)
    d2 = get_steps(listo[1], 0)

    # big numbers so they 'lose' first min operation
    # these variables need to be out of loop
    dist = 1000000
    steps = 1000000

    # dicts support set operations like intersection
    # https://www.it-swarm.net/fr/python/fusion-de-dictionnaires-de-dictionnaires/940336489/
    for pos in d1.keys() & d2.keys():
        # actual problem
        dist = min(abs(pos[0]) + abs(pos[1]), dist)
        steps = min(d1[pos] + d2[pos], steps)

    return (dist, steps)


if __name__ == "__main__":
    f = open("input.txt")
    print(dist(f, 0, 0))
    f.close()
