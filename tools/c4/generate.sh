#!/bin/bash

export SAMPLE_ENV=fake:2137
export LOG_LEVEL=DEBUG

rm -rf out
mkdir out

go run .

if ! type "plantuml" > /dev/null; then
  echo "Please install plantuml to generate PNG diagrams automatically"
fi

for f in out/*.plantuml; do plantuml "$f"; done
