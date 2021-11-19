package providerlog

import (
	"fmt"
	"io"
)

type ProviderLogWriter struct {
	w    io.Writer
	line []byte
}

func NewProviderLogWriter(w io.Writer) *ProviderLogWriter {
	return &ProviderLogWriter{w: w}
}

func (w *ProviderLogWriter) Write(p []byte) (n int, err error) {
	start := 0
	for i, b := range p {
		// TODO: This probably does not hold up against encoded UTF-8
		if b != '\n' {
			continue
		}

		w.line = append(w.line, p[start:i]...)
		fmt.Fprintf(w.w, "[DEBUG] %s\n", w.line)
		w.line = w.line[:]
		start = i
	}

	if start != len(p)-1 {
		w.line = append(w.line, p[start:len(p)-1]...)
	}
	return len(p), nil
}

func (w *ProviderLogWriter) Close() error {
	fmt.Fprintf(w.w, "[DEBUG] %s\n", w.line)
	return nil
}
