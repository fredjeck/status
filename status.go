package status

import (
	"fmt"
	"github.com/shiena/ansicolor"
	"os"
	"strconv"
)

// Not really usefull for now, we could use bare strings but would be
// cleaner if we'd like to expose the status.
type status string

const (
	// Status : pending (operation in progress)
	stPending = status("[....]")
	// Status : success (yes! we made it)
	stSuccess = status("[\x1b[1m\x1b[32mPASS\x1b[0m]")
	// Status : failure (woops something went wrong)
	stWarning = status("[\x1b[1m\x1b[33mWARN\x1b[0m]")
	// Status : failure (houston we have a problem here)
	stFailure = status("[\x1b[1m\x1b[31mFAIL\x1b[0m]")
)

var out = ansicolor.NewAnsiColorWriter(os.Stdout)

// Writer writes so called status messages in a *nix init fashion
// i.e : Message           [status]
// Writer provides utility methods to print such formatted messages (Success, Successf, Warn, Warnf...) as well as update methods
// which will only update the status part of the message, leaving the message untouched (MkSuccess, MkWarn...)
type Writer struct {
	lastMsg string // last printed message - used when updating a status
	lastSts status // last printed status
	width   int    // maximum message width (incl. status)
}

// NewWriter creates a new Writer configured to print messages limited to termWidth length
func NewWriter(termWidth int) *Writer {
	return &Writer{"", stPending, termWidth}
}

// Done prints out a newline character
func (w *Writer) Done() {
	fmt.Print("\n")
}

// Prints the status message and stores the issued message for a future update
func (w *Writer) printf(s status, message string, obj ...interface{}) {
	if obj == nil {
		w.lastMsg = message
	} else {
		w.lastMsg = fmt.Sprintf(message, obj...)
	}
	w.lastMsg = format(w.lastMsg, s, w.width)
	w.lastSts = s
	fmt.Printf("%v%v", w.lastMsg, s)
}

// Prints out the last message with the given status overwriting the last printed line
func (w *Writer) update(s status) {
	fmt.Fprintf(out, "\r%v%v", w.lastMsg, s)
}

// Format formats the status string's left part (message) according to the screen real estate.
// If the message is bigger than the available space, the message will be truncated adding the horizontal ellipsis as a truncation mark.
// If not, the message will be left padded with spaces.
func format(message string, s status, width int) string {
	// Screen real estate : available terminal width minus the length of the status.
	// We also substract one to leave a space right to the status.
	avw := width - len(s) - 1

	m := message
	if len(m) > avw {
		m = message[0:avw-1] + "…"
	}

	return fmt.Sprintf("%-"+strconv.Itoa(avw)+"v", m)
}

// ----------------------------
//      Helpers functions
//-----------------------------

// Success

// MkSuccess marks the last status message as being a success.
func (w *Writer) MkSuccess() *Writer {
	w.update(stSuccess)
	return w
}

// Success prints the given message with the success status
func (w *Writer) Success(message string) {
	w.printf(stPending, message)
}

// Successf prints the given message format string interpolated with obj.
// See the fmt.printf function for the format string details.
func (w *Writer) Successf(message string, obj ...interface{}) {
	w.printf(stPending, message, obj...)
}

// Warning

// MkWarning marks the last status message as being a warning.
func (w *Writer) MkWarning() *Writer {
	w.update(stWarning)
	return w
}

// Warning prints the given message with the success status
func (w *Writer) Warning(message string) {
	w.printf(stWarning, message)
}

// Warningf prints the passed message format string interpolated with obj.
// See the fmt.printf function for the format string details.
func (w *Writer) Warningf(message string, obj ...interface{}) {
	w.printf(stWarning, message, obj...)
}

// Failure

// MkFailure marks the last status message as being a failure.
func (w *Writer) MkFailure() *Writer {
	w.update(stFailure)
	return w
}

// Failure prints the given message with the failure status
func (w *Writer) Failure(message string) {
	w.printf(stFailure, message)
}

// Failuref prints the passed message format string interpolated with obj.
// See the fmt.printf function for the format string details.
func (w *Writer) Failuref(message string, obj ...interface{}) {
	w.printf(stFailure, message, obj...)
}

// Pending

// MkPending marks the last status message as being in progress.
func (w *Writer) MkPending() *Writer {
	w.update(stPending)
	return w
}

// Pending prints the given message with the "in progress" status
func (w *Writer) Pending(message string) {
	w.printf(stPending, message)
}

// Pendingf prints the passed message format string interpolated with obj.
// See the fmt.printf function for the format string details.
func (w *Writer) Pendingf(message string, obj ...interface{}) {
	w.printf(stPending, message, obj...)
}
