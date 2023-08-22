#!/bin/bash

start_time=$(date +%s)
for i in {1..8}
do
    ./ken attach --datadir ../data --preload attack.js &
done
wait
end_time=$(date +%s)

# elapsed time with second resolution
elapsed=$(( end_time - start_time ))
echo "ELAPSED TIME: $elapsed"
