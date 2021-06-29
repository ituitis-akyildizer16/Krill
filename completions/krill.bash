# krill bash completion
# source this file: source completions/krill.bash

_krill() {
    local cur prev
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"

    local commands="ask suggest review status init models version"
    local files="--file -f --context-only --inline --shell --staged"

    case "$prev" in
        ask|suggest)
            COMPREPLY=( $(compgen -f -- "$cur") )
            return 0
            ;;
        -f|--file)
            COMPREPLY=( $(compgen -f -- "$cur") )
            return 0
