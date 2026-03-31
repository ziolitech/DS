- Types of graph:
 - Matrix Graph - a matrix item represents a node and shows where it is connected and where it is not.

 - Adj Matrix
 - Adj List
 - Edge Set - we are given set of edges {slice of pairs} and using that determine if src is connected to the destination.
 - Compositional Representation of Graph.

Tasks:
 - Create Graph Creation Code (✅)
 Next 
 - BFS and DFS.


 TODO:
  - Create Video to explain all these.


Problems (TODO):
1. No of provinces/connected components (both bfs and dfs works.) 
- Rotton Oranges (Interesting BFS - problem to find the min time to make oranges rotten in a grid.)
2. Flood fill algorithm with colors - can be done with bfs/dfs. (No constrainst like min time .. so DFS can also be used
   simple recursion way would work for this.)
3. Cycle detection in undirected Graph. (bfs) (✅)
4. Cycle detection in undirected Graph (dfs.) 
5. Distance of the nearest cell from '0' (✅)


TO Code:
6. surrounded regions. 👈 👨‍💻
7. Number of enclaves problem (similar to surrounded region problem) - we solve using BFS  👨‍💻 👈
8. word ladder 1:
 - brute force way .. where we change every char from a to z.
 - he analysed the brute force and then arrived the fact that .. we need to only care for the words which are there in the word list and thus .. choose that and move ahead ...
   - And due to change at every level we can say that we need to use the BFS traversal.
   - for each popped element .. you are changing each char and checking if that exists in the word list if yes then push again in the queue.
    ending point if queue is empty or we find a match.

9. word ladder 2: (one of the toughest problem of BFS ...)
   difference : print all the possible sequences from the start to the end.
   - Non competetive way
     - Performs the bfs with 2 key distinctions 1. Bfs is performed with keeping the sequence until now in the queue queue of slice
      - IMP : we don't delete from the word list until a certain level is completed as in that level there could be use of the same string with a different combination from the parent level.
      - in implementation track the sequences used in a level and then delete it from the wordlist.

   - Competetive prog way.
    - step 1 : Follow word ladder 1 (i.e find the min steps) and the store the steps where it occurred on the level in a map kind of structure.
    - step 2 : back track from end to begin (in the map which maintains the order.) to get the answer
          - basic idea is that .. we reduce the number of ways in which we need to reach to answer by storing in the map of the shortest path.


10. number of distinct islands...
 - Performed the BFS traversal to detect connected nodes in 8 directions and thus finally count the island. 
 - Used coordinates to push into the queue {x,y} and then pop them and then visit their neighbours in 8 directions.


11. bipartite graph. cycle detection in graph DFS
   - paint the graph with. 2 col in such a way that all the neighbour nodes have different color .. can only be done when you have odd nodes in the cycle.

12. Cycle detection in directed graph:
   - This is different from detecting a cycle in an undirected graph. The directed graph causes to visit the same node but they can visit the same node from a different path with the same origin and hence .. that simple dfs algo with path visit will not work . we need to take care of path as well for that we take a temp path visit array as well .. which we fill while visit the path and clear the elements while doing the back tracking.


------------------------- Topo Sort -------------------------

13:
 - Topo sort. (using DFS)
  - Keep on navigating a path deeply and keep on marking them as visited until it does not have anything to explore further thats when we put it in stack .. just to preserve the order in which we were navigating the path.

14. Topo sort (using slightly modified version of BFS - using the indegree array) - Kahn's Algorithm
 - First calculate the indegree of all the nodes using the adjacency list.
 - so the graph will start with nodes whose .. indegree is 0, remember there can be many nodes whose indegree is 0 .. we need to put all the nodes in the queue .. basically this will act as the starting point node. Put it in the queue. Now we will take out the element .. and then go to adj list and remove the edges temp to the outgoing nodes in that way we reduce the indegree of the nodes.

- while doing above if the indegree of a node reduces to 0 then we push it in the queue .. which means that this becomes the candidate for the next position in the Topo sort.
- this we keep doing until we have all the elements in the adjacency list.

15. Detect Cycle in a Graph using the topo sort 
   - This is performed using the khans algorithm if we detect that there is nothing to push n he qqueue this may be coz of the fact that there is a cycle.
   IMP Point : If the size of the topo sort is less than N then it indicates it has a cycle.


16:
 - Course Schedule I and II:
  - in this problem we are given the course schedule where course a needs to performed before course b. and similar dependent conditions .. and we in 1st probelem we needed to check if we can finish all the courses. - we use khan's algo to find one of the topo sort.
  In 2nd problem, we needed to find one of the topo sort with slight modification in th problem set.

17: (Rivisit .. it)
 - Eventual Safe States
  - Safe node - every possible path originiting from that node .. ends in a terminal node.
  - terminal node
From the YT comments:
  i just tried created such a graph where atleast 1 node i wanted to create in such a way that one of its path leads to a terminal node and one path leads to some other node. I realised any path that does not lead to terminal node will always have a cycle

   Fact:
   In a directed graph, a node is not safe if:
         - It cannot reach any terminal node, and
         - Every path from it eventually gets stuck in a cycle.

   Why?
   - Every path in a finite directed graph has only two possibilities:
   - It eventually stops at a terminal node (outdegree = 0).
   - It never stops — meaning it must loop → cycle.
   - There’s no third option, because the graph is finite.

- Read more comments for the intuition.

18. Alien's Dictionary
 - something before something indicates presence/or it being a candidate for the Topo sort.
 - express all chars in terms of the number and then based in analysis ... we need to create a directed graph of chars or numbers.
 - and to determine the order of the char we need to perform the topological sorting.

Core concept compare each pair and see .. which char appear before another pair.
 - if one char is unequal then we dont need to go further that becomes the differentiating factor.


Shortest Path Algorithms:

1. Shortest path in undirected Graph with unit weights.
 - BFS with storing node and dist and then keep a distance array as well with all having int_max or infintity as default val and then replacing those vals with the min distance each time while navigating through the BFS until the queue is empty.
 - Here DFS is not valid as it goes deep and can visit every node without giving the shortest path algorithm and in BFS it goes level by level distance by distance so node at d=1 is visited first as compared to node at d=2 and so on.

2. Shortest path in a Directed Acyclic Graph (DAG) with weights using the Topological sort.
  - step1 : Do a topo sort on the Graph using DFS and push elements to stack as the element coming last will be push first and the origin will be pushed on top of the stack.
  - step2 : You take out the nodes from the stack one by one and then update the distance array - here we keep track of the old distance and only update if the obtained distance via the current node is lesser.
  pre-req: Have a distance array with each node indicated by index having a distance of 'd' (intially we keep it as the int max). The source node can be set as 0 as the distance from the starting origin node to the end node is 0 itself.
  Here it gives you the shortest distance for each node.

3. Dijkstra's Algorithm :
 -  From a given source to all the destination in the graph what will be the shortest distance. 
 we can implement the Dijkshtra algo using three methods:
 1. Priority Queue. (PQ)
 2. Set (Fastest)
 3. Queue (this will take lots of time. - has unnecessary paths to cover)

 Dij algo may or may not work for the negative edges .. This just works for the edges with postive edges as it may visit a node and it wont allow to visit it if it is already visited.
  - Greedy Approach of selecting the edges. (Watch Abdul sir's lecture.)

- Dij says that find the node with the shortest path in a greedy way and update its total cost and the try to see if there are other nodes via this node where we can again relax the edges(find the shiortest path via that node.) means we can chose to get the shortest path.

if d[u] + c(u,v) < d[v] {
   d[v] = d[u] + c(u,v)
}


4. Bellman ford algo