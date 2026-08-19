# ============================================================
# CELL 14: FORMATTED ORIGINAL RESULTS
# ============================================================

original_summary = original_df[
    [
        "Qubits",
        "Gate_Count",
        "Circuit_Depth",
        "Median_Time_sec",
        "Mean_Time_sec",
        "Std_Time_sec"
    ]
].copy()

display(original_summary)
