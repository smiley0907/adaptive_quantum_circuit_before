# ============================================================
# CELL 16: ORIGINAL EXECUTION TIME GRAPH
# ============================================================

plt.figure(figsize=(9, 5))

plt.plot(
    original_df["Qubits"],
    original_df["Median_Time_sec"],
    marker="o",
    linewidth=2
)

plt.xlabel("Number of Qubits")
plt.ylabel("Median Execution Time (s)")
plt.title("Original Circuit Execution Performance")

plt.grid(True)
plt.tight_layout()
plt.show()
