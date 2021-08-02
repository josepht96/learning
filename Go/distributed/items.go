package main

import "fmt"

//Song struct uh
type Song struct {
	ID     int
	Title  string
	Artist string
}

func (s Song) String() string {
	return fmt.Sprintf("Title: %s, Artist: %s, ID: %v", s.Title, s.Artist, s.ID)
}

var songs = []Song{
	{
		ID:     0,
		Title:  "Turtles All the Way Down",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     1,
		Title:  "Life of Sin",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     2,
		Title:  "Living the Dream",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     3,
		Title:  "Voices",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     4,
		Title:  "Long White Line",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     5,
		Title:  "The Promise",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     6,
		Title:  "A Little Light",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     7,
		Title:  "Just Let Go",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     8,
		Title:  "It Ain't All Flowers",
		Artist: "Sturgill Simpson",
	},
	{
		ID:     9,
		Title:  "Panbowl",
		Artist: "Sturgill Simpson",
	},
}
