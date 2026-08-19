# ============================================================
# CELL 15: SAVE ORIGINAL EXECUTION RESULTS
# ============================================================

original_df.to_csv(
    "original_circuit_results.csv",
    index=False
)

print(
    "Original execution dataset saved as "
    "'original_circuit_results.csv'"
)
