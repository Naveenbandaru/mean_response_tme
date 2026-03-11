import (
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

type Node struct {
    id int
    load int
}

func (n *Node) processRequest(req string) string {
    delay := rand.Intn(100) + n.load
    time.Sleep(time.Duration(delay) * time.Millisecond)
    return fmt.Sprintf("Node %d handled request: %s", n.id, req)
}

type Cluster struct {
    nodes []*Node
    mu    sync.Mutex
}

func (c *Cluster) selectNode() *Node {
    c.mu.Lock()
    defer c.mu.Unlock()
    var chosen *Node
    for _, n := range c.nodes {
        if chosen == nil || n.load < chosen.load {
            chosen = n
        }
    }
    chosen.load += 10
    return chosen
}

func (c *Cluster) handleRequest(req string) string {
    node := c.selectNode()
    result := node.processRequest(req)
    c.mu.Lock()
    node.load -= 5
    c.mu.Unlock()
    return result
}

func newCluster(size int) *Cluster {
    nodes := make([]*Node, size)
    for i := 0; i < size; i++ {
        nodes[i] = &Node{id: i + 1, load: 0}
    }
    return &Cluster{nodes: nodes}
}

func main() {
    cluster := newCluster(5)
    http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
        result := cluster.handleRequest("data")
        fmt.Fprintln(w, result)
    })
    fmt.Println("Proposed cluster running on port 8080")
    http.ListenAndServe(":8080", nil)
}
