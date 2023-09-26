#!/bin/bash

quest01=("how-2-js" "primitives" "declarations" "mutability" "more-or-less" "returns" "last-first-kiss" "concat-str" "change" "dog-years" "circulars" "biggie-smalls" "method-man" "abs" "min-max" "sign" "is" "physics" "collections" "block-chain")


quest02=("repeat" "elementary" "index-of" "slicer" "keep-cut" "cut-corners" "reverser" "unbreakable" "chunky" "nasa" "molecules-cells" "get" "triangle" "pyramid" "flat" "finc-expression" "sums")

quest03=("its-a-match" "letter-space-number" "vowels-dots" "same-amount" "ion-out" "greedy-url" "group-price" "valid-ip")

Raid1="crossword"

quest04=("dr-strange" "date-is" "get-some-time" "rebecca-blank" "match-cron" "count-leap-years" "day-leap-year" "bloody-sunday" "unicode-technical-report-35")

quest05=("using-map" "using-filter" "using-reduce" "long-words" "for-each" "mapper" "filter" "reduce" "quantifiers" "sweet-curry" "has-city" "flow" "currify")

quest06=("get-them-all" "build-brick-and-break" "harder-bigger-bolder-stronger" "pimp-my-style" "fifty-shades-of-cold" "keycodes-symphony" "mouse-trap" "pick-and-click" "gossip-grid" "where-do-we-go")

Raid2="sortable"

quest07=("invert" "pick-omit" "fusion" "manipulate-values" "manipulate-keys" "manipulate-entries" "curry-entries" "deep-copy" "replica" "pronoun" "neuron" "flagger")

quest08=("debounce" "throttle" "keep-trying-or-give-up" "get-json" "is-winner" "gougle-search" "all" "series" "race")

Raid3="clonernews"

quest09=("verydisco" "verydisco-forever" "tell-me-how-many" "tell-me-who" "tell-me-vip" "tell-it-cypher" "friend-support" "uninvited" "gatecrashers" "personal-shopper" "happiness-manager")

# mkdir quest09
# cd quest09

# for element in "${quest08[@]}"
# do
    # wget --no-verbose https://learn.zone01dakar.sn/git/root/public/raw/branch/master/subjects/${element}/README.md

    # mv README.md ${element}.md
# done

wget --no-verbose https://learn.zone01dakar.sn/git/root/public/raw/branch/master/subjects/crossword/README.md

mv README.md Raid1.md

wget --no-verbose https://learn.zone01dakar.sn/git/root/public/raw/branch/master/subjects/sortable/README.md

mv README.md Raid2.md

wget --no-verbose https://learn.zone01dakar.sn/git/root/public/raw/branch/master/subjects/clonernews/README.md

mv README.md Raid3.md
