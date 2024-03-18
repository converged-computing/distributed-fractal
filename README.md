# Distributed Fractale

I wanted a distributed, grpc-based application for testing compatibility metadata with a scheduler in the cloud.
I wanted it to improve with strong scaling, and have network as a variable. I think Go is also interesting
for how it uses memory and garbage collection, etc.
So here we have, distributed fractal! This is so much cooler (and less annoying) than LAMMPS. Sorry LAMMPS...

![mandelbrot.png](mandelbrot.png)

## Usage

The default commands in the [Makefile](Makefile) for `make leader` and `make worker` generate metrics, primarily for the leader.
This is what the output looks like for the leader after a run is finished:

```bash
fractal leader --metrics --quiet
```
```console
METRICS LEADER time: 38.35129787s
METRICS LEADER /cgo/go-to-c-calls:calls: 1
METRICS LEADER /cpu/classes/gc/mark/assist:cpu-seconds: 0.003188
METRICS LEADER /cpu/classes/gc/mark/dedicated:cpu-seconds: 0.056393
METRICS LEADER /cpu/classes/gc/mark/idle:cpu-seconds: 0.106529
METRICS LEADER /cpu/classes/gc/pause:cpu-seconds: 0.015728
METRICS LEADER /cpu/classes/gc/total:cpu-seconds: 0.181838
METRICS LEADER /cpu/classes/idle:cpu-seconds: 454.936002
METRICS LEADER /cpu/classes/scavenge/assist:cpu-seconds: 0.000000
METRICS LEADER /cpu/classes/scavenge/background:cpu-seconds: 0.000000
METRICS LEADER /cpu/classes/scavenge/total:cpu-seconds: 0.000001
METRICS LEADER /cpu/classes/total:cpu-seconds: 499.728416
METRICS LEADER /cpu/classes/user:cpu-seconds: 44.610576
METRICS LEADER /gc/cycles/automatic:gc-cycles: 19
METRICS LEADER /gc/cycles/forced:gc-cycles: 0
METRICS LEADER /gc/cycles/total:gc-cycles: 19
METRICS LEADER /gc/heap/allocs-by-size:bytes: 9.000000
METRICS LEADER /gc/heap/allocs:bytes: 3800955080
METRICS LEADER /gc/heap/allocs:objects: 13332537
METRICS LEADER /gc/heap/frees-by-size:bytes: 9.000000
METRICS LEADER /gc/heap/frees:bytes: 3419495792
METRICS LEADER /gc/heap/frees:objects: 12665006
METRICS LEADER /gc/heap/goal:bytes: 411198456
METRICS LEADER /gc/heap/objects:objects: 667531
METRICS LEADER /gc/heap/tiny/allocs:objects: 37795374
METRICS LEADER /gc/limiter/last-enabled:gc-cycle: 0
METRICS LEADER /gc/pauses:seconds: 0.000025
METRICS LEADER /gc/stack/starting-size:bytes: 2048
METRICS LEADER /memory/classes/heap/free:bytes: 32071680
METRICS LEADER /memory/classes/heap/objects:bytes: 381459288
METRICS LEADER /memory/classes/heap/released:bytes: 294912
METRICS LEADER /memory/classes/heap/stacks:bytes: 4423680
METRICS LEADER /memory/classes/heap/unused:bytes: 5375144
METRICS LEADER /memory/classes/metadata/mcache/free:bytes: 1200
METRICS LEADER /memory/classes/metadata/mcache/inuse:bytes: 14400
METRICS LEADER /memory/classes/metadata/mspan/free:bytes: 189920
METRICS LEADER /memory/classes/metadata/mspan/inuse:bytes: 1017760
METRICS LEADER /memory/classes/metadata/other:bytes: 16400736
METRICS LEADER /memory/classes/os-stacks:bytes: 0
METRICS LEADER /memory/classes/other:bytes: 2213187
METRICS LEADER /memory/classes/profiling/buckets:bytes: 4957
METRICS LEADER /memory/classes/total:bytes: 443466864
METRICS LEADER /sched/gomaxprocs:threads: 12
METRICS LEADER /sched/goroutines:goroutines: 7
METRICS LEADER /sched/latencies:seconds: 0.000000
METRICS LEADER /sync/mutex/wait/total:seconds: 0.000440
```

I haven't found a good way to ping the worker that everything is done (to output metrics) but probably there is a good way -
I thought maybe the total tasks could at least be sent over to the node, but I'm not sure where I'd keep them (there is no state).
I'm thinking maybe they could be streamed over and then retrieved before exit.

Note that you can read about the metrics [here](https://go.dev/src/runtime/metrics/description.go) - some are indeed cumulative!
Note that this example shows running in non quiet mode, and without metrics. The default in the Makefile uses quiet mode
and metrics. Run the leader:

```bash
make leader
```

And as many workers as you like:

```console
# in separate terminals
make worker
make worker
```

And watch the generation happen! Here is the leader:

```console
$ make leader
mkdir -p /home/vanessa/Desktop/Code/distributed-fractal/bin
GO111MODULE="on" go build -o /home/vanessa/Desktop/Code/distributed-fractal/bin/fractal ./cmd/fractal/fractal.go
/home/vanessa/Desktop/Code/distributed-fractal/bin/fractal leader

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /start                    --> github.com/converged-computing/distributed-fractal/pkg/core.(*Leader).Init.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :9092
Rendering image....[GIN] 2024/03/16 - 16:51:25 | 200 |  5.689972138s |             ::1 | POST     "/start"
...

Mandelbrot set rendered into `mandelbrot.png`
```

And a worker (you'll notice they receive different shards of work, and one can finish before another):

```console
Received work: xmin:-0.0655945 ymin:0.741986625 xmax:0.060094499999999995 ymax:0.8362533750000001 iy:5900 iters:800 width:8192
Received work: xmin:-0.0655945 ymin:0.741986625 xmax:0.060094499999999995 ymax:0.8362533750000001 iy:5901 iters:800 width:8192
```

And right now I'm pinging an http endpoint to start, likely I can add params here instead:

```console
curl -X POST http://localhost:9092/start
```

And then you get the beautiful image after the workers finish! They are designed to process one row (Y) of the height at a time (across X).
I tried it first by pixel and it was horrificly slow, heh. To use in a headless environment and force exit:

```bash
fractal leader --force-exit
fractal leader --force-exit || true
```

I'm sure there are better ways to do that, I'm just too tired to figure it out fully now.

## Images

Here are some to generate:

```bash
mkdir -p ./img
./bin/fractal leader --palette "Hippi" --xpos -0.0091275 --ypos 0.7899912 --escape-radius .01401245 --outfile "img/mandelbrot.png"
./bin/fractal leader --palette "Plan9" --xpos -0.0091275 --ypos 0.7899912 --escape-radius .01401245 --outfile "img/test2.png"
./bin/fractal leader --palette "Vivid" --xpos -0.00991275 --ypos 0.7899912 --escape-radius .02401245 --outfile "img/test3.png" --iters 800 --step 600 --smoothness 10 --width 1920 --height 1080
./bin/fractal leader --palette "Hippi" --xpos -0.00275 --ypos 1.012 --escape-radius .089999 --outfile "img/test4.png" --iters 800 --step 600 --smoothness 10 --width 1920 --height 1080
./bin/fractal leader --palette "Hippi" --xpos -0.00275 --ypos 0.78912 --escape-radius .1256789 --outfile "img/test5.png" --iters 800 --step 6000 --smoothness 10 --width 1920 --height 1080
./bin/fractal leader --palette "AfternoonBlue" --xpos -0.0091275 --ypos 0.7899912 --escape-radius .01401245 --outfile "img/test6.png"
./bin/fractal leader --palette "SummerBeach" --xpos -0.0091275 --ypos 0.7899912 --escape-radius .01401245 --outfile "img/test7.png"
./bin/fractal leader --palette "Biochemist" --xpos -0.0091275 --ypos 0.7899912 --escape-radius .01401245 --outfile "img/test8.png" --smoothness 10
```

## Thank You

 - I learned about the simple distributed system architecture from [this post](https://dev.to/tikazyq/golang-in-action-how-to-implement-a-simple-distributed-system-2n0n)
 - The fractale I added is derived from [here](https://github.com/esimov/gobrot) (it's beautiful) ([MIT LICENSE](https://github.com/esimov/gobrot/commit/078c5bc391a187fea2d0663d8f4192732e61869e))

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/rainbow/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/rainbow/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/rainbow/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614

