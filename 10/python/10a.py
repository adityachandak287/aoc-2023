import sys

sys.argv.append("10/sample-10a.txt")

assert len(sys.argv) > 1, "Arguments not passed"

lines = []
with open(sys.argv[1]) as infile:
    lines = infile.read().strip().split("\n")

print(lines)

grid = [[c for c in row] for row in lines]

start = ()
for row in range(len(grid)):
    for col in range(len(grid[0])):
        if grid[row][col] == "S":
            start = (row, col)
            print(start)

"""
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
"""

r, c = start
seen = [start]

def check(directions: list[str]):
    for d in directions:
        match d:
            case "north":
                if r > 0 and grid[r-1][c] in "|7F" and ((r-1, c) not in seen or (r-1, c) == start):
                    print("going north")
                    seen.append((r-1, c))
                    return (r-1, c)
            case "south":
                if r < len(grid) -1 and grid[r+1][c] in "|LJ" and ((r+1, c) not in seen or (r+1, c) == start):
                    print("going south")
                    seen.append((r+1, c))
                    return (r+1, c)
            case "west":
                if c > 0 and grid[r][c-1] in "-LF" and ((r, c-1) not in seen or (r, c-1) == start):
                    print("going west")
                    seen.append((r, c-1))
                    return (r, c-1)
            case "east":
                if c < len(grid[0])-1 and grid[r][c+1] in "-J7" and ((r, c+1) not in seen or (r, c+1) == start):
                    print("going east")
                    seen.append((r, c+1))
                    return (r, c+1)
    return -1, -1

while True:
    match grid[r][c]:
        case "S":
            r, c = check(["north", "south", "east", "west"])
        case "-":
            r, c = check(["east", "west"])
        case "|":
            r, c = check(["north", "south"])
        case "L":
            r, c = check(["north", "east"])
        case "J":
            r, c = check(["north", "west"])
        case "7":
            r, c = check(["south", "west"])
        case "F":
            r, c = check(["south", "east"])
        case ".":
            pass
        case _:
            assert False
    
    if(r, c) == (-1, -1):
        break

print("SEEN", len(seen), "ANSWER", len(seen)//2)
