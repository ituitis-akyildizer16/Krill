# krill PowerShell completion
# Add to $PROFILE: . ./completions/krill.ps1

Register-ArgumentCompleter -Native -CommandName krill -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)

    $commands = @('ask', 'suggest', 'review', 'status', 'init', 'models', 'version')
    $tokens = $commandAst.CommandElements | ForEach-Object { $_.Extent.Text }
    $currentCommand = $tokens[1]
