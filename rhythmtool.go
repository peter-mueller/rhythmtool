// Package rhythmtool provides a toolset for generating and manipulating simple
// rhythmns.
package rhythmtool

import "math/rand"

// Note is a simple type that describes one single unit of time.
// It can either be a HIT or a PAUSE.
type Note bool

// HIT represents an active (or played) Note.
const HIT Note = true

// PAUSE represents a inactive (or rested) Note.
const PAUSE Note = false

// bjorklundStore is a container for all information required in the
// Bjorklund algorithmn.
type bjorklundStore struct {
	Counts     []int
	Remainders []int
	Rhythm     Rhythm
}

// A Rhythm is a sequential arrangement of Notes.
type Rhythm []Note

// String can convert Notes to readable strings.
func (n Note) String() string {
	if n == HIT {
		return "x"
	}
	return "-"
}

// Random creates a Rhythm that contains random Notes.
// Note that the rand Seed can be changed.
func Random(length int) Rhythm {
	rhythm := make(Rhythm, length)
	for i := range rhythm {
		rhythm[i] = rand.Intn(2) == 0
	}
	return rhythm
}

// GenerateBjorklund creates a Rhythm with specified length that has
// the number of provided hits.
//
// It uses the Bjorklund algorithm to create euclidean rhythmns.
func GenerateBjorklund(pulses, length int) Rhythm {
	if pulses > length {
		panic("failed")
	}

	divisor := length - pulses
	remainders := []int{pulses}
	counts := []int{}
	level := 0

	for {
		counts = append(counts, divisor/remainders[level])
		remainders = append(remainders, divisor%remainders[level])
		divisor = remainders[level]
		level++

		if remainders[level] <= 1 {
			break
		}
	}
	counts = append(counts, divisor)

	store := bjorklundStore{
		Counts:     counts,
		Remainders: remainders,
	}
	store.buildBjorklund(level)
	if store.Rhythm[0] == PAUSE && store.Rhythm[len(store.Rhythm)-1] == HIT {
		store.Rhythm = store.Rhythm.Reverse()
	}
	return store.Rhythm
}

// Subdivide takes a Rhythm and splits it up to multiple Rhythmns based on the
// given subdivision.
//
// It panics, if n subdivisions do not fit perfectly into the length of
// the Rhythmn.
func (rhythm Rhythm) Subdivide(subdivision int) []Rhythm {
	if len(rhythm)%subdivision != 0 {
		panic("error")
	}
	var rhythms []Rhythm
	for i := 0; i < len(rhythm); i += subdivision {
		rhythms = append(rhythms, rhythm[i:i+subdivision])
	}
	return rhythms
}

// Rotate shifts a Rhythm to the right or left (negative offset).
func (rhythm Rhythm) Rotate(offset int) Rhythm {
	offset = offset * -1 % len(rhythm)
	if offset < 0 {
		offset += len(rhythm)
	}
	return append(rhythm[offset:], rhythm[:offset]...)
}

// Reverse flips the Rhythm. Like playing it backwards.
func (rhythm Rhythm) Reverse() Rhythm {
	for i := len(rhythm)/2 - 1; i >= 0; i-- {
		opp := len(rhythm) - 1 - i
		rhythm[i], rhythm[opp] = rhythm[opp], rhythm[i]
	}
	return rhythm
}

// MergeWith combines the Rhythm with the given rhythm.
func (r1 Rhythm) MergeWith(r2 Rhythm) Rhythm {
	r := make(Rhythm, max(len(r1), len(r2)))
	for i, note := range r1 {
		r[i] = note
	}
	for i, note := range r2 {
		r[i] = r[i] || note
	}
	return r
}

func (store *bjorklundStore) buildBjorklund(level int) {
	if level == -1 {
		store.Rhythm = append(store.Rhythm, PAUSE)
		return
	}
	if level == -2 {
		store.Rhythm = append(store.Rhythm, HIT)
		return
	}

	for i := 0; i < store.Counts[level]; i++ {
		store.buildBjorklund(level - 1)
	}
	if store.Remainders[level] != 0 {
		store.buildBjorklund(level - 2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
