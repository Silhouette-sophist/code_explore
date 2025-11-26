package main

import (
	"context"
	"errors"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func main() {
	ctx := context.Background()

	const (
		nodeOfL1 = "invokable"
		nodeOfL2 = "streamable"
		nodeOfL3 = "transformable"
	)

	type testState struct {
		ms []string
	}

	gen := func(ctx context.Context) *testState {
		return &testState{}
	}

	sg := compose.NewGraph[string, string](compose.WithGenLocalState(gen))

	l1 := compose.InvokableLambda(func(ctx context.Context, in string) (out string, err error) {
		return "InvokableLambda: " + in, nil
	})

	l1StateToInput := func(ctx context.Context, in string, state *testState) (string, error) {
		state.ms = append(state.ms, in)
		return in, nil
	}

	l1StateToOutput := func(ctx context.Context, out string, state *testState) (string, error) {
		state.ms = append(state.ms, out)
		return out, nil
	}

	_ = sg.AddLambdaNode(nodeOfL1, l1,
		compose.WithStatePreHandler(l1StateToInput), compose.WithStatePostHandler(l1StateToOutput))

	l2 := compose.StreamableLambda(func(ctx context.Context, input string) (output *schema.StreamReader[string], err error) {
		outStr := "StreamableLambda: " + input

		sr, sw := schema.Pipe[string](utf8.RuneCountInString(outStr))

		// nolint: byted_goroutine_recover
		go func() {
			for _, field := range strings.Fields(outStr) {
				sw.Send(field+" ", nil)
			}
			sw.Close()
		}()

		return sr, nil
	})

	l2StateToOutput := func(ctx context.Context, out string, state *testState) (string, error) {
		state.ms = append(state.ms, out)
		return out, nil
	}

	_ = sg.AddLambdaNode(nodeOfL2, l2, compose.WithStatePostHandler(l2StateToOutput))

	l3 := compose.TransformableLambda(func(ctx context.Context, input *schema.StreamReader[string]) (
		output *schema.StreamReader[string], err error) {

		prefix := "TransformableLambda: "
		sr, sw := schema.Pipe[string](20)

		go func() {

			defer func() {
				panicErr := recover()
				if panicErr != nil {
					logger.CtxInfof(ctx, "panic occurs: %v\n", err)
				}

			}()

			for _, field := range strings.Fields(prefix) {
				sw.Send(field+" ", nil)
			}

			for {
				chunk, err := input.Recv()
				if err != nil {
					if err == io.EOF {
						break
					}
					// TODO: how to trace this kind of error in the goroutine of processing sw
					sw.Send(chunk, err)
					break
				}

				sw.Send(chunk, nil)

			}
			sw.Close()
		}()

		return sr, nil
	})

	l3StateToOutput := func(ctx context.Context, out string, state *testState) (string, error) {
		state.ms = append(state.ms, out)
		logger.CtxInfof(ctx, "state result: ")
		for idx, m := range state.ms {
			logger.CtxInfof(ctx, "    %vth: %v", idx, m)
		}
		return out, nil
	}

	_ = sg.AddLambdaNode(nodeOfL3, l3, compose.WithStatePostHandler(l3StateToOutput))

	_ = sg.AddEdge(compose.START, nodeOfL1)

	_ = sg.AddEdge(nodeOfL1, nodeOfL2)

	_ = sg.AddEdge(nodeOfL2, nodeOfL3)

	_ = sg.AddEdge(nodeOfL3, compose.END)

	run, err := sg.Compile(ctx)
	if err != nil {
		logger.CtxInfof(ctx, "sg.Compile failed, err=%v", err)
		return
	}

	out, err := run.Invoke(ctx, "how are you")
	if err != nil {
		logger.CtxInfof(ctx, "run.Invoke failed, err=%v", err)
		return
	}
	logger.CtxInfof(ctx, "invoke result: %v", out)

	stream, err := run.Stream(ctx, "how are you")
	if err != nil {
		logger.CtxInfof(ctx, "run.Stream failed, err=%v", err)
		return
	}

	for {

		chunk, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.CtxInfof(ctx, "stream.Recv() failed, err=%v", err)
			break
		}

		logger.CtxInfof(ctx, "%v", chunk)
	}
	stream.Close()

	sr, sw := schema.Pipe[string](1)
	sw.Send("how are you", nil)
	sw.Close()

	stream, err = run.Transform(ctx, sr)
	if err != nil {
		logger.CtxInfof(ctx, "run.Transform failed, err=%v", err)
		return
	}

	for {

		chunk, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.CtxInfof(ctx, "stream.Recv() failed, err=%v", err)
			break
		}

		logger.CtxInfof(ctx, "%v", chunk)
	}
	stream.Close()
}
