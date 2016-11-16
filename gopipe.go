// Implements a DAG data structure as a means to encapsulate a datapipeline and downstream
// channel communication
package zd

// TYPES ////////////////////////////////////////

type Weight interface{}

type Packet struct {
	Value interface{}
	Error error
}

type Stage interface{
	Process(in <-chan Packet) chan Packet, error
}

type GenericStage struct {
	fn func(in <-chan Packet) chan Packet
}

type Vertex struct {
  Edges []*Edge
  Data Stage 
}

type Edge struct {
  Origin *Vertex
  Target *Vertex
  Weight interface{}
}

type Pipeline struct {
  Root *Vertex
}

// METHODS //////////////////////////////////////

func (p *Pipeline) Go() error {
  
}

// executes lambda and something
func (s GenericStage) Process(in <-chan Packet) chan Packet, error {
  return s.fn(in)
}

func(v *Vertex) AddStage(func(in <-chan Packet) chan Packet) {

}


// STATIC ///////////////////////////////////////

// creates a new pipe and returns pointer to;
// NOTE: start with vertex or stage?? 
func NewPipe(s Stage) *Pipe {
  return &Pipe{ Vertex{ Data: s, } }
}

func newVertex(s Stage) *Vertex {
  v := Vertex{
    Edges: make([]*Edge, 5)
  }
}
