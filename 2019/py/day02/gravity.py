from operator import add, mul

TARGET = 19690720
OPCODES = {
    1:add,
    2:mul,
    99: False
}

PART1_POS1 = 12
PART1_POS2 = 2

def program_alarm_1202(pos, listo):
    # destructure instructions to 4 variables
    opcode, op1, op2, nex = listo[pos: pos + 4]

    # determine if opcode continues programs or halts
    if not OPCODES[opcode]:
        return listo
    else:
        # execute instruction
        listo[nex] = OPCODES[opcode](listo[op1], listo[op2])
        # keep going until we reach a halt
        return program_alarm_1202(pos + 4, listo)

def gravity(f, val1, val2):
    for l in f:
        l = [int(i) for i in l.strip().split(",")]

        # Try all (noun, verb) until we found a good match
        for noun in range(100):
            for verb in range(100):
                # make a copy of list
                ll = [x for x in l]
                # set the initial states
                ll[1] = noun
                ll[2] = verb
                # run stateful changes on list
                ll = program_alarm_1202(0, ll)

                # a safe guess that program hasn't ended yet
                # really a heuristic based on my input
                if noun == PART1_POS1 and verb == PART1_POS2:
                    # part 1 complete
                    val1 = ll[0]
                if ll[0] == TARGET:
                    # pair works for part 2
                    # part 2 complete
                    val2 = (100 * noun) + verb

                    # we're done
                    return (val1, val2)


def main():
    f = open("input.txt")
    # print part and part 2
    print(gravity(f, 0, 0))
    f.close()


if __name__ == "__main__":
    main()
