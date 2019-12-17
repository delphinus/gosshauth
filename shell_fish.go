package main

type fish struct{}

// FISH is a shell `fish`.
var FISH = fish{}

func (fish) Export(p string) string {
	return "set -x SSH_AUTH_SOCK " + p
}
func (fish) Hook() string {
	return `function _sshauthsock_hook --on-event fish_prompt
  set -l previous_exit_status $status
  eval ("` + Me() + `" fixup fish)
  return $previous_exit_status
end`
}
