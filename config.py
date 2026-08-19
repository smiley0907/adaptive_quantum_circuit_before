# ============================================================
# CELL 3: EXPERIMENT CONFIGURATION
# ============================================================

QUBIT_CONFIGS = [3, 5, 7, 9, 11]

SHOTS = 1024

WARMUP_RUNS = 2

MEASUREMENT_RUNS = 10

RANDOM_SEED = 42

print("Qubit configurations :", QUBIT_CONFIGS)
print("Shots                :", SHOTS)
print("Warm-up executions   :", WARMUP_RUNS)
print("Measurement runs     :", MEASUREMENT_RUNS)
print("Random seed          :", RANDOM_SEED)
