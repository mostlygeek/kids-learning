# CVC word generator

This is a PDF worksheet generator that reads a list
of words from stdin and creates a 5 page PDF with a randomized
list of the words.

It reads from stdin so I can create wordlists with endings
to tune the practice and growth week to week.


* Generate PDF: `cat ad_it_ot.txt | go run generate.go ad_it_ot.pdf`
* Create wordlist: `grep -E '.(ed|ot)' wordlist.txt > ed_ot.txt`

