from collections import deque

def play(max_players, beginning, end, circle, scores):
    for marble in xrange(beginning, end + 1):
        if marble % 23 == 0:
            for right_rotations in range(7):
                circle.appendleft(circle.pop())
            scores[marble % max_players] += marble + circle.pop()
            for left_rotations in range(1):
                circle.append(circle.popleft())
        else:
            for left_rotations in range(1):
                circle.append(circle.popleft())
            circle.append(marble)
    return (circle, max(scores), scores)

raw = [l.rstrip().split(' ') for l in open("input.txt")][0]
circle, score, scores = play(int(raw[0]), 1, int(raw[6]), deque([0]), [0] * int(raw[0])) # part 1
circle, score2, scores = play(int(raw[0]), int(raw[6]) + 1, (int(raw[6]) * 100), circle, scores) # part 2
print score, score2
