# krill PowerShell completion
# Add to $PROFILE: . ./completions/krill.ps1

Register-ArgumentCompleter -Native -CommandName krill -ScriptBlock {
    param($wordToComplete, $commandAst, $cursorPosition)

