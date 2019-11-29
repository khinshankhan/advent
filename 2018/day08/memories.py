data = map(int, [l.rstrip().split(' ') for l in open("input.txt")][0])

# i love reddit, my code but bettered with naming and smarter recursion
def parse(data):
    nodes, metas = data[:2]
    data = data[2:]
    scores = []
    totals = 0

    for i in range(nodes):
        total, score, data = parse(data)
        totals += total
        scores.append(score)

    totals += sum(data[:metas])

    if nodes == 0:
        return (totals, sum(data[:metas]), data[metas:])
    else:
        return (
            totals,
            sum(scores[k - 1] for k in data[:metas] if k > 0 and k <= len(scores)),
            data[metas:]
        )

total, value, remaining = parse(data)
print total, value
