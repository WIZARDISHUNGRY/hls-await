package stream

import "flag"

type flags struct {
	URL            string
	DumpHttp       bool
	VerboseDecoder bool
	AnsiArt        int
	Threshold      int
	Flicker        bool
	FastStart      int
	FastResume     bool
	DumpFSM        bool
	OneShot        bool
	Sixel          bool
}

func WithFlags() StreamOption {
	return func(s *Stream) error {
		s.flags = someFlags
		return nil
	}
}

func getFlags() flags {
	f := flags{}
	flag.BoolVar(&f.DumpHttp, "dump-http", false, "dumps http headers")
	flag.BoolVar(&f.VerboseDecoder, "verbose-decoder", false, "ffmpeg debuggging info")
	flag.IntVar(&f.AnsiArt, "ansi-art", 0, "output ansi art on modulo frame")
	flag.IntVar(&f.Threshold, "threshold", 2, "need this much to output a warning")
	flag.BoolVar(&f.Flicker, "flicker", false, "reset terminal in ansi mode")
	flag.IntVar(&f.FastStart, "fast-start", 1, "start by only processing this many recent segments")
	flag.BoolVar(&f.FastResume, "fast-resume", true, "if we see a bunch of new segments, behave like fast start")
	flag.BoolVar(&f.DumpFSM, "dump-fsm", false, "write graphviz src and exit")
	flag.BoolVar(&f.OneShot, "one-shot", true, "render an ansi frame when entering up state")
	flag.BoolVar(&f.Sixel, "sixel", false, "output ansi images as sixels")
	return f
}

var someFlags = getFlags()
