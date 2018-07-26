package gosshauth

type bash struct{}

// BASH is a shell `bash`.
var BASH = bash{}

func (bash) Export(p string) string {
	return "export SSH_AUTH_SOCK=" + p
}

func (bash) Hook() string {
	return `_sshauthsock_hook() {
    local previous_exit_status=$?;
    eval "$("` + Me() + `" fixup bash)";
    return $previous_exit_status;
}
if ! [[ "$PROMPT_COMMAND" =~ _sshauthsock_hook ]]; then
    PROMPT_COMMAND="_sshauthsock_hook;$PROMPT_COMMAND";
fi
`
}
