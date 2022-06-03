Dijkstras algorithm over the US
Every state capital is a node
Every edge is a straight path to adjacent states capital
calculate the fastest route to the node from the node of choice
update each node you check
store the shortest distance from start node to that visited node
by adding shortest distance from current node
if the current nodes route to the visited node is less than its current
shortestDistance (starts -1 or something 
to indicate it hasnt been visited) then replace shortestDistance
Keep track of visited nodes. After checking each adjacent node,
go to next node unless its been visited


dijkstras algorithm jumps to next unvisited node if you hit a dead end,
or if all edges from that node have been visited
type WeightedGraph struct {
	StatesMap map([string]State)
	Routes []Route
}

type State struct {
	StateName string
	ShortestPath
	PreviousState
	Routes []*Route
}

type Route struct {
	CityA string
	CityB string
	Distance int
}

Build WeightedGraph
Add State
Add Routes
For each route, add pointer to route to CityA and
CityB Routes array in StatesMap

For each key in StatesMap
