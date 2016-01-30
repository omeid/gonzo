package gonzo

import "github.com/omeid/gonzo/context"

func makestage(stage Stage, ctx context.Context, in <-chan File) Pipe {
	out := make(chan File)

	next, cancel := context.WithCancel(ctx)
	go func() {
		err := stage(ctx, in, out)
		close(out)
		if err != nil {
			cancel()
		}
	}()

	return pipe{files: out, context: next}
}
