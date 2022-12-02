package main

import (
	"bufio"
	"os"
)

func doPart1() int {
	plays := p1_readFile()
	points := 0
	for _, play := range plays {
		points += p1_calculatePoints(play)
	}
	return points
}

func p1_readFile() []Play {
	readFile, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)
	plays := []Play{}

	for scanner.Scan() {
		line := scanner.Text()
		plays = append(plays, p1_buildPlay(line))
	}

	readFile.Close()
	return plays
}

func p1_buildPlay(s string) Play {
	return Play{
		opponent: readOpponent(s[0]),
		player:   readPlayer(s[2]),
	}
}

func readPlayer(c byte) int {
	if c == 'X' {
		return ROCK
	} else if c == 'Y' {
		return PAPER
	} else if c == 'Z' {
		return SCISSORS
	}
	return INVALID
}

func p1_calculatePoints(p Play) int {
	pl := p.player
	op := p.opponent

	if pl == op {
		return pl + DRAW
	} else if pl == ROCK && op == PAPER {
		return pl + LOSE
	} else if pl == ROCK && op == SCISSORS {
		return pl + WIN
	} else if pl == PAPER && op == ROCK {
		return pl + WIN
	} else if pl == PAPER && op == SCISSORS {
		return pl + LOSE
	} else if pl == SCISSORS && op == ROCK {
		return pl + LOSE
	} else if pl == SCISSORS && op == PAPER {
		return pl + WIN
	}
	return 0
}
