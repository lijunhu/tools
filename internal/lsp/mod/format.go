package mod

import (
	"context"

	"golang.org/x/tools/internal/lsp/protocol"
	"golang.org/x/tools/internal/lsp/source"
	"golang.org/x/tools/internal/telemetry/trace"
)

func Format(ctx context.Context, snapshot source.Snapshot, fh source.FileHandle) ([]protocol.TextEdit, error) {
	ctx, done := trace.StartSpan(ctx, "mod.Format")
	defer done()

	file, m, err := snapshot.ModHandle(ctx, fh).Parse(ctx)
	if err != nil {
		return nil, err
	}
	formatted, err := file.Format()
	if err != nil {
		return nil, err
	}
	// Calculate the edits to be made due to the change.
	diff := snapshot.View().Options().ComputeEdits(fh.Identity().URI, string(m.Content), string(formatted))
	return source.ToProtocolEdits(m, diff)
}
