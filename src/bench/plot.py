import pandas as pd
import plotly.graph_objects as go

for bit_size in [32, 64]:
    csv_file = f"random_bits{bit_size}.csv"
    output_file = f"random_bits{bit_size}.png"
    columns = ["bits", "dragonbox_ns", "ryu_ns"]

    df = pd.read_csv(csv_file)

    fig = go.Figure()
    fig.add_trace(go.Scatter(
        x=df["bits"],
        y=df["dragonbox_ns"],
        mode='markers',
        marker=dict(size=1, opacity=0.5),
        name="Dragonbox",
    ))
    fig.add_trace(go.Scatter(
        x=df["bits"],
        y=df["ryu_ns"],
        mode='markers',
        marker=dict(size=1, opacity=0.5),
        name="Ryu",
    ))

    mean = df.mean()
    median = df.median()
    std = df.std()

    text = ""
    text += "dragonbox:<br>"
    text += f"mean={mean['dragonbox_ns']:.2f} median={median['dragonbox_ns']:.2f} std={std['dragonbox_ns']:.2f}"
    text += "<br>"
    text += "ryu:<br>"
    text += f"mean={mean['ryu_ns']:.2f} median={median['ryu_ns']:.2f} std={std['ryu_ns']:.2f}"

    fig.add_annotation(
        text=text,
        x=0.01, y=0.99, # top left
        xref="paper", yref="paper",
        showarrow=False,
        align="left",
        font=dict(size=11),
        bordercolor="black",
        borderwidth=1,
        borderpad=6,
        bgcolor="white",
    )

    fig.update_layout(
        title=f"Dragonbox vs Ryu Performance ({bit_size}-bit)",
        xaxis_title="Bit Patterns",
        yaxis_title="Time (ns)",
        legend=dict(itemsizing="constant"),
        yaxis=dict(range=[30, 70]),
        xaxis=dict(
            range=[0, 2**bit_size],
            tickvals=[i*2**(bit_size-2) for i in range(1, 4+1)],
            ticktext=[f"{i}×2<sup>{bit_size-2}</sup>" for i in range(1, 4+1)],
        )
    )
    fig.write_image(output_file, width=1200, height=675)


for bit_size in [32, 64]:
    csv_file = f"random_digits{bit_size}.csv"
    output_file = f"random_digits{bit_size}.png"

    df = pd.read_csv(csv_file)

    mean = df.groupby("digits").mean().reset_index()
    median = df.groupby("digits").median().reset_index()
    q25 = df.groupby("digits").quantile(0.25).reset_index()
    q75 = df.groupby("digits").quantile(0.75).reset_index()

    fig = go.Figure()
    fig.add_trace(go.Scatter(
        x=mean["digits"],
        y=mean["dragonbox_ns"],
        mode="lines",
        name="Dragonbox Mean",
        line=dict(color="blue", width=1),
    ))
    fig.add_trace(go.Scatter(
        x=median["digits"],
        y=median["dragonbox_ns"],
        mode="lines",
        name="Dragonbox Median",
        line=dict(color="blue", dash="dot", width=1),
    ))
    fig.add_trace(go.Scatter(
        # reverse q75 to create polygon
        x=pd.concat([q25["digits"], q75["digits"][::-1]]),
        y=pd.concat([q25["dragonbox_ns"], q75["dragonbox_ns"][::-1]]),
        fill="toself",
        fillcolor="rgba(31,119,180,0.2)", # blue
        line=dict(color="rgba(255,255,255,0)"), # transparent
        name="Dragonbox IQR",
        showlegend=True,
    ))

    fig.add_trace(go.Scatter(
        x=mean["digits"],
        y=mean["ryu_ns"],
        mode="lines",
        name="Ryu Mean",
        line=dict(color="orange", width=1),
    ))
    fig.add_trace(go.Scatter(
        x=median["digits"],
        y=median["ryu_ns"],
        mode="lines",
        name="Ryu Median",
        line=dict(color="orange", dash="dot", width=1),
    ))
    fig.add_trace(go.Scatter(
        # reverse q75 to create polygon
        x=pd.concat([q25["digits"], q75["digits"][::-1]]),
        y=pd.concat([q25["ryu_ns"], q75["ryu_ns"][::-1]]),
        fill="toself",
        fillcolor="rgba(255, 127, 14, 0.2)", # orange
        line=dict(color="rgba(255, 255, 255, 0)"), # transparent
        name="Ryu IQR",
        showlegend=True,
    ))

    fig.update_layout(
        title=f"Dragonbox vs Ryu Performance ({bit_size}-bit)",
        xaxis_title="Number of Digits",
        yaxis_title="Time (ns)",
        legend=dict(itemsizing="constant"),
        yaxis=dict(range=[30, 70]),
        xaxis=dict(range=[1, 17]),
    )
    fig.write_image(output_file, width=1200, height=675)
