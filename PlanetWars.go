package main

import (
	"log"
	"fmt"
	"math"
	"bufio"
	"os"
	"container/vector"
)

const (
	NEUTRAL = iota
	ME
	)


func main() {
	in := bufio.NewReader(os.Stdin)
	lines := vector.StringVector(make([]string, 0, 500))

	for {
		for {
			l, err := in.ReadString('\n')
			
			if err != nil {
				log.Stderr("Error reading map, expected more input\n")
				return
			}

			if l == "go\n" {
				pw := ParseGameState(lines)
				DoTurn(pw)
				pw.EndTurn()		
				break
			} else {
				lines.Push(l)
			}
		}

		lines = lines[0:0]
	}
}




type Fleet struct {
	ID, Owner, NumShips, Source, Dest, TripLength, TurnsRemaining int 
}

func (this *Fleet) String() string {
	return fmt.Sprintf("F%d: %d %d %d %d %d\n", this.ID, this.Owner, this.NumShips, this.Source, this.Dest, this.TripLength, this.TurnsRemaining)
}

type Planet struct {
	ID int
	Owner int
	NumShips int
	GrowthRate int
	X float64
	Y float64
}

func (this *Planet) String() string {
	return fmt.Sprintf("P%d: %f %f %d %d %d\n", this.ID, this.X, this.Y, this.Owner, this.NumShips, this.GrowthRate)
}

type PlanetWars struct {
	Planets []*Planet
	Fleets []*Fleet
}


func (this *PlanetWars) MyPlanets() []*Planet {
	mp := make([]*Planet, len(this.Planets))
	
	i := 0
	
	for _, p := range this.Planets {
		if p.Owner == ME {
			mp[i] = p
			i++
		}
	}
	
	return mp[0:i]
}

func (this *PlanetWars) NotMyPlanets() []*Planet {
	mp := make([]*Planet, len(this.Planets))
	
	i := 0
	
	for _, p := range this.Planets {
		if p.Owner != ME {
			mp[i] = p
			i++
		}
	}
	
	return mp[0:i]
}

func (this *PlanetWars) EnemyPlanets() []*Planet {
	mp := make([]*Planet, len(this.Planets))
	
	i := 0
	
	for _, p := range this.Planets {
		if p.Owner > ME {
			mp[i] = p
			i++
		}
	}
	
	return mp[0:i]
}

func (this *PlanetWars) NeutralPlanets() []*Planet {
	mp := make([]*Planet, len(this.Planets))

	i := 0

	for _, p := range this.Planets {
		if p.Owner == NEUTRAL {
			mp[i] = p
			i++
		}
	}
	
	return mp[0:i]
}

func (this *PlanetWars) MyFleets() []*Fleet {
	mf := make([]*Fleet, len(this.Fleets))
	
	i := 0
	
	for _, f := range this.Fleets {
		if f.Owner == ME {
			mf[i] = f
			i++
		}
	}
	
	return mf[0:i]
}

func (this *PlanetWars) EnemyFleets() []*Fleet {
	mf := make([]*Fleet, len(this.Fleets))
	
	i := 0
	
	for _, f := range this.Fleets {
		if f.Owner > ME {
			mf[i] = f
			i++
		}
	}
	
	return mf[0:i]
}

func (this *PlanetWars) String() (rep string) {
	for _, p := range this.Planets {
		rep += p.String()
	}

	for _, f := range this.Fleets {
		rep += f.String()
	}

	return
}

func (this *PlanetWars) Distance(p1, p2 int) int {
	if p1 >= len(this.Planets) || p2 >= len(this.Planets) || p1 < 0 || p2 < 0 {
		log.Stderr("Distance called with out of range planet #\n")
		return -1
	}

	dx := this.Planets[p1].X - this.Planets[p2].X
	dy := this.Planets[p1].Y - this.Planets[p2].Y

	return int(math.Ceil(math.Sqrt(dx*dx + dy*dy)))
}

func (this *PlanetWars) IssueOrder(s, d, nShips int) {
	fmt.Printf("%d %d %d\n", s, d, nShips)
}

func (this *PlanetWars) EndTurn() {
	fmt.Print("go\n")
}

func (this *PlanetWars) IsAlive(playerID int) bool {
	for _, p := range this.Planets {
		if p.Owner == playerID {
			return true
		}
	}

	for _, f := range this.Fleets {
		if f.Owner == playerID {
			return true
		}
	}

	return false
}

func ParseGameState(lines []string) *PlanetWars {
	pw := &PlanetWars{make([]*Planet, len(lines)), make([]*Fleet, len(lines))}
	pNum, fNum := 0, 0

	for _, ln := range lines {

		switch ln[0] {
		case 'P' :
			p := &Planet{ID : pNum}
			read, e := fmt.Sscanf(ln[2:], "%f %f %d %d %d", &p.X, &p.Y, &p.Owner, &p.NumShips, &p.GrowthRate)

			if read < 5 || e != nil {
				log.Stderrf("Bad line in input: %s\n", ln)
			}
			pw.Planets[pNum] = p
			pNum++
		case 'F' :
			f := &Fleet{ID : fNum}
			read, e := fmt.Sscanf(ln[2:], "%d %d %d %d %d", &f.Owner, &f.NumShips, &f.Source, &f.Dest, &f.TripLength, &f.TurnsRemaining)

			if read < 5 || e != nil {
				log.Stderrf("Bad line in input: %s\n", ln)
			}
			pw.Fleets[fNum] = f
			fNum++
		default :
			log.Stderr("Error parsing gamestate: First char of line not 'P' or 'F'\n")
			return nil
		}
	}

	pw.Fleets = pw.Fleets[0:fNum]
	pw.Planets = pw.Planets[0:pNum]
	return pw
}
