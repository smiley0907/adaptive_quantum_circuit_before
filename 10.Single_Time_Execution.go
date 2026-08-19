# ============================================================
# CELL 10: SINGLE TIMED EXECUTION
# ============================================================

def execute_and_measure(
    simulator,
    circuit,
    shots=SHOTS,
    seed=RANDOM_SEED
):
    """
    Execute one quantum circuit and return its execution time.
    """

    start_time = time.perf_counter()

    job = simulator.run(
        circuit,
        shots=shots,
        seed_simulator=seed
    )

    result = job.result()

    end_time = time.perf_counter()

    execution_time = end_time - start_time

    return execution_time, result
