package main

import (
	"math"
)

// The DoTurn function is where your code goes. The PlanetWars object contains
// the state of the game, including information about all planets and fleets
// that currently exist. Inside this function, you issue orders using the
// pw.IssueOrder() function. For example, to send 10 ships from planet 3 to
// planet 8, you would say pw.IssueOrder(3, 8, 10).
//
// There is already a basic strategy in place here. You can use it as a
// starting point, or you can throw it out entirely and replace it with your
// own. Check out the tutorials and articles on the contest website at
// http://www.ai-contest.com/resources.

func DoTurn(pw *PlanetWars) {
	//If we already have a fleet in flight, do nothing
	if len(pw.MyFleets()) > 1 {
		return
	}

	//Otherwise find my strongest planet...
	mp := pw.MyPlanets()
	mships, source := -1, -1

	for _, p := range mp {
		if p.NumShips > mships {
			mships = p.NumShips
			source = p.ID
		} 
	}

	//And the weakest enemy or neutral planet...
	tp := pw.NotMyPlanets()
	tships, dest := math.MaxInt32, -1

	for _, p := range tp {
		if p.NumShips < tships {
			tships = p.NumShips
			dest = p.ID
		} 
	}

	pw.IssueOrder(source, dest, mships / 2)
}
