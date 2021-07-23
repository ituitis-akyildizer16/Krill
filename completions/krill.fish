# krill fish completion
# source this file: source completions/krill.fish

complete -c krill -f

complete -c krill -n "__fish_use_subcommand" -a ask -d "ask a question about the current repo"
