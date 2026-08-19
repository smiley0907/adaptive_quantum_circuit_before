# ============================================================
# CELL 6: ORIGINAL CIRCUIT CHARACTERISTICS
# ============================================================

circuit_characteristics = []

for n in QUBIT_CONFIGS:

    circuit = original_circuits[n]

    circuit_characteristics.append({
        "Qubits": n,
        "Gate_Count": circuit.size(),
        "Circuit_Depth": circuit.depth()
    })

original_characteristics_df = pd.DataFrame(
    circuit_characteristics
)

display(original_characteristics_df)
