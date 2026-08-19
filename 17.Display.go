# ============================================================
# CELL 17: STORE INDIVIDUAL EXECUTION MEASUREMENTS
# ============================================================

original_measurements = []

for n in QUBIT_CONFIGS:

    circuit = execution_circuits[n]

    execution_times = measure_original_circuit(
        circuit,
        shots=SHOTS,
        warmup_runs=WARMUP_RUNS,
        measurement_runs=MEASUREMENT_RUNS,
        seed=RANDOM_SEED + 500 + n
    )

    for repetition, execution_time in enumerate(
        execution_times,
        start=1
    ):

        original_measurements.append({
            "Qubits": n,
            "Repetition": repetition,
            "Execution_Time_sec": execution_time
        })

original_measurements_df = pd.DataFrame(
    original_measurements
)

display(original_measurements_df)
