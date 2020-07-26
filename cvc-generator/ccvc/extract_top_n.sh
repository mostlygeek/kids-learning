#!/bin/sh

# this extracts the top N (by number of words) from wordlist

NUM=10
if [ $# -eq 1 ]; then
    NUM=$1
fi

for ending in $(./unique-endings.sh | head -n $NUM); do
    grep "..$ending" wordlist-ccvc.txt
done
