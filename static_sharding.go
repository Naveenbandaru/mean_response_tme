import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

type Node struct {
    id int
}
func (n *Node) processRequest(req string) string {
    delay := rand.Intn(200)
    time.Sleep(time.Duration(delay) * time.Millisecond)
    return fmt.Sprintf("Node %d handled request: %s", n.id, req)
}
type Cluster struct {
    nodes []*Node
}
func (c *Cluster) handleRequest(req string) string {
    index := rand.Intn(len(c.nodes))
    return c.nodes[index].processRequest(req)
}
func newCluster(size int) *Cluster {
    nodes := make([]*Node, size)
    for i := 0; i < size; i++ {
        nodes[i] = &Node{id: i + 1}
    }
    return &Cluster{nodes: nodes}
}
func main() {
    cluster := newCluster(5)
    http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
        result := cluster.handleRequest("data")
        fmt.Fprintln(w, result)
    })
    server := &http.Server{
        Addr: ":8080",
    }
    fmt.Println("Cluster running on port 8080")
    server.ListenAndServe()
}
