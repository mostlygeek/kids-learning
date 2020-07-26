#!/bin/sh

# this extracts the vowel, consonsant at the end of the
# words to show the unique endings
#
sed 's/..\(..\)/\1/' wordlist-ccvc.txt | sort | uniq -c | sort -rn
