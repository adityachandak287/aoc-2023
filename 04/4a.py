import re

INPUT_LINES=[]
with open("04/input-4a.txt") as infile:
    INPUT_LINES = infile.read().split("\n")

score = 0
for line in INPUT_LINES:
    game_id, all_numbers = line.split(": ")
    winning_str, numbers_str = all_numbers.split(" | ")
    winning = [int(n) for n in re.split("\s+", winning_str.strip())]
    numbers = [int(n) for n in re.split("\s+", numbers_str.strip())]

    common = 0
    for num in numbers:
        if num in winning:
            common += 1

    card_score = pow(2, max(common-1, 0)) if common > 0 else 0
    score += card_score
    print(game_id, card_score)

print("SCORE", score)
