export function randId() {
    return Math.floor(Math.random() * 999999999);
}

export function pathToNodes(flow) {
    let nodes = []
    let edges = []

    for (let node of flow.nodes) {
        nodes.push({
            id: node.id, type: "fpnode", position: { x: node.x, y: node.y }, data: { node: node }
        })
        if (node.listen_from != undefined) {
            for (let id of node.listen_from) {
                edges.push({
                    id: ("e-" + randId()), source: id, target: node.id
                })
            }
        }
    }

    return { nodes: nodes, edges: edges }

}   

export function nodesToPath(nodes, edges, title){
    console.log(nodes, edges)

    let nodesJ = []

    let path = {
        title: title,
        nodes: nodesJ
    }

    for (let node of nodes){
        let nodeData = node.data.node

        nodeData.x = node.position.x
        nodeData.y = node.position.y

        if(nodeData.type == "inlet"){
            nodesJ.push(nodeData)
        }else{
            //different here because we need to calculate connections
            //need to find edges where the target is this node
            let listenFrom = []
            for(let e of edges){
                if(e.target == nodeData.id){
                    listenFrom.push(e.source)
                }
            }

            nodeData.listen_from = listenFrom
            
            nodesJ.push(nodeData)

        }
    }


    return path
}