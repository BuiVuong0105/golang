package pattern

type Position struct {
	lat float64
	lon float64
}

func (p *Position) Move(lat, lon float64) {
	p.lat += lat
	p.lon += lon
}

func (p *Position) Teleport(lat, lon float64) {
	p.lat = lat
	p.lon = lon
}

type Player struct {
	*Position
}

type Enmery struct {
	*Position
}

func createPlayer() *Player {

	return &Player{
		Position: &Position{
			lat: 6,
			lon: 8,
		},
	}

	// player := Player{}
	// player.lat = 6
	// player.lon = 8

	// return &player
}
