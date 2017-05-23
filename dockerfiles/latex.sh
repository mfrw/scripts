#!/bin/bash
docker container run --rm -it -h latex -v $PWD:/data -u  1000 mfrw/latex
