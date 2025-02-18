package main

import (
	"fmt"
)

func main() {
	// Example function executions
	fmt.Println(numTilePossibilities("AAB"))    // Expected output: 8
	fmt.Println(numTilePossibilities("AAABBC")) // Expected output: 188
	fmt.Println(numTilePossibilities("V"))      // Expected output: 1
}

func numTilePossibilities(tiles string) int {
	result := make(map[string]bool, 0)
	solveTilesProblem("", []rune(tiles), result)
	return len(result)
}

func solveTilesProblem(currentSequence string, tiles []rune, sequences map[string]bool) {
	if currentSequence != "" {
		sequences[currentSequence] = true
	}

	for i, t := range tiles {
		newTiles := make([]rune, len(tiles)-1)
		copy(newTiles[:i], tiles[:i])
		copy(newTiles[i:], tiles[i+1:])
		solveTilesProblem(currentSequence+string(t), newTiles, sequences)
	}

}
