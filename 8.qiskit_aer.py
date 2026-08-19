# ============================================================
# CELL 8: CONFIGURE QISKIT AER SIMULATOR
# ============================================================

def create_simulator():
    """
    Create the Qiskit Aer simulator used for the experiment.
    """

    simulator = AerSimulator(
        method="statevector",
        device="CPU"
    )

    return simulator


simulator = create_simulator()

print("Qiskit Aer simulator initialized successfully.")
