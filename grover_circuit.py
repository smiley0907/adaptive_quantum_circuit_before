# ============================================================
# CELL 4: GROVER CIRCUIT GENERATOR
# ============================================================

def create_grover_circuit(n_qubits):
    """
    Create a Grover search circuit.

    The marked state is represented using an oracle,
    followed by the Grover diffusion operation.
    """

    qc = QuantumCircuit(n_qubits, n_qubits)

    # --------------------------------------------------------
    # Initial superposition
    # --------------------------------------------------------
    qc.h(range(n_qubits))

    # --------------------------------------------------------
    # Oracle for the marked state
    # --------------------------------------------------------
    qc.x(range(n_qubits))

    if n_qubits == 1:
        qc.z(0)
    else:
        qc.h(n_qubits - 1)

        qc.mcx(
            list(range(n_qubits - 1)),
            n_qubits - 1
        )

        qc.h(n_qubits - 1)

    qc.x(range(n_qubits))

    # --------------------------------------------------------
    # Diffusion operator
    # --------------------------------------------------------
    qc.h(range(n_qubits))
    qc.x(range(n_qubits))

    if n_qubits == 1:
        qc.z(0)
    else:
        qc.h(n_qubits - 1)

        qc.mcx(
            list(range(n_qubits - 1)),
            n_qubits - 1
        )

        qc.h(n_qubits - 1)

    qc.x(range(n_qubits))
    qc.h(range(n_qubits))

    # --------------------------------------------------------
    # Measurement
    # --------------------------------------------------------
    qc.measure(
        range(n_qubits),
        range(n_qubits)
    )

    return qc
