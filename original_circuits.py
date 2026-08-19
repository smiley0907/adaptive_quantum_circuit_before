# ============================================================
# CELL 5: GENERATE ORIGINAL CIRCUITS
# ============================================================

original_circuits = {}

for n in QUBIT_CONFIGS:

    original_circuits[n] = create_grover_circuit(n)

print("Original circuits generated successfully.")
