#!/bin/bash
docker container run --rm -it -v $PWD:/data -u  1000 mfrw/latex
