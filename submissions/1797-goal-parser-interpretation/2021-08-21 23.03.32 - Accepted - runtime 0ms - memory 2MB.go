import "strings"

func interpret(command string) string {
    t := strings.ReplaceAll(command, "()", "o")
    return strings.ReplaceAll(t, "(al)", "al")
}