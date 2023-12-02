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

sum_power_cubes = 0
for game in all_games:
    min_red = max(reading.red for reading in game.readings)
    min_green = max(reading.green for reading in game.readings)
    min_blue = max(reading.blue for reading in game.readings)
    sum_power_cubes += min_red * min_green * min_blue

print("SUM", sum_power_cubes)
