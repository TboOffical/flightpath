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

/*
[
    {
        "id": "in_pushbullet_9024",
        "type": "fpnode",
        "position": {
            "x": 0,
            "y": 0
        },
        "data": {
            "node": {
                "id": "in_pushbullet_9024",
                "type": "inlet",
                "module": "pushbullet",
                "x": 0,
                "y": 0,
                "config": {
                    "api_key": "o.oXJJe7fDpJ8Bnv1VJDz1D69O4tEGKIRN"
                }
            }
        },
        "measured": {
            "width": 140,
            "height": 40
        }
    },
    {
        "id": "mod_prefix_0022",
        "type": "fpnode",
        "position": {
            "x": 0,
            "y": 100
        },
        "data": {
            "node": {
                "id": "mod_prefix_0022",
                "type": "modifier",
                "module": "text",
                "x": 0,
                "y": 100,
                "task": "prefix",
                "listen_from": [
                    "in_pushbullet_9024"
                ],
                "config": {
                    "prefix": "New Message:"
                }
            }
        },
        "measured": {
            "width": 87,
            "height": 40
        }
    },
    {
        "id": "outlet_email",
        "type": "fpnode",
        "position": {
            "x": 0,
            "y": 200
        },
        "data": {
            "node": {
                "id": "outlet_email",
                "type": "outlet",
                "module": "email",
                "x": 0,
                "y": 200,
                "listen_from": [
                    "mod_prefix_0022"
                ],
                "config": {
                    "to_address": "wireman272@gmail.com",
                    "subject": "Test Email",
                    "server_address": "smtp.zoho.com",
                    "server_port": 465,
                    "server_username": "info@chfservices.com ",
                    "server_password": "xcg8b%qU"
                }
            }
        },
        "measured": {
            "width": 96,
            "height": 40
        }
    }
]
*/

// [
//     {
//         "id": "e-34826233",
//         "source": "in_pushbullet_9024",
//         "target": "mod_prefix_0022"
//     },
//     {
//         "id": "e-921115324",
//         "source": "mod_prefix_0022",
//         "target": "outlet_email"
//     }
// ]

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