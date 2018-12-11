def powerlevel(x, y, serial):
    rackid = x + 10
    incr = (((rackid * y + serial) * rackid) // 100) % 10
    return incr - 5

def powers(serial):
    full = []
    for i in range(1, 301):
        temp = []
        str1 = ""
        for j in range(1, 301):
            temp.append([j,i])
        full.append(temp)
    levels = []
    for i in full:
        temp = []
        for j in i:
            temp.append(powerlevel(j[0], j[1], serial))
        levels.append(temp)
    return levels

def power_sum(grid, x, y, length=3):
    return sum(grid[i][j] for i in range(x, x+length) for j in range(y, y+length))

def p8_remix(serial, box_length):
    grid = powers(serial)
    result = (None, None)
    best_power_sum = -9000
    for x in range(300 - box_length):
        for y in range(300 - box_length):
            p = power_sum(grid, x, y, box_length)
            if p > best_power_sum:
                best_power_sum = p
                result = (y + 1, x + 1)
    return [result, best_power_sum]
print p8_remix(7803, 3)
ranged_power = []
for i in range(int(300**0.5)*2):
    ranged_power.append([p8_remix(7803, i), i])
print max(ranged_power, key=lambda x: x[0][1])
