# Dragonbox Benchmarks

Benchmark, fuzz test, and profile the Dragonbox algorithm using a locally built Go compiler and standard library.

## Setup

Clone the repository with its submodule:

```console
git clone --recurse-submodules git@github.com:taichimaeda/dragonbox-bench.git
```

Build the Go compiler and standard library (`go1.25.0` or later):

```console
cd ./go/src
./make.bash
```

### Benchmarking

Run benchmarks using your local Go build:

```console
export PATH="$PWD/go/bin:$PATH"
cd ./src/bench
go run main.go
```

To visualize the results:

```console
cd ./src/bench
pip install -r requirements.txt
python plot.py
```

### Fuzz Testing

Run fuzz tests against Ryu to check correctness of the implementation:

```console
export PATH="$PWD/go/bin:$PATH"
cd ./src/fuzz
go run main.go
```

### Profiling

Collect CPU and memory usage data:

```console
export PATH="$PWD/go/bin:$PATH"
cd ./src/profile
go run main.go
```

## Results

The following benchmarks and profiles were obtained on an M1 MacBook Air (8 cores/16GB memory).

Each chart shows performance of the Dragonbox algorithm under different input scenarios:

![](./src/bench/random_bits32.png)  
![](./src/bench/random_bits64.png)  
![](./src/bench/random_digits32.png)  
![](./src/bench/random_digits64.png)

![](./src/profile/random_bits32.png)
![](./src/profile/random_bits64.png)
![](./src/profile/random_digits32.png)
![](./src/profile/random_digits64.png)