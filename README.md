## gitsync

[![Go Report Card](https://goreportcard.com/badge/github.com/seibert-media/gitsync)](https://goreportcard.com/report/github.com/seibert-media/gitsync)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/513590eff4e54095a25b66bf65bd1323)](https://www.codacy.com/app/kwiesmueller/gitsync?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=seibert-media/gitsync&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/seibert-media/gitsync.svg?branch=master)](https://travis-ci.org/seibert-media/gitsync)
[![Docker Repository on Quay](https://quay.io/repository/seibertmedia/gitsync/status "Docker Repository on Quay")](https://quay.io/repository/seibertmedia/gitsync)

gitsync provides a service acting as middleware to sync the provided git repository when it's webhook gets called and then calls the next provided webhook while passing all return results back to the original caller


## Installing

Simply run `go get github.com/seibert-media/gitsync/cmd/gitsync`. There should now be a command `gitsync` in your `$GOPATH/bin`.  

## Usage

gitsync provides a `-help` flag which prints usage information and exits. 

## Dependencies

All dependencies inside this project are being managed by [dep](https://github.com/golang/dep) and are checked in.
After pulling the repository, it should not be required to do any further preparations aside from `make deps` to prepare the dev tools.

When adding new dependencies while contributing, make sure to add them using `dep ensure --add "importpath"`.

## Testing

To run tests you can use:
```bash
make test
# or
go test
# or
ginkgo -r
```

## Contributing

This application is developed using behavior driven development. 
Please keep in mind to use this development method when contribution to the project.

If you are new to BDD have a look at this video which inspired this project to use BDD:
 
https://www.youtube.com/watch?v=uFXfTXSSt4I

Feedback and contributions are highly welcome. Feel free to file issues or pull requests.

## Attributions

* [Kolide for providing `kit`](https://github.com/kolide/kit)
