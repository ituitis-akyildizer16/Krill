# krill fish completion
# source this file: source completions/krill.fish

complete -c krill -f

complete -c krill -n "__fish_use_subcommand" -a ask -d "ask a question about the current repo"
complete -c krill -n "__fish_use_subcommand" -a suggest -d "suggest a shell command"
complete -c krill -n "__fish_use_subcommand" -a review -d "summarize the current diff"
