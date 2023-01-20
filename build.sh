#!/bin/sh
cd ./worker &&\
go build &&\
cd ../producer &&\
go build