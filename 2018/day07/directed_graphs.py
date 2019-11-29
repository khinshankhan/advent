import networkx # needs pip install
from collections import defaultdict # defaultdict is really useful

# Used reddit for interesting approach idea

# get data
def data():
    return [l.rstrip().split(' ') for l in open("input.txt")]
# organize data as almost adjacency matrix
def G(lines):
    return [(i[1], i[7]) for i in lines]
# store for now and later
connections = [(i[1], i[7]) for i in data()]
# uses networkx's nice graphical functions to just solve
def solve():
    return ''.join(networkx.lexicographical_topological_sort(networkx.DiGraph(connections)))
# end part 1
print solve()

# proletariat
class Worker:
    def __init__(self):
        self.task, self.done, self.time_rem = None, False, 0
    def set_task(self, task):
        self.task, self.done, self.time_rem = task, False, (ord(task) - ord('A') + 1) + 60
    def step(self):
        if not self.done:
            self.time_rem -= 1
            self.done = True if (self.time_rem == 0) else False

# for later usage
befores = defaultdict(list)
afters = defaultdict(list)
# 'unzip' the connections for dependency lists
for (src, dst) in connections:
    befores[src].append(dst)
    afters[dst].append(src)
# instantiate 5 workers into an array for access
workers = [Worker() for _ in range(5)]
# alpha contains jobs as a heap, order is jobs done (so we can end once alphabet is over)
alpha, order = [], []
alpha.extend(set(befores.keys()) - set(afters.keys())) # add keys without dependecies
alpha.sort(reverse=True) # reverse sort to treat as priority heap
# time
seconds = 0
# iterate until all jobs are donezo
while len(order) != 26:
    # get num of workers free
    w = [worker for worker in workers if worker.time_rem <= 0]
    # assign any possible tasks to available workers
    for worker in w:
        if alpha:
            next_task = alpha.pop()
            worker.set_task(next_task)
    # workers work
    for worker in workers:
        worker.step()
        # keep track of jobs done
        if worker.done:
            order.append(worker.task)
            # next batch of available jobs
            for s in befores[worker.task]:
                preds = afters[s]
                if s not in order and all(p in order for p in preds):
                    alpha.append(s)
                    alpha.sort(reverse=True)
            worker.done = False
    # where does the time go :")
    seconds += 1

print(seconds)
