ROCK = "A"
PAPER = "B"
SCISSOR = "C"

SHAPE_SCORES = {
    "A": 1,
    "B": 2,
    "C": 3,
}

LOSS = "X"
DRAW = "Y"
WIN = "Z"

MOVES_TO_POINTS = {
    LOSS: 0,
    DRAW: 3,
    WIN: 6
}

MOVES_TO_SCORES = {
    ROCK: {
        ROCK: DRAW,
        PAPER: WIN,
        SCISSOR: LOSS,
    },
    PAPER: {
        ROCK: LOSS,
        PAPER: DRAW,
        SCISSOR: WIN,
    },
    SCISSOR: {
        ROCK: WIN,
        PAPER: LOSS,
        SCISSOR: DRAW,
    },
}

OUTCOME_TO_MOVES = {}
for opp_move in MOVES_TO_SCORES:
    OUTCOME_TO_MOVES[opp_move] = dict((v, k) for k,v in MOVES_TO_SCORES[opp_move].items())



def calc_round_score(round):
    round_guide = round.split(' ')
    opp_move = round_guide[0]
    outcome = round_guide[1]

    player_move = OUTCOME_TO_MOVES[opp_move][outcome]
    return MOVES_TO_POINTS[outcome] + SHAPE_SCORES[player_move]

def calc_total_score(strat_file_path):
    total_score = 0
    with open(strat_file_path, "r") as strat_file:
        for round in strat_file:
            total_score += calc_round_score(round.rstrip())
    return total_score

print(calc_total_score("./input"))
