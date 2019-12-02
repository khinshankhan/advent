import re

coord_vel = [map(int, re.findall('-?\d+', l)) for l in open("input.txt")]

for t in range(100000):
    min_x = min([x for x,y,_,_ in coord_vel])
    max_x = max([x for x,y,_,_ in coord_vel])
    min_y = min([y for x,y,_,_ in coord_vel])
    max_y = max([y for x,y,_,_ in coord_vel])
    W = 100 # screen basically. outrageous if i cant eyeball it, so just assume
    if min_x+W >= max_x and min_y + W >= max_y:
        print t,min_x, max_x, min_y, max_y # time and the size variables (just in case... debugging)
        for y in range(min_y, max_y+1):
            for x in range(min_x, max_x+1):
                if (x,y) in [(px,py) for px,py,_,_ in coord_vel]:
                    print '#',
                else:
                    print '.',
            print

    for p in coord_vel:
        p[0] += p[2]
        p[1] += p[3]

