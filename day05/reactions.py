class Polymer():
    def __init__(self, line = [l.rstrip() for l in open("input.txt")][0]):
        self.polymers, self.line = [], line
        for i in self.line:
            if self.polymers and self.reactive(i, self.polymers[-1]):
                self.polymers.pop()
            else:
                self.polymers.append(i)
    def reactive(self, unit, unit2):
        return (unit.lower() == unit2.lower()) and (unit != unit2)
    def length(self):
        return len(self.polymers)
    def removals(self):
        for i in range(26):
            upper, lower = chr(ord('A') + i), chr(ord('a') + i)
            yield self.line.replace(upper, '').replace(lower, '')
    def shortest(self):
        return min((Polymer(l).length() for l in self.removals()))

p = Polymer()
print(p.length())
print(p.shortest())
