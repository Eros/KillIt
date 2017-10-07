local cmd output
cmd=("$words[@]" --_shell-completion zsh:$CURRENT)
output=$("$cmd[@]" 2>/dev/null)

if [[ $output == "#compdef "* ]]; then
    # Looks like we got a valid completion function - so eval it to produce
    # the completion matches.
    eval $output
else
    echo "\nCompletion error running command:" ${(qqq)cmd}
    echo -n "If output below is unhelpful you may need to edit this file and "
    echo    "redirect stderr to a file."
    echo "Expected completion function, but instead got:"
    echo $output
    return 1
fi
