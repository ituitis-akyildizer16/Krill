# krill PowerShell completion
# Add to $PROFILE: . ./completions/krill.ps1

Register-ArgumentCompleter -Native -CommandName krill -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)

    $commands = @('ask', 'suggest', 'review', 'status', 'init', 'models', 'version')
    $tokens = $commandAst.CommandElements | ForEach-Object { $_.Extent.Text }
    $currentCommand = $tokens[1]

    if ($tokens.Count -le 2) {
        $commands | Where-Object { $_ -like "$wordToComplete*" } |
            ForEach-Object { [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_) }
        return
    }

    switch ($currentCommand) {
        'ask' {
            @('--file', '-f', '--context-only') |
                Where-Object { $_ -like "$wordToComplete*" } |
                ForEach-Object { [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterName', $_) }
        }
        'suggest' {
            @('--inline', '--shell') |
                Where-Object { $_ -like "$wordToComplete*" } |
                ForEach-Object { [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterName', $_) }
        }
    }
}