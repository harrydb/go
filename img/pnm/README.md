Package pnm
===========

Package pnm implements a PBM, PGM and PPM image decoder and encoder.

This package is compatible with Go version 1.


### Installation

	go install github.com/harrydb/go/img/pnm


### Documention

See: https://gopkgdoc.appspot.com/pkg/github.com/harrydb/go/img/pnm


### Limitations

Not implemented are:

* Writing pnm files in raw format.
* Writing images with 16 bits per channel.
* Writing images with a custom Maxvalue.
* Reading/and writing PAM images.

(I would be happy to accept patches for these.)
