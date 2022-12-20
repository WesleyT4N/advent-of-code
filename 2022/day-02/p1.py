ROCK = "X"
PAPER = "Y"
SCISSOR = "Z"

OPP_ROCK = "A"
OPP_PAPER = "B"
OPP_SCISSOR = "C"

SHAPE_SCORES = {
    "X": 1,
    "Y": 2,
    "Z": 3,
}

LOSS = 0
DRAW = 3
WIN = 6


OUTCOME_SCORES = {
    OPP_ROCK: {
        ROCK: DRAW,
        PAPER: WIN,
        SCISSOR: LOSS,
    },
    OPP_PAPER: {
        ROCK: LOSS,
        PAPER: DRAW,
        SCISSOR: WIN,
    },
    OPP_SCISSOR: {
        ROCK: WIN,
        PAPER: LOSS,
        SCISSOR: DRAW,
    },
}

def calc_round_score(round):
    moves = round.split(' ')
    opp_move = moves[0]
    player_move = moves[1]

    return OUTCOME_SCORES[opp_move][player_move] + SHAPE_SCORES[player_move]

def calc_total_score(strat_file_path):
    total_score = 0
    with open(strat_file_path, "r") as strat_file:
        for round in strat_file:
            total_score += calc_round_score(round.rstrip())
    return total_score

print(calc_total_score("./input"))
