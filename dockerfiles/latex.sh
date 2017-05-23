#!/bin/bash
docker container run --rm --name latex -it -h latex -v $PWD:/data -u  1000 mfrw/latex
