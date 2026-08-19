# ============================================================
# CELL 11: ORIGINAL CIRCUIT MEASUREMENT
# ============================================================

def measure_original_circuit(
    circuit,
    shots=SHOTS,
    warmup_runs=WARMUP_RUNS,
    measurement_runs=MEASUREMENT_RUNS,
    seed=RANDOM_SEED
):
    """
    Execute the original circuit using:
        - 2 warm-up executions
        - 10 measurement executions

    Returns the individual execution times.
    """

    # --------------------------------------------------------
    # Warm-up executions
    # --------------------------------------------------------

    for i in range(warmup_runs):

        execute_and_measure(
            simulator,
            circuit,
            shots=shots,
            seed=seed + i
        )

    # --------------------------------------------------------
    # Measurement executions
    # --------------------------------------------------------

    execution_times = []

    for i in range(measurement_runs):

        execution_time, result = execute_and_measure(
            simulator,
            circuit,
            shots=shots,
            seed=seed + warmup_runs + i
        )

        execution_times.append(execution_time)

    return execution_times
