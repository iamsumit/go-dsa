package g

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

type Graph struct {
	Vertices []*Vertex
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(key int) *Vertex {
	v := &Vertex{Key: key}
	g.Vertices = append(g.Vertices, v)
	return v
}

func (g *Graph) AddEdge(v1, v2 *Vertex) {
	v1.Adjacent = append(v1.Adjacent, v2)
	v2.Adjacent = append(v2.Adjacent, v1)
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) GetVertices() []*Vertex {
	return g.Vertices
}

func (g *Graph) GetAdjacent(v *Vertex) []*Vertex {
	return v.Adjacent
}

func (g *Graph) GetAdjacentKeys(v *Vertex) []int {
	keys := make([]int, len(v.Adjacent))
	for i, a := range v.Adjacent {
		keys[i] = a.Key
	}
	return keys
}
