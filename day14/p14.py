raw = [l.rstrip() for l in open("input.txt")][0]
score, e1, e2 = '37', 0, 1
while raw not in score[-7:]:
    t1 = int(score[e1])
    t2 = int(score[e2])
    score += str(t1 + t2)
    e1 = (e1 + t1 + 1) % len(score)
    e2 = (e2 + t2 + 1) % len(score)
print score[int(raw):int(raw)+10]
print score.index(raw)
