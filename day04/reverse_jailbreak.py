def sort_data():
    return sorted([l.rstrip() for l in open("input.txt")])

class Time():
    def __init__(self):
        self.minutes = [0] * 59
        
    def __init__(self, time_start, interval):
        self.minutes = [0] * 59
        self.edit(time_start, interval)
        
    def edit(self, time_start, interval):
        for i in range(interval):
            self.minutes[time_start + i] += 1
            
    def max_value(self):
        return max(self.minutes)
    
    def max_index(self):
        return self.minutes.index(max(self.minutes))
        
def sleepy():
    times, guards, id, time_start = sort_data(), [0] * 5000, -1, -1
    for i in times:
        if "#" in i:
            id = int(i.split(' ')[3].split('#')[1])
        if "asleep" in i:
            time_start = int(i.split(' ')[1].split(']')[0].split(':')[1])
        if "wakes" in i:
            time_end = int(i.split(' ')[1].split(']')[0].split(':')[1])
            if guards[id] != 0:
                guards[id][1].edit(time_start, time_end - time_start)
                guards[id][0] += time_end - time_start
            else:
                guards[id] = [time_end - time_start, Time(time_start, time_end - time_start)]
    interval, value = -1, -1
    for i in range(5000):
        if guards[i] != 0:
            if interval < guards[i][0]:
                interval, value = guards[i][0], i * guards[i][1].max_index()
    return value
    
print(sleepy())

def sleepy2():
    times, guards, id, time_start = sort_data(), [0] * 5000, -1, -1
    for i in times:
        if "#" in i:
            id = int(i.split(' ')[3].split('#')[1])
        if "asleep" in i:
            time_start = int(i.split(' ')[1].split(']')[0].split(':')[1])
        if "wakes" in i:
            time_end = int(i.split(' ')[1].split(']')[0].split(':')[1])
            if guards[id] != 0:
                guards[id].edit(time_start, time_end - time_start)
            else:
                guards[id] = Time(time_start, time_end - time_start)
    id, minute, value = -1, -1, -1
    for i in range(5000):
        if guards[i] != 0:
            if value < guards[i].max_value():
                id, minute, value = i, guards[i].max_index(), guards[i].max_value()
    return id * minute

print(sleepy2())
