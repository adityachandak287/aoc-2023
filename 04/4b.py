import re
from dataclasses import dataclass


@dataclass
class ScoreRecord:
    game_id: str
    winning_count: int


INPUT_LINES = []
with open("04/input-4a.txt") as infile:
    INPUT_LINES = infile.read().split("\n")

score = 0

all_games = []

for line in INPUT_LINES:
    game_id, all_numbers = line.split(": ")
    winning_str, numbers_str = all_numbers.split(" | ")
    winning = [int(n) for n in re.split("\s+", winning_str.strip())]
    numbers = [int(n) for n in re.split("\s+", numbers_str.strip())]

    common = 0
    for num in numbers:
        if num in winning:
            common += 1

    # print(game_id, common)

    all_games.append(ScoreRecord(game_id=game_id, winning_count=common))

print("====")

all_counts = {}


def incr_count(game_id: str, incr: int):
    if game_id in all_counts:
        all_counts[game_id] += incr
    else:
        all_counts[game_id] = incr


for game_idx, game in enumerate(all_games):
    card_count = 1 + all_counts.get(game.game_id, 0)
    # print(f"Found {card_count} cards of {game.game_id}")

    # print("----", game_idx)
    for idx in range(game_idx + 1, game_idx + 1 + game.winning_count):
        # print(game_idx, idx, all_games[idx].game_id)
        incr_count(all_games[idx].game_id, card_count * 1)

    all_counts[game.game_id] = card_count

# print(all_counts)
print("SCORE", sum(all_counts.values()))
