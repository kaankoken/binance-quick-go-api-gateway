#!/bin/bash

rm -rf ./vendor
go mod tidy
go mod vendor
