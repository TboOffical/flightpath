<script>
    import {
        currentRoute,
        currentServerAddr,
        currentServerApiKey,
        currentServerName,
        docs,
    } from "../stores";
    import { fade, fly, scale, slide } from "svelte/transition";
    import { onMount } from "svelte";
    import {
        GetDocsFromServer,
        GetInfoFromServer,
        GetPathsFromServer,
        SavePath,
    } from "../../wailsjs/go/main/App";

    import "@fluentui/web-components/button.js";
    import "@fluentui/web-components/spinner.js";
    import "@fluentui/web-components/tablist.js";
    import "@fluentui/web-components/tab.js";
    import "@fluentui/web-components/divider.js";
    import "@fluentui/web-components/dialog.js";
    import "@fluentui/web-components/dialog-body.js";
    import "@fluentui/web-components/text-input.js";

    let loading = $state(true);
    let flowsOpen = $state(false);

    let paths = [];
    let pathsObjs = [];
    let srvInfo = {};

    let loadingMessage = $state("Init");

    let activePath = $state(0);

    onMount(async () => {
        loadingMessage = "Getting Paths";
        paths = await GetPathsFromServer(
            $currentServerAddr,
            $currentServerApiKey,
        );
        loadingMessage = "Getting Metadata";
        srvInfo = await GetInfoFromServer(
            $currentServerAddr,
            $currentServerApiKey,
        );

        let d = await GetDocsFromServer(
            $currentServerAddr,
            $currentServerApiKey,
        );

        docs.set(d);

        console.log($docs);

        loadingMessage = "Processing";

        for (let p of paths) {
            pathsObjs.push(JSON.parse(p.Data));
        }

        loadPaths();

        loading = false;
    });

    import {
        Background,
        Controls,
        MiniMap,
        Panel,
        SvelteFlow,
    } from "@xyflow/svelte";

    import "@xyflow/svelte/dist/style.css";
    import { nodesToPath, pathToNodes, randId } from "../pathToNodes";
    import FPNode from "../FPNode.svelte";

    const nodeTypes = { fpnode: FPNode };

    let nodes = $state.raw([
        // { id: "1", position: { x: 0, y: 0 }, data: { label: "1" } },
        // { id: "2", position: { x: 0, y: 100 }, data: { label: "2" } },
    ]);

    let escape = function (str) {
        return str
            .replace(/[\\]/g, "\\\\")
            .replace(/[\"]/g, '\\"')
            .replace(/[\/]/g, "\\/")
            .replace(/[\b]/g, "\\b")
            .replace(/[\f]/g, "\\f")
            .replace(/[\n]/g, "\\n")
            .replace(/[\r]/g, "\\r")
            .replace(/[\t]/g, "\\t");
    };

    let edges = $state.raw([]);

    function uppercase(str) {
        return str[0].toUpperCase() + str.slice(1);
    }

    function loadPaths() {
        let data = pathsObjs[activePath];
        let nds = pathToNodes(data);
        nodes = nds.nodes;
        edges = nds.edges;
    }

    async function save() {
        loading = true;
        loadingMessage = "Compiling";
        let path = nodesToPath(nodes, edges, pathsObjs[activePath].title);
        pathsObjs[activePath] = path;
        try {
            loadingMessage = "Uploading";
            await SavePath(
                $currentServerAddr,
                $currentServerApiKey,
                paths[activePath].ID,
                JSON.stringify(path),
            );
        } catch (e) {
            //todo: add error popup
            console.log(e);
            loading = false;
            return;
        }
        loading = false;
    }

    function addNode(d) {
        let id = "n-" + randId();

        let n = {
            id: id,
            type: d.Type == 0 ? "inlet" : "outlet",
            x: 0,
            y: 0,
            module: d.Module,
            config: {},
        };

        for (let item of d.Params) {
            n.config[item.ParamJsonField] = "";
        }

        pathsObjs[activePath].nodes.push(n);

        loadPaths();
    }

    function addNodeTask(d, t) {
        let id = "n-" + randId();

        let n = {
            id: id,
            type: "modifier",
            x: 0,
            y: 0,
            module: d.Module,
            task: t.Name,
            config: {},
        };

        for (let item of t.Params) {
            n.config[item.ParamJsonField] = "";
        }

        pathsObjs[activePath].nodes.push(n);

        loadPaths();
    }

    let nodeInserter = $state(false);
</script>

{#if nodeInserter}
    <div
        transition:fly={{ y: 100 }}
        class="fixed overflow-y-hidden flex overflow-x-scroll bottom-0 left-0 w-full h-[200px] backdrop-blur-lg z-40"
        style="background-color: rgba(0, 0, 0, .7);"
    >
        {#each $docs as d}
            <div>
                {#if d.Type == 0 || d.Type == 1}
                    <div
                        class="h-11/12 flex flex-col justify-center items-center w-[200px] shadow-md rounded-md m-2"
                        style="background-color: rgba(0, 0, 0, .6);"
                    >
                        <p class="text-2xl mb-2">{uppercase(d.Module)}</p>
                        <fluent-button
                            role="none"
                            on:click={() => {
                                addNode(d);
                            }}>Add</fluent-button
                        >
                    </div>
                {:else}
                    <!-- {JSON.stringify(d)} -->
                    <div
                        class="h-11/12 flex flex-col justify-center items-center w-[200px] shadow-md rounded-md m-2"
                        style="background-color: rgba(0, 0, 0, .6);"
                    >
                        <p class="text-2xl mb-2">{uppercase(d.Module)}</p>
                        {#each d.Tasks as t}
                            <fluent-button
                                role="none"
                                on:click={() => {
                                    addNodeTask(d, t);
                                }}>{t.Name}</fluent-button
                            >
                        {/each}
                    </div>
                {/if}
            </div>
        {/each}
    </div>
{/if}

<div class="m-2 fixed top-0 z-50" style="width: 98%;">
    <div
        class="flex justify-between p-2 rounded-md items-center w-full"
        style="background-color: rgba(0,0,0,.8); backdrop-filter: blur(10px);"
    >
        <div class="flex items-center w-1/3 s">
            <fluent-button>
                {#if loading}
                    <fluent-spinner size="tiny"></fluent-spinner>
                    <p class="ml-2">{loadingMessage}</p>
                {:else}
                    <div class="flex items-center" in:fly={{ x: 10 }}>
                        <img
                            alt="publish"
                            class="size-4 mr-2"
                            src="/src/assets/check.svg"
                        />
                        {pathsObjs.length} Path {pathsObjs.length > 1
                            ? "s"
                            : ""}
                    </div>
                {/if}
            </fluent-button>
            <p class="ml-2">
                {$currentServerName}
            </p>
        </div>
        <div class="w-1/3 flex justify-center space-x-2">
            <fluent-button icon-only>
                <img
                    in:scale
                    alt="add"
                    class="size-5 opacity-80"
                    src="/src/assets/add.png"
                />
            </fluent-button>
            <fluent-button
                role="none"
                icon-only
                on:click={() => {
                    flowsOpen = !flowsOpen;
                }}
                appearance={flowsOpen ? "primary" : ""}
            >
                <img
                    in:scale
                    alt="flows"
                    class="size-5 opacity-80"
                    src="/src/assets/flows.png"
                />
            </fluent-button>
            <div class="w-0.5 rounded-full bg-white opacity-30"></div>
            <fluent-button
                appearance={nodeInserter ? "primary" : ""}
                role="none"
                on:click={() => {
                    nodeInserter = !nodeInserter;
                }}
                icon-only
            >
                <img
                    in:scale
                    alt="nodes"
                    class="size-5 opacity-80"
                    src="/src/assets/nodes.png"
                />
            </fluent-button>
            <fluent-button icon-only role="none" on:click={save}>
                <img
                    in:scale
                    alt="publish"
                    class="size-5 opacity-80"
                    src="/src/assets/check.svg"
                />
            </fluent-button>
            <fluent-button icon-only>
                <img
                    in:scale
                    alt="test"
                    class="size-5 opacity-80"
                    src="/src/assets/play.png"
                />
            </fluent-button>
        </div>
        <div class="w-1/3 flex justify-end">
            <fluent-button
                role="none"
                icon-only
                on:click={() => {
                    currentRoute.set("/home");
                }}
            >
                <img
                    class="invert size-5"
                    alt="home"
                    src="/src/assets/images/icons8-home.svg"
                />
            </fluent-button>
        </div>
    </div>
</div>

<div class="w-full h-dvh overflow-hidden flex-col" in:fly={{ x: -10 }}>
    <div class="flex h-full w-full">
        {#if flowsOpen && !loading}
            <div class="w-4/12 m-0" transition:slide={{ axis: "x" }}>
                <div
                    class="w-11/12 ml-2 p-2 h-10/12 rounded-lg mt-[70px]"
                    style="background-color: rgba(0,0,0,.4);"
                >
                    <p class="m-1 text-xl mb-3">Paths</p>
                    <div
                        style="display: flex; flex-direction: column; gap: 1rem;"
                    >
                        <fluent-tablist orientation="vertical" size="medium">
                            {#each pathsObjs as p}
                                <fluent-tab style="height: 45px;">
                                    <div class="ml-3">
                                        {p.title}
                                    </div>
                                </fluent-tab>
                            {/each}
                        </fluent-tablist>
                    </div>
                </div>
            </div>
        {/if}
        <div class="w-full h-full">
            {#if flowsOpen}
                <div transition:slide={{ axis: "y" }} class="h-16"></div>
            {/if}
            <SvelteFlow
                class="transition-all"
                style={flowsOpen
                    ? `border-radius: 15px 0px 0px 0px;`
                    : "border-radius: 15px 15px 0px 0px;"}
                bind:nodes
                bind:edges
                {nodeTypes}
            >
                <!-- <Panel position="center-left">
                    <h1 style="color: black;">My Flow</h1>
                </Panel> -->
                <MiniMap />
                <Controls />
                <Background />
            </SvelteFlow>
        </div>
    </div>
</div>
