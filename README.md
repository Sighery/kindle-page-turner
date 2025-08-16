# Example Go application for kindle-bt-api

Example Go application making use of [`kindle-bt-api`][kindle-bt-api]. This
makes use of [CGO][]. This will require building `kindle-bt-api` and having the
proper shared object (`libkindlebt.so`), as well as access to the `kindlebt`
header files.

To build the object, as well as to get the headers, look into the
[`kindle-bt-api`][kindle-bt-api] library.

## Setting up the Go project

Cross compiling in Go is usually easy, but the Kindle setup requires a few extra
changes. The flow goes something like this:

1. Set up koxtoolchain
2. Set up kindle-sdk
3. Source the koxtoolchain environment
4. Set up needed environment variables

Even though this sounds easy, in practice this requires quite a few steps.
Setting up `koxtoolchain` and `kindle-sdk` are already covered in the
`kindle-bt-api` repository.

The last two steps, you can see a working documentation of in the
[`.envrc` file][envrc]. If you use [Direnv][], you can automate these steps
whenever you `cd` into the directory. If not, you can use this as reference of
the environment variables you'll need to set.

The most important ones are:

* source koxtoolchain: The path will depend on where you have downloaded the
  `koxtoolchain` repository into your system, as well as the target (the
  `.envrc` assumes `kindlehf`)
* set up `CC`: this should point to the `gcc` compiler binary for your target
* set up `CGO_CFLAGS`: compiler flags. This should point to the path where the
  `kindlebt` and `ace_bt` header files are in your system.
* set up `CGO_LDFLAGS`: linker flags. This should point to the path where
  `libkindlebt.so` is in your system

Instead of manually setting these values, you may also just edit the `.envrc`
file to have appropriate values for your system paths/Kindle target, and then
let direnv refresh the environment.

## Running on the Kindle

Once you have copied the binary into the Kindle, if `libkindlebt.so` is not yet
on the system, or if you have it on a non-standard path, you can manually force
the Go binary to look for objects at a specific path:

```sh
LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/mnt/us ./kindlebt_go
```

[kindle-bt-api]: https://github.com/Sighery/kindle-bt-api
[CGO]: https://go.dev/wiki/cgo
[envrc]: ./.envrc
[Direnv]: https://direnv.net/
