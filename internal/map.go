package internal

type Map struct {
	Tilesets      map[string]*Tileset
	TilesetsReady bool
	Chunks        map[int]map[int]*Chunk
	StartChunkX   int
	StartChunkY   int
}

func NewMap() *Map {
	m := &Map{
		Chunks:   make(map[int]map[int]*Chunk),
		Tilesets: make(map[string]*Tileset),
	}
	m.LoadTilesets()
	return m
}

func (m *Map) LoadTilesets() {
	m.Tilesets["main"] = NewTileset("./assets/TilesA2.png", 16, 16)
}

func (m *Map) GenerateChunk(x, y int) {
	_, yOK := m.Chunks[y]
	if !yOK {
		m.Chunks[y] = make(map[int]*Chunk)
	}

	_, xOK := m.Chunks[y][x]
	if !xOK {
		m.Chunks[y][x] = NewChunk(x, y)
	}
}

func (m *Map) Update() {

}

func (m *Map) Draw() {
	// draw chunks
	for y := m.StartChunkY; y < m.StartChunkY+4; y++ {
		for x := m.StartChunkX; x < m.StartChunkX+6; x++ {
			m.DrawChunk(m.Tilesets["main"], x, y)
		}
	}
}

func (m *Map) DrawChunk(tileset *Tileset, x, y int) {
	_, yOK := m.Chunks[y]
	if !yOK {
		m.Chunks[y] = make(map[int]*Chunk)
	}
	_, xOK := m.Chunks[y][x]
	if !xOK {
		m.Chunks[y][x] = NewChunk(x, y)
	}
	m.Chunks[y][x].Draw(tileset)
}
