import re

INPUT_LINES = []
with open("04/input-04.txt") as infile:
    INPUT_LINES = infile.read().split("\n")

score = 0

all_games = []

line_pattern = re.compile(":\s+| \| ")
space_pattern = re.compile("\s+")

for line in INPUT_LINES:
    game_id, winning_str, numbers_str = line_pattern.split(line)
    winning = [int(n) for n in space_pattern.split(winning_str.strip())]
    numbers = [int(n) for n in space_pattern.split(numbers_str.strip())]

    common = 0
    for num in numbers:
        if num in winning:
            common += 1

    all_games.append({"game_id": game_id, "winning_count": common})

all_counts = {}

for game_idx, game in enumerate(all_games):
    card_count = 1 + all_counts.get(game["game_id"], 0)

    for idx in range(game_idx + 1, game_idx + 1 + game["winning_count"]):
        all_counts[all_games[idx]["game_id"]] = (
            all_counts.get(all_games[idx]["game_id"], 0) + card_count
        )

    all_counts[game["game_id"]] = card_count


print("SCORE", sum(all_counts.values()))
