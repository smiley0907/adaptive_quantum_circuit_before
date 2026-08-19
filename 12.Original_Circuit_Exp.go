# ============================================================
# CELL 12: ORIGINAL CIRCUIT EXPERIMENT
# ============================================================

original_results = []

for n in QUBIT_CONFIGS:

    circuit = execution_circuits[n]

    print()
    print("-" * 60)
    print(f"Running Original Circuit: {n} qubits")
    print("-" * 60)

    execution_times = measure_original_circuit(
        circuit,
        shots=SHOTS,
        warmup_runs=WARMUP_RUNS,
        measurement_runs=MEASUREMENT_RUNS,
        seed=RANDOM_SEED + n
    )

    # --------------------------------------------------------
    # Statistical measurements
    # --------------------------------------------------------

    median_time = float(
        np.median(execution_times)
    )

    mean_time = float(
        np.mean(execution_times)
    )

    std_time = float(
        np.std(execution_times, ddof=1)
    )

    min_time = float(
        np.min(execution_times)
    )

    max_time = float(
        np.max(execution_times)
    )

    original_results.append({
        "Qubits": n,
        "Gate_Count": original_circuits[n].size(),
        "Circuit_Depth": original_circuits[n].depth(),
        "Shots": SHOTS,
        "Median_Time_sec": median_time,
        "Mean_Time_sec": mean_time,
        "Std_Time_sec": std_time,
        "Min_Time_sec": min_time,
        "Max_Time_sec": max_time
    })

    print(
        f"Median : {median_time:.6f} sec"
    )

    print(
        f"Mean   : {mean_time:.6f} sec"
    )

    print(
        f"Std    : {std_time:.6f} sec"
    )
