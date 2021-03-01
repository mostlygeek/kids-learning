#!/bin/sh

# creates a list of words start or ending with the most common 2 letter combos
grep -E '(st|sl|sh|ch|fl|tr|sp|wh|sn)(ow|at|ew|ot|ay|ap|am|it|aw|ug|op|ug)' wordlist-ccvc.txt
