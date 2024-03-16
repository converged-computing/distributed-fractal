# Distributed Fractale

I wanted a distributed, grpc-based application for testing compatibility metadata with a scheduler in the cloud.
I wanted it to improve with strong scaling, and have network as a variable. I think Go is also interesting
for how it uses memory and garbage collection, etc.
So here we have, distributed fractal! This is so much cooler (and less annoying) than LAMMPS. Sorry LAMMPS...

![mandelbrot.png](mandelbrot.png)

## Usage

Run the leader:

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
I tried it first by pixel and it was horrificly slow, heh.

## TODO

 - vary parameters (via start)?

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

