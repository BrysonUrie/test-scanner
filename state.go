package main

type state int

const (
	s0 state = iota
	s1
	s2
	s3
	se
)

var stateMap = [][]state{
	{s1, se}, // s0
	{se, s2}, // s1
	{s3, se}, // s2
	{se, s2}, // s3
}
