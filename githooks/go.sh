#!/bin/sh
# Partly copy from https://tip.golang.org/misc/git/pre-commit
# Copyright 2012 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# git gofmt pre-commit hook
#
# To use, store as .git/hooks/pre-commit inside your repository and make sure
# it has execute permissions.
#
# This script does not handle file names that contain spaces.

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "$gofiles" ] && exit 0

echo "Running goimports check"
# goimports `go get golang.org/x/tools/cmd/goimports`
unformatted=$(goimports -local go-common -l $gofiles)
if [ ! -z "$unformatted" ]; then
    echo >&2 "Go files must be formatted with goimports. Please run:"
    for fn in $unformatted; do
        echo >&2 "  goimports -local go-common -w $PWD/$fn"
    done
    exit 1
fi

echo "Running gofmt check"
# gofmt
unformatted=$(gofmt -s -l $gofiles)
if [ ! -z "$unformatted" ]; then
    # Some files are not gofmt'd. Print message and fail.

    echo >&2 "Go files must be formatted with gofmt. Please run:"
    for fn in $unformatted; do
        echo >&2 "  gofmt -s -w $PWD/$fn"
    done
    exit 1
fi

# go get -u golang.org/x/lint/golint
echo "Running golint check"
# set -min_confidence 0.81 to disable abbreviation check..
golint -min_confidence 0.81 -set_exit_status $gofiles
ret=$?
if [ $ret -ne 0 ]; then
    echo "golint fail"
    exit 1
fi

exit 0
