#!/bin/sh

# expect ../generate binary to exist (pre-compiled).  Build it with
# go build generate.go before hand

set -e

cd $(dirname $0)

# generates sets of pdfs
#./top-start-end.sh | grep -E '^(sh|ch|sn)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set1_sh_ch_sn.pdf
#./top-start-end.sh | grep -E '^(wh|st|fl)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set2_wh_st_fl.pdf
#./top-start-end.sh | grep -E '^(sl|tr|sp)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set3_sl_tr_sp.pdf

./top-start-end.sh | grep -E '^(sh|ch)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set1_sh_ch.pdf
./top-start-end.sh | grep -E '^(wh|st)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set2_wh_st.pdf
./top-start-end.sh | grep -E '^(sl|tr)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set3_sl_tr.pdf
./top-start-end.sh | grep -E '^(sn|fl)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set4_sn_fl.pdf
./top-start-end.sh | grep -E '^(ch|sp)' | sed 's/\(..\)\(..\)/\1\2/' | ../generate set5_ch_sp.pdf
