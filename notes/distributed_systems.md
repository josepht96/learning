# CAP Theorem:
the CAP theorem, also named Brewer's theorem after computer scientist Eric Brewer, 
states that it is impossible for a distributed data store to simultaneously provide 
more than two out of the following three guarantees:
	Consistency: Every read receives the most recent write or an error
	Availablity: Every request receives a non-error response, without the guarantee that it 
		     contains most recent write
	Partition tolerance: The system continues to operate despite an arbitrary number of messages 
			     being dropped or delayed by the network between nodes
Primarily this is a choice between consistency and availability

# Patterns for Building Distributed Systems for The Enterprise
CAP Theorem
Theorem describes behavior of distributed system. Connection of interconnected nodes that all share data. Guarantees: Only 2 of 3 can be guaranteed: Consistency, Availability, Partition Tolerance. Consistency means system guarantees its Read data is as fresh as the data it just wrote (non-stale data). Availability means a non-failing node will giving the client a reasonable response within a reasonable amount of time. Meaning you can read and write relatively quickly. Partition Tolerance guarantees a distributed system will continue to function in the event of a network partition, nodes cannot communicate with each other. 
Fallacies of Distributed Systems: 
1.	Network is reliable
2.	Latency == 0
3.	Bandwidth is infinite
4.	Network is secure
5.	Topology doesn’t change
6.	There is one administrator 
7.	Transpiration cost is 0
8.	Network is homogenous
Networks on not reliable in the same sense that function calls are reliable. Requests, responses might fail, could be related in transportation issue or logic. Often hard to determine why something failed. Time is a major factor, while its not in OOP. 
Bandwidth has limitations. Passing large data packets has a higher chance of failing or passing incomplete. 
