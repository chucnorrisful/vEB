# vEB
Go implementation of the van Emde Boas tree data structure: Priority queue for positive whole numbers in O(log log u) time.

## Supports the following Priority-Queue operations:

**Insert(int)** insert a positive number, allowed range: [0, u)

**Delete(int)** deletes a previously inserted number

**Succ(int)** finds the next larger number that is already stored inside the tree. Note that the parameter of succ doesn't neet to exist inside the tree. Specially, Succ(-1) gives the minimum (smallest stored element) of the tree.

## Runtime and Space:

All operations run in **O(log log u)** time with u being a large integer provided at initialisation providing an upper limit for the allowed numbers to be inserted.
Space requirement of the (fully filled with all u elements) tree is O(u), as well as initialisation time.

Current implementation uses lazy initialisation, so the init-time is **O(sqrt(u))**, and Insert may run slower until all of the substructures of the tree were used at least once. I might add a switch to toggle both modes at some point in the future.

## Usage:

'''
github.com/chucnorrisful/vEB
'''

See test/main.go for examples.

## Todos:

- add safety features (member check on insertion, deletion etc.)
- small optimisations: use bitmasks instead of integer division and mod2 calculations
- add switch for lazyInit vs. fullInit
- add sparse-mode: usage of hashmaps instead of arrays in internal structure may yield in massive space savings if inserted elements are sparse in a large universe.
- add linked-list for nodes of tree: enables Succ, Pred in constant time
- edit PrioQ protocol to also consume Member, Pred, Min, Max
- extend to negative numbers
