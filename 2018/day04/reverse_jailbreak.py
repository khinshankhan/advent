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

def analyze():
    times, sleep_sched, interval_sched, id, time_start = sort_data(), [0] * 5000, [0] * 5000, -1, -1
    for i in times:
        if "#" in i:
            id = int(i.split(' ')[3].split('#')[1])
        if "asleep" in i:
            time_start = int(i.split(' ')[1].split(']')[0].split(':')[1])
        if "wakes" in i:
            time_end = int(i.split(' ')[1].split(']')[0].split(':')[1])
            if sleep_sched[id] != 0:
                sleep_sched[id].edit(time_start, time_end - time_start)
                interval_sched[id] += time_end - time_start
            else:
                sleep_sched[id] = Time(time_start, time_end - time_start)
                interval_sched[id] = time_end - time_start
    return (sleep_sched, interval_sched)

def sleepy():
    sleeps, intervals = analyze()
    interval, value = -1, -1 # part 1
    id, minute, value2 = -1, -1, -1 # part 2
    for i in range(5000):
        if sleeps[i] != 0 and interval < intervals[i]: # part 1
            interval, value = intervals[i], i * sleeps[i].max_index()
        if sleeps[i] != 0 and value2 < sleeps[i].max_value(): # part 2
            id, minute, value2 = i, sleeps[i].max_index(), sleeps[i].max_value()
    return (value, id * minute)

print(sleepy())
