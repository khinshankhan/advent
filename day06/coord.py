class gcs():
  def __init__(self):
    self.coords = [map(int, x) for x in [l.rstrip().split(',') for l in open("input.txt")]]
    self.rows = (max(i[1] for i in self.coords) + 1)
    self.cols = (max(i[0] for i in self.coords) + 1)
    self.grid = [[-1] * self.cols for x in xrange(self.rows)]
  def min_plot(self):
    for i in range(self.rows):
      for j in range(self.cols):
        letter, value, count = 9000, 9000, 0
        for k in self.coords:
          if self.grid[i][j] == -1:
            if manhattan_distance(i, j, k[1], k[0]) == value:
              letter = '.'
            if manhattan_distance(i, j, k[1], k[0]) < value:
              letter, value = chr(ord('a') + count), manhattan_distance(i, j, k[1], k[0])
            count += 1
          else:
            letter = self.grid[i][j]
        self.grid[i][j] = letter
  def edge_letters(self):
    edges = []
    for i in xrange(self.rows):
      for j in xrange(self.cols):
        edges.append(self.grid[0][j])
        edges.append(self.grid[i][0])
        edges.append(self.grid[self.rows - 1][j])
        edges.append(self.grid[i][self.cols - 1])
    self.edges = set(edges)
  def minimize_danger(self):
    self.edge_letters()
    self.minimized = [[-1] * self.cols for x in xrange(self.rows)]
    for i in xrange(self.rows):
      for j in xrange(self.cols):
        self.minimized[i][j] = '.' if self.grid[i][j] in self.edges else self.grid[i][j]
      self.minimized[i][j] = '.'

def manhattan_distance(r, c, rr, cc):
    return abs(rr - r) + abs(cc - c)
def flattened(list_2d):
  return [item for sublist in list_2d for item in sublist]

# part 1
a = gcs()
a.min_plot()
a.minimize_danger()
filtered = filter(lambda x: x != '.', flattened(a.minimized))
occur = filtered.count((max(set(filtered), key = filtered.count)))
print occur
# part 2
points = []
for i in range(a.rows):
  for j in range(a.cols):
    value = sum(manhattan_distance(j, i, k[1], k[0]) for k in a.coords)
    if value < 10000:
      points.append(value)
print len(points)
