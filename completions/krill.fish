# krill fish completion
# source this file: source completions/krill.fish

complete -c krill -f

complete -c krill -n "__fish_use_subcommand" -a ask -d "ask a question about the current repo"
complete -c krill -n "__fish_use_subcommand" -a suggest -d "suggest a shell command"
complete -c krill -n "__fish_use_subcommand" -a review -d "summarize the current diff"
complete -c krill -n "__fish_use_subcommand" -a status -d "show a repo digest"
complete -c krill -n "__fish_use_subcommand" -a init -d "write a default config file"
complete -c krill -n "__fish_use_subcommand" -a models -d "list models on the local Ollama server"
complete -c krill -n "__fish_use_subcommand" -a version -d "print version and build info"

