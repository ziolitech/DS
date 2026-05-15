- Types of graph:
 - Matrix Graph - a matrix item represents a node and shows where it is connected and where it is not.
 - Adj Matrix
 - Adj List
 - Edge Set - we are given set of edges {slice of pairs} and using that determine if src is connected to the destination.
 - Compositional Representation of Graph. (Ref : https://medium.com/@snassr/graphs-with-go-golang-part-i-3e0f9392c294)

Tasks:
 - Create Graph Creation Code (✅)

 Traversals:
 - BFS 
  - Goes in levels - visits the nodes in distance from the starting node.

 - DFS.



 TODO:
  - Create Video to explain all these.


Problems (TODO):
1. No of provinces/connected components (both bfs and dfs works.) 

2. Rotton Oranges (Interesting BFS - problem to find the min time to make oranges rotten in a grid.)

3. Flood fill algorithm with colors - can be done with bfs/dfs. (No constraint like min time .. so DFS can also be used
   simple recursion way would work for this.)

4. Cycle detection in undirected Graph. (bfs) (✅) - TODO write notes
   - Have a vis array to store if a element is visited or not.
   - Put the elements in the queue along with the information  about the parent node where it came from.
   - Keep on visiting sibling .. if the sibling is not visited then visit them and put them in the queue with (node, parentNode) as data.
   - we need parentNode to not visit it via the adjacency list.

5. Cycle detection in undirected Graph (dfs.) - TODO write notes.
 - Here you call the DFS func from the start node and also pass the parent node while calling and then keep on visiting sibliings if the sibling is not visited then keep calling further .. the moment you found some sibling as visited - in case its non parent .. and its already visited then it indicates that there is cycle where this element was already in the path and its kind of closing the circle/loop in the cycle.

6. Distance of the nearest cell from '0' having 1 (✅)
 - main intuiton is starting with 1 to enque as we know the 3 vars (x, y, d = known - for 1 it is 0). If we enque 0 then we dont know the distance of 1 from this coordinate (x, y , d(unknown)).


To Code:
7. surrounded regions. 👈 👨‍💻 
- Intution - if a zero is on the boundry then all the connected zeros will never be converted to an X. the remaining zeros will be converted to an X.


8. Number of enclaves problem (similar to surrounded region problem) - we solve using BFS  👨‍💻 👈, can be solved using DFS also..by first enqueing.



######################################################## TODO : Focus later ########################################################

9. Word ladder 1:
 - brute force way .. where we change every char from a to z.
 - he analysed the brute force and then arrived the fact that .. we need to only care for the words which are there in the word list and thus .. choose that and move ahead ...
   - And due to change at every level we can say that we need to use the BFS traversal.
   - for each popped element .. you are changing each char and checking if that exists in the word list if yes then push again in the queue.
    ending point if queue is empty or we find a match.


10. word ladder 2: (one of the toughest problem of BFS ...)
   difference : print all the possible sequences from the start to the end.
   - Non competetive way
     - Performs the bfs with 2 key distinctions 1. Bfs is performed with keeping the sequence until now in the queue of slice
      - IMP : we don't delete from the word list until a certain level is completed as in that level there could be use of the same string with a different combination from the parent level.
      - in implementation track the sequences used in a level and then delete it from the wordlist.

   - Competetive prog way.
    - step 1 : Follow word ladder 1 (i.e find the min steps) and the store the steps where it occurred on the level in a map kind of structure.
    - step 2 : back track from end to begin (in the map which maintains the order.) to get the answer
          - basic idea is that .. we reduce the number of ways in which we need to reach to answer by storing in the map of the shortest path.


########################################################################################################################


11. number of distinct islands ...
 - Performed the BFS traversal (we can use DFS also here.) to detect connected nodes in 8 directions and thus finally count the island.
 - Used coordinates to push into the queue {x,y} to identify a vertex (can be encapsulated in a Data struct as well) and then pop them and then visit their neighbours in 8 directions.
 Nuances:
 - Maintain a 2D matrix/array to mark a node as visited.
 - Go through all the elements in the matrix to see if we have a unvisited 1 .. (this will help detect other clusters in the matrix.)


12. bipartite graph. cycle detection in graph DFS. (can be done with BFS also but probably more iteration .. BS with GPT.)
   - paint the graph with. 2 col in such a way that all the neighbour nodes have different color.
   Intuition:
    - For a linear graph without any cycle this can be easily done.
    - In case Graph has cycle there are 2 cases:
      - 1 : Graph has odd number of nodes - we can color in this case ... see and check as last will coinside with the first color
      - 2 : Graph has even number of nodes - No issues with coloring with 2 colors.


13. Cycle detection in directed graph:
   - This is different from detecting a cycle in an undirected graph. The directed graph causes to visit the same node but they can visit the same node from a different path with the same origin and hence .. that simple dfs algo with path visit will not work . we need to take care of path as well for that we take a temp path visit array as well .. which we fill while visit the path and clear the elements while doing the back tracking.
   - Node has to be revisited in the same path and then we can consider it to be a cycle.
   - we loop over all the elements (to cater to connected components), and then if a node is not visited then call the dfs (if dfs returns true then it means that it has detected a cycle)
    - in dfs it keeps on updating visited and path visited array as an when it finds the unvisted nodes.
    - while backtracking when a node does not have a neighbour node anymode ..it keeps on remove the node from path visited. If it has more nodes to cover then it does not remove from path visited array.
    - IMP : to have a cycle a node has to be detected as visited and path visited both.

   - Note : we maintain 2 arrays here .. 1. Visited 2. Path visited array.


------------------------- Topo Sort -------------------------
How to do topo sort:

14:
 - Topo sort. (using DFS)
  - Topo sort is only applicable for the DAGs (directed graphs without any cycle.)
  - Keep on navigating a path deeply and keep on marking them as visited until it does not have anything to explore further thats when we put it in stack .. just to preserve the order in which we were navigating the path. 
  - Popping from the stack will give us the topological sort.


15. Topo sort (using slightly modified version of BFS - using the indegree array) - Kahn's Algorithm
 - First calculate the indegree of all the nodes using the adjacency list.
 - so the graph will start with nodes whose .. indegree is 0, remember there can be many nodes whose indegree is 0 .. we need to put all the nodes in the queue .. basically this will act as the starting point node. Put it in the queue. 
 - Now we will take out the element .. and then go to adj list and remove the edges temp to the outgoing nodes in that way we reduce the indegree of the nodes.
- while doing above if the indegree of a node reduces to 0 then we push it in the queue .. which means that this becomes the candidate for the next position in the Topo sort.
- this we keep doing until we visit all the elements in the adjacency list.


Topo. Sort usages:

16. Detect Cycle in a Directed Graph using the topo sort (using BFS - sligtly modified kahn's Algo version.)
 - Remember Topo is only possible in a directed acyclic Graph
 - If the size of the topo sort is less than N then it indicates it has a cycle.


17:
 - Course Schedule I and II:
  - in this problem we are given the course schedule where course 'a' needs to performed before course 'b'. and similar dependent conditions .. and we in 1st problem we needed to check if we can finish all the courses. - we use khan's algo to find one of the topo sort.
  In 2nd problem, we needed to find one of the topo sort with slight modification in th problem set.

18: (Rivisit .. it)
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



19. Alien's Dictionary
 - something before something indicates presence/or it being a candidate for the Topo sort.
 - express all chars in terms of the number and then based in analysis ... we need to create a directed graph of chars or numbers.
 - and to determine the order of the char we need to perform the topological sorting.
Core concept compare each pair and see .. which char appear before another pair.
 - if one char is unequal (among the comparison of the entire string) then we don't need to go further that becomes the differentiating factor.
[EXPLAIN CLEARLY]


Shortest Path Algorithms:

1. Shortest path in undirected Graph with unit weights from a given start node.
 - BFS with storing node and dist and then keep a distance array as well with all having int_max or infintity as default val and then
  replacing those vals with the min distance each time while navigating through the BFS until the queue is empty.

 - Here DFS is not valid as it goes deep and can visit every node without giving the shortest path algorithm and in BFS it goes level by
  level distance by distance so node at d=1 is visited first as compared to node at d=2 and so on.


2. Shortest path in a Directed Acyclic Graph from source (one of the nodes with indegree 0) to all the nodes.
(DAG - you dont get stuck in a cycle rather .. you have a source node and a defined terminal node as well) with weights (any) using the Topological sort (DFS Topo).
 [Important : from understanding perspective]
  - step1 : Do a topo sort on the Graph using DFS and push elements to stack as the element coming
   last will be push first and the origin will be pushed on top of the stack.
  - step2 : You take out the nodes from the stack one by one and then update the distance array -
   here we keep track of the old distance and only update if the obtained distance via the current node is lesser. (This is while navigating 
   through its neigbour.)
  pre-req: Have a distance array with each node indicated by index having a distance of 'd' (intially we keep it as the int max). 
  The source node can be set as 0 as the distance from the starting origin node to the end node as itself is 0.
  Here it gives you the shortest distance for each node.

  Topo sort in a way gives a reachability graph - we know the source and we know in what flow
  - we would be able to reach the destination nodes.
   - in that way in step 2 when we keep covering each combination of path we can determine the
   shortest distance.

3.  Dijkstra's Algorithm :
 -  From a given source to all the destination in the graph 
 what will be the shortest distance. 
 we can implement the Dijkshtra algo using three methods:
 1. Priority Queue. (PQ)
  - Here we push the element in the PQ with {dist,node} as the seed element and then we pop and keep 
  exploring neighbours to find the better shortest distance if we find then we push in pQ to account for that path via this node.
   - Keeping the dist first in the pQ .. we always try to take the min dist for a given node in a greedy way.

 2. Set (Fastest)
  - For a given node if we keep on finding the better paths we keep that in the pQ .. set avoids to keep multiple paths 
  .. we always keep the best path in the set and no need to check on every path.

 3. Queue (this will take lots of time. - has unnecessary paths to cover)


- Single source shortest path
 Dij algo may or may not work for the negative edges .. This just always works for the edges with postive edges as it may visit a node and
  it wont allow to visit it if it is already visited.
  - Greedy Approach of selecting the edges. (Watch Abdul sir's lecture.)

- Dij says that find the node with the shortest path in a greedy way and update its total cost and 
then try to see if there are other nodes via this node where we can again relax the edges(find the shiortest path via that node.) 
means we can chose to get the shortest path.

   if d[u] + c(u,v) < d[v] {
      d[v] = d[u] + c(u,v)
   }

[Need better explanation] 



. Bellman ford algorithm (Single source shortst path problem ) - [Will revisit in depth as per the lecture order.]
 - As Dijkshtra does not work with negative weights .. so we use the bellman ford algo.
  - Follows a Dynamic programming strategy and solves the problem of negatives edges.
   - Finds out all the possible solution and picks up the best one.
   - Does not work if the graph comtains a cycle.

Algo:
 - It says that we need to relax all the edges for the n-1 times n = V (number of vertices).
 Why n-1 edges ?
  - Reason is if we have to relax the longest path containing all the edges then it would be n-1 edges. 
  (This is the atmost amount of iteration which we may need to relax all the edges)
  - Write all the edges in any order (order does not matter that much.)
- choose the origin and mark the distance as 0 and rest of the vertex as ∞ (opt + 5 keyboard short.).




4. Print shortest path Djikshtra Algo:
 - Hint : Just remember from where you are coming from.
 - In this problem we simply, have to maintain one distance array and another parent array to store the options where we find that if there is a better path exists from source to that node then update the distance plus also update the parent.6


5. Shortest Distance in a Binary Maze.
 - Here we create a 2D distance array to mark the source distance as 0 and rest cell we can keep as int_max(infinity)
 - Also, we can simply use queue with Djikshtra here as traversal is in 4 directions with a unit distance. [So, we need not use a PQ here.]
 - Can we simply use BFS algo in 4 direction - I believe yes .. as it also acheives that with simple queue? - Acc to GPT BFS is preferred and has less time complexity.


6. Cheapest flights within K stops [IMP Problem - Require a deep dive]:
 - There are 2 ways, I saw people have solved it .. need to deep dive to understand it better.
 Striver's way - in this way -  we need to maintain the pQ(priority queue)/queue with the nodeData as 
 {stops, node, distance} - which means that we are giving priority to stops over 
 distance(can be cost, weight) as we do in normal Dijkshtra algorithm.
 If we use hops/stops as the first element in the pQ then we do not need pQ specifically as we are
 moving in bfs fashion .. one hop/stop at a time and in that way maintaining a counter for the min
 distance.


 In the other Algo, by Shiran in this(https://www.youtube.com/playlist?list=PLT3bGNUOvbdIIX2JSC-57103xetAI4Cl_)
  series, she uses 2 arrays - 1. to maintain the distance 2. to maintain the hops.
 - in this she uses the priority queue with nodeData as {distance, node, stops} only but pushes
 into the pQ if we find a better distance for that node or we find lesser number of hops for that node.
 [If lesser number of hops are found - we dont update the distance as we still want to return the
 cheapest/lesser distance].

- In mazhar video, he has explained how simply we can perform the bfs ..and solve this question by visiting each neighbour path. (https://www.youtube.com/watch?v=VmUpydhNmuw). But here number of iterations will be more as we are exploring each path.


7. Network Delay time - Attempt it.

8. Minimum multiplications to reach the end.
 - Simple Djikshtra .. but just need to take care of the modulo condition and array containing elements upto (modulo_value -1) .. i.e if it is 

9. No of ways to arrive at destination.
 - Here we need to consider combination involving subpaths as well .. as u would be coming from various different directions.

10. Bellman Ford Algorithm [single source shortest path algo.]: 
  - Applicable for DG (directed Graph). 
  For an undirected graph -> convert it to a directed graph by associating edges from a given source to other destination.
  In this algo, we have to iterate over the edge connections (u,v,w) [u - source, v - dest node, w - weight/dist of the edge]

  - Here we need to relax all the edges N-1 times sequentially.
    - Why N-1 iterations ? : this is because in a graph with largest path in a graph with N vertices is N-1 edges and 
    as edges can be given in any order, and if the order is given such that in a iteration we can just calculate 1 value of shortest path as for
    just one value it will be given.
     
  - Check for each edge ...
  - Does not work with Graph having a negative cycle.
  - To detect a cycle .. we can perform Nth iteration if the graph gets readjusted again then it means that ... it has got a nagative cycle.


11. Floyd Warshall Algo. ( Multi source shortest path algorithm)
 - Multi source shortest path problem .. (Find the shortest distance from a given node to every other node.)
 - Concept of via a subpath .. to determine the shortest path.
 - Helps detect negative cycle as well .
 - Intuition : Check for each node for the presence of better path via another node ..as there can be path which is is better than
  the direct path via other nodes. This algo has n*n*n complexity.
 - Makes use of the concept from the Dynamic programming where we leverage the precomputed values.
 min(d[i][k] +d[k][j]) for calc the distance for d[i][j]

12. Find the city with smallest no of neighbours at a threshold distance. - (Application of Floyd warshall.) 
 - TODO - to practice.

13. Minimum Spanning Tree
  - comes the prim's algo underneath it.
  - A tree within the graph which connects each and every node and has the minimum weight to get connected.
  Intuition - It is a subgraph basically a tree without cycle .. to connect every node with min possible cost.

14. Prim's Algo
 - Pre-req : we take a min-Heap and a visited array and MST result array. {weight, node, parent} in the pQ starting with the first element pushed and market visited.
 - then we start exploring the neighbours if they are not visited and then push it to the PQ. While again taking the min out put it in the MST array and add it in the MST sum. Repeat the process.

15. Disjoint Set Data Structure. (DSU)
 - This is a custom data structure (backed by 2 array - one to store parent for a given node and other to store the rank for a given node. - Acc. to abdul sir.. can be implement using linked list as well.)
  which exposes 2 methods - 1. To do union of two disjoint set 2. Find ultimate common parent.
 - Here we dynamically form the data structure when edges are given in any order we start forming components using the union operation.
   - union has 2 cases either their parent is same which means they are already connected in the component or they are not connected .. in that case parent[u] and parent[v] we need to check which one has got higher rank or weight.. one with higher weight/rank will become the parent of the other.


16. Krushkal's Algo to find the MST using the disjoint data structure. 
 - Here goal is to find the MST .. given the number of edges.
 - Here we first sort the edges by weight .. {weight, u , v} where u is the start node of edge and v is the end node of the edge.
 - and then we start forming the MST by taking the least weight edge and ignore the edge with u,v belonging the MST (or the same component.)4
 - in this way we create the MST.


17. No of provinces using disjoint set data structure.
 - This can be solved using D set data stricture ... we need to make union by edges as they are there and then we need to count - how many nodes come with same parent [considering every node has a parent to itself] which indicates total number of parents or the bosses.

18. No of operations to make the network connected.
 - Prob : to move the extra edges from a component to use them as a supporting edge to connect other components.
 - In this problem .. we need to find the extra edges from a given component.
 - Min edges required to connect n components is = n-1.
 - if extra edges >= nC-1 this implies that ⟹ connecting components is possible.


19. Account Merge :
- in this problem .. we are given .. accounts set (containing mail ids) .. where ... in some set .. duplicate mailID exist ..
 and we need to .. merge the accout such that .. account set with duplicate or repeated mailIDs get merged to each other.
 - First we build a reverse map of the account id to the set they belong to ...we number the set from 0.. meaning we index them.
 - Basically we form the emailId -> Index mapping and then .. while doing this if we find a mailID belonging to different set then we link this set(index) with the other set as parent.
 In the 2nd step, we can go through each .. component/set and then see where does the element belong to ..as a ultimate parent .. if the 2 guys .. have the same element parent .. then we merge the set.

20. Number of Islands - II - Online queries
 - In this problem, we need to return the total distinct component at each step and form the answer to that as a list.
 - Here, approach involves first .. increasing the component count by considering the new guy as a isolated component and when we check if it is surrounder by prev visited/or counted component then we reduce the component count.
 - Also, we need to learn to identify each cell in a matrix as a unique number to consider them as node .. by that unique number we can get the coordinate cell_no = (row_num) * row_size + col_num 


21.  Making a Large Island - DSU.
 - Here we are given a matrix and we can convert 0 -> 1 to form the biggest island.. Acc to understanding .. 
 - Step1 : we form the disjoint connected components and number them as a single number fro the coordinate. Disjoint set will form the connected component and will share the unique ultimate single parent. Also, disjoint set can be represented as the flattened version or path compressed version where all the nodes are directly connected to the ultimate parent.
 - Step2 : Try converting every 0 -> 1 one by one and see if they are able to form the larger connected component/island.
  - Check left -> if there is cell(numbered which has 1) and find its ultimate parent and gets its size (as we have size and rank present in a DSU) - add totals from left, right top and botton (Edge case - remember to not count the same connected component twice or more .. this can be done by maintaining unique parents in a set data structure)

22. Most stones removed with same row and column :
 - Crux -> Identify the connected components as for 1 connected components all the elements/nodes can be removed except 1.
  - so, max(stones which can be removed) = number of nodes - (no of total connected components)

- Imp : Entire row is treated as Node. (All the elements in that row is treated as part of the same node). 
 - Also a column is also treated as a node and mapped with entire order of node. In his way, we go through each element in the matrix and see which row and column it is part of and then similarly we establish connections between nodes to form the disjoint set which ultimately gives us the unique ultimate parents.


 23. Strongly Connected Components - Kosaraju's Algorithm
  - Applicable for the directed Graph Algorithm
  - In this algorithm, we need to find number of stronly connected component or print the stronly connected component.
  - By definition, this component can be defined as the component where each node can reach every other node even if we reverse the edges in the graph or we take the normal edges.
 - We acheive this in 3 steps.
   - Perform the DFS in the graph and store the last finished nodes in a stack to keep track of the nodes which finished first vs last.
   - Reverse the graph.
   - Go through the stack to take last finished elements first .. basically these indicate the start of the graph even after Graph reversal we start from the correct element in the Graph.
   - DFS will let us detect the strongly connected components.


24. Finding Bridges in the Graph - Tarjan's Algo - using insert time and (min adj time except parent).
 - Here we need to detect - how many bridges exist in a graph.
   - For this we perform DFS and keep noting down their insertion time/step along with min adj time.
    - This min adj helps to know that if parent is reachable from the subtree even if the edge to the parent(or caller in the DFS path) is removed to see the bridge applicability. [Basically, 2 concepts named - insertion time and lowest time of insertion time.]

     - If reachable then it is not a bridge else it is a bridge.

25. Articulation Point.
 - Basically an algo to detect the single point of failure in a Graph. We perform a DFS and keep time of visit of a node in the graph and check the lowest adj number except parent and visited nodes.
 - Unlike with Bridge .. here the vertex itself is removed so we need to see if the parent of node (lets say node is x) to be removed is reachable by the children node. {parent[node] is reachable by the child[node] in that case it is node a articulation point else it is.}
 Acc to Abdul sir Vid (https://www.youtube.com/watch?v=jFZsDDB0-vo) - if u vertext is parent and v is child .. d[u] >= u[v]
 if L[v] >= d[u]{L stands for lowest and d stands for discovery number} this indicates that child can not discover nodes parent if given u node is removed, then u will be an articulation point. Above is valid for all the nodes except root.
  - Root works differently as it can have either 1 child or more child on the basis of it will become articulation point.`