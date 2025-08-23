import matplotlib.pyplot as plt
import pandas as pd

df = pd.read_csv("random_bits32.csv")
df = df[:1000]

plt.figure(figsize=(8, 6))
plt.scatter(df["bits"], df["dragonbox_ns"], label="Dragonbox", s=1, alpha=0.7)
plt.scatter(df["bits"], df["ryu_ns"], label="Ryu", s=1, alpha=0.7)

plt.xlabel("Bits")
plt.ylabel("Time (ns)")
plt.title("Dragonbox vs Ryu Performance")
plt.legend()
plt.grid(True)

plt.ylim(0, 100)
plt.show()
