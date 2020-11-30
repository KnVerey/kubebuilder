#!/bin/sh

# hack script to generated pkged file because it has the wrong package name by default
rm pkged.go
go run github.com/markbates/pkger/cmd/pkger
mv ../pkged.go .
sed -i '' 's/package kubebuilder/package main/' ./pkged.go