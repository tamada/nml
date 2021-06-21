__nml() {
    local i cur prev opts cmds
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    prev="${COMP_WORDS[COMP_CWORD-1]}"
    cmd=""
    opts=""

    case "${prev}" in
        --unit | -u)
            COMPREPLY=($(compgen -W "nanoseconds microseconds milliseconds seconds minutes hours" -- "${cur}"))
            return 0
            ;;
    esac
    opts="-H --with-header -u --unit -h --help"
    if [[ "$cur" =~ ^\- ]]; then
        COMPREPLY=( $(compgen -W "${opts}" -- "${cur}") )
        return 0
    else
        compopt -o filenames
        COMPREPLY=($(compgen -d -- "$cur"))
    fi
}

complete -F __nml -o bashdefault -o default nml
