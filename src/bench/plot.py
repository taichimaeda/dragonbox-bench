import sys

import pandas as pd
import plotly.express as px

for bit_size in [32, 64]:
    csv_file = f"random_bits{bit_size}.csv" 
    output_file = f"random_bits{bit_size}.png"

    df = pd.read_csv(csv_file)
    fig = px.scatter(
        df,
        x="bits",
        y=["dragonbox_ns", "ryu_ns"],
        labels={"value": "Time (ns)", "bits": "Bit Patterns"},
        title="Dragonbox vs Ryu Performance"
    )

    fig.update_traces(marker=dict(size=1, opacity=0.5))
    fig.update_yaxes(range=[30, 70])
    fig.write_image(output_file, width=1200, height=675)

for bit_size in [32, 64]:
    csv_file = f"random_digits{bit_size}.csv" 
    output_file = f"random_digits{bit_size}.png"

    df = pd.read_csv(csv_file)
    df = df.groupby("digits").mean().reset_index()
    fig = px.line(
        df,
        x="digits",
        y=["dragonbox_ns", "ryu_ns"],
        labels={"value": "Time (ns)", "digits": "Number of Digits"},
        title="Dragonbox vs Ryu Performance"
    )

    fig.update_traces(line=dict(width=2))
    fig.update_yaxes(range=[30, 70])
    fig.write_image(output_file, width=1200, height=675)