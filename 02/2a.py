from dataclasses import dataclass
import re


@dataclass
class GameReading:
    red: int = 0
    green: int = 0
    blue: int = 0


@dataclass
class Game:
    game_id: int
    readings: list[GameReading]


games_raw: list[str] = []
with open("02/input-2a.txt", encoding="utf8") as infile:
    games_raw = infile.readlines()

all_games: list[Game] = []
for game in games_raw:
    game = game.strip()

    game_name, readings = game.split(": ")

    game_id = int(game_name.split("Game ")[1])

    game_readings: list[GameReading] = []

    current_game = Game(game_id=game_id, readings=[])

    for reading in readings.split("; "):
        matches = re.finditer("(?P<count>\d+) (?P<colour>red|green|blue)", reading)
        game_reading = GameReading()
        for m in matches:
            colour = m.group("colour")
            count = int(m.group("count"))
            match colour:
                case "red":
                    game_reading.red = count
                case "green":
                    game_reading.green = count
                case "blue":
                    game_reading.blue = count
        current_game.readings.append(game_reading)

    all_games.append(current_game)

MAX_RED = 12
MAX_GREEN = 13
MAX_BLUE = 14

sum_valid_game_ids = 0
for game in all_games:
    valid = True
    for reading in game.readings:
        if (
            reading.red > MAX_RED
            or reading.green > MAX_GREEN
            or reading.blue > MAX_BLUE
        ):
            valid = False
    if valid:
        sum_valid_game_ids += game.game_id

    # if not any(reading.red > MAX_RED or reading.green > MAX_GREEN or reading.blue > MAX_BLUE for reading in game.readings):
    #     sum_valid_game_ids += game.game_id

print("SUM", sum_valid_game_ids)
