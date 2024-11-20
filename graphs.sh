#!/bin/sh

for file in ./output/graph_*.gv; do
  dot -Tjpg "$file" > "${file%.gv}.jpg"
done