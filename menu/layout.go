package menu

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

// Handle building menu layout
func layoutMenu(w *tabwriter.Writer, cmds []CommandOption, width int) {
	fmt.Fprintln(w, "*\tCommand\tDescription\t")
	for i := range cmds {
		// Write command
		fmt.Fprintf(w, "*\t%s\t", cmds[i].Command)

		// Check description length
		description_length := len(cmds[i].Description)

		if description_length <= width {
			fmt.Fprintf(w, "%s\t\n", cmds[i].Description)
			continue
		}

		if description_length > width {
			layoutLongDescription(w, cmds[i].Description, width)
		}

	}
	fmt.Fprintln(w)
	w.Flush()
}

// Return tokens up cumulative maxsize
func getDescriptionRange(tokens []string, start int, maxsize int) ([]string, int) {
	total := 0
	token_part := tokens[start:]
	for i := range token_part {
		length := len(token_part[i])
		if total+length > maxsize {
			return token_part[0 : i-1], start + i
		}
		total = total + length
	}
	return token_part[0:], -1
}

func layoutLongDescription(w *tabwriter.Writer, d string, width int) {

	// Tokenize description
	tokens := strings.Fields(d)

	// Get description for range
	description, lastIndex := getDescriptionRange(tokens, 0, width)

	// Write first MAX_LENGTH of description
	fmt.Fprintf(w, "%s\t\n", strings.Join(description, " "))

	for {
		if lastIndex == -1 {
			break
		}

		description, lastIndex = getDescriptionRange(tokens, lastIndex, width)
		fmt.Fprintf(w, "*\t\t%s\t\n", strings.Join(description, " "))
	}

}
