package rollinghash

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

// This algorithm first reads the originalFilename and updatedFilename versions of the file
// in chunks,computing the hash of each chunk using the SHA1 hash function.
// It then compares the hashes of the chunks from the originalFilename and updatedFilename files

const chunkSize = 1024 // chunk size in bytes

type chunk struct {
	hash  [sha1.Size]byte
	bytes []byte
}

// computeChunkHashes computes the hash of each chunk in the given file
func computeChunkHashes(file *os.File) ([]chunk, error) {
	chunks := make([]chunk, 0)
	buf := make([]byte, chunkSize)

	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		hash := sha1.Sum(buf[:n])
		chunks = append(chunks, chunk{hash: hash, bytes: buf[:n]})
		fmt.Println("number of chunks", len(chunks))
	}

	return chunks, nil
}

// ComputeDelta generates a description of the differences between the originalFilename and updatedFilename versions of the file
func ComputeDelta(original, updated *os.File) ([]byte, error) {
	buf := make([]byte, chunkSize)

	for {
		n, err := original.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	originalChunks, err := computeChunkHashes(original)
	if err != nil {
		return nil, err
	}
	updatedChunks, err := computeChunkHashes(updated)
	if err != nil {
		return nil, err
	}

	// keep track of the current position in each list of chunks
	originalPos := 0
	updatedPos := 0
	var delta []byte

	for updatedPos < len(updatedChunks) {
		// if the hashes of the current chunk from the originalFilename and updatedFilename files match,
		// it means the chunk can be reused and we can move to the next chunk in both lists
		if originalPos < len(originalChunks) && originalChunks[originalPos].hash == updatedChunks[updatedPos].hash {
			originalPos++
			updatedPos++
			continue
		}

		// otherwise, the chunk has been added or modified, so we need to add it to the delta
		delta = append(delta, updatedChunks[updatedPos].bytes...)
		updatedPos++
	}

	return delta, nil
}
