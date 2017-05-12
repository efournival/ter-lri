# Installation instructions for DANSE

1. `git clone --recursive https://github.com/efournival/ter-lri.git`
2. `cd ter-lri/go-numeric-monoid && CGO_CPPFLAGS='-DMAX_GENUS=35' go install` (replace `35` by the desired genus)
3. `cd ../danse`
4. On each worker, setup a `config.json` (an example is provided as `config.json.sample`) with only the master machine's IP address
5. On the master machine, put all the workers' IP addresses in the configuration file

## Starting the computation
Change the maximum genus to compute at [the line 15](https://github.com/efournival/ter-lri/blob/master/danse/danser.go#L15) of `danser.go`.

Go binaries are statically linked and everything is included. Consequently, build DANSE with `go build` and copy the generated binary on each machine (workers and master).

Launch the workers with `./danse` and the master with `./danse --master`.
