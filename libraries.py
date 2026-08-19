# ============================================================
# CELL 1: ENVIRONMENT SETUP
# Adaptive Quantum Circuit Optimization Using Execution Feedback
# ============================================================

%pip install -q qiskit qiskit-aer pandas numpy matplotlib

# ============================================================
# CELL 2: IMPORT REQUIRED LIBRARIES
# ============================================================

import time
import platform
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

import qiskit
from qiskit import QuantumCircuit, transpile
from qiskit_aer import AerSimulator

print("Qiskit version :", qiskit.__version__)
print("Python version :", platform.python_version())
print("Platform       :", platform.platform())
