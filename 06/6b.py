import sys
import re

assert len(sys.argv) > 1, "Arguments not passed"

lines = []
with open(sys.argv[1]) as infile:
    lines = infile.read().strip().split('\n')

print(lines)

times = int("".join(re.split("\s+",re.split("Time:\s+", lines[0])[1])))
distances = int("".join(re.split("\s+",re.split("Distance:\s+", lines[1])[1])))

answer = 1

for t, d in  zip([times], [distances]):
    print(t,d)
    combos = 0
    for i in range(0, t+1):
        distance_travelled = (t-i) * i
        if distance_travelled > d:
            combos += 1
    print("combos for", t,d,combos)
    assert combos > 0, "atleast 1 possible way to win"
    answer *= combos

print(answer)
