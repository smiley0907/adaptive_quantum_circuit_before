# ============================================================
# CELL 9: PREPARE CIRCUITS FOR EXECUTION
# ============================================================

execution_circuits = {}

for n in QUBIT_CONFIGS:

    execution_circuits[n] = transpile(
        original_circuits[n],
        simulator,
        optimization_level=0,
        seed_transpiler=RANDOM_SEED
    )

print("All circuits prepared for execution.")
