
package mp4

import (
	"github.com/nareix/mp4/atom"
	"io"
)

const (
	H264 = 1
	AAC = 2
)

type Track struct {
	Type int
	SPS []byte
	PPS []byte
	TrackAtom *atom.Track
	r io.ReadSeeker

	sample *atom.SampleTable
	sampleIndex int

	sampleOffsetInChunk int64
	syncSampleIndex int

	dts int64
	sttsEntryIndex int
	sampleIndexInSttsEntry int

	cttsEntryIndex int
	sampleIndexInCttsEntry int

	chunkGroupIndex int
	chunkIndex int
	sampleIndexInChunk int
}

