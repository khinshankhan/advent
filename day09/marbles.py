from collections import deque

def play(max_players, beginning, end, circle):
    scores = [0] * max_players
    for marble in xrange(beginning, end + 1):
        if marble % 23 == 0:
            circle.rotate(7)
            scores[marble % max_players] += marble + circle.pop()
            circle.rotate(-1)
        else:
            circle.rotate(-1)
            circle.append(marble)
    return (circle, max(scores))

raw = [l.rstrip().split(' ') for l in open("input.txt")][0]
circle, score = play(int(raw[0]), 1, int(raw[6]), deque([0])) # part 1
circle, score2 = play(int(raw[0]), int(raw[6]) + 1, (int(raw[6]) * 100) + 1, circle) # part 2
print score, score2
