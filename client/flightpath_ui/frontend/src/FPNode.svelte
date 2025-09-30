<script>
    import { Handle, Position, useSvelteFlow } from "@xyflow/svelte";
    import { onMount } from "svelte";

    import "@fluentui/web-components/button.js";
    import "@fluentui/web-components/text-input.js";
    import { flip } from "svelte/animate";
    import { fade, fly, scale, slide } from "svelte/transition";
    import { docs } from "./stores";

    let { id, data } = $props();

    function uppercase(str) {
        return str[0].toUpperCase() + str.slice(1);
    }

    function getStyle() {
        if (data.node.type == "inlet") {
            return "background-color: rgba(0,50,0,.6);";
        }
        if (data.node.type == "outlet") {
            return "background-color: rgba(50,0,0,.6);";
        }
        if (data.node.type == "modifier") {
            return "background-color: rgba(0,0,0,.6);";
        }
    }

    let { updateNodeData } = useSvelteFlow();
    let editOpen = $state(false);
</script>

{#if editOpen}
    <div
        transition:slide={{ axis: "x" }}
        class="fixed flex h-fit z-40 w-[400px] backdrop-blur-sm flex-col justify-between rounded-md shadow-2xl text-white"
        style="background-color: rgba(0, 0, 0, .7);"
    >
        <div
            class="p-2 rounded-t-md flex items-center"
            style="background-color: rgba(0, 0, 0, .6);"
        >
            Editing <p class="font-mono ml-1">{data.node.id}</p>
        </div>
        <div class="h-full m-3">
            {#each $docs as m}
                {#if m.Module == data.node.module}
                    {#if data.node.task != undefined}
                        {#each m.Tasks as t}
                            {#if t.Name == data.node.task}
                                {#each t.Params as p}
                                    <p>{p.ParamName}</p>
                                    <input
                                        class="border w-full focus-visible:outline-0 font-mono p-1 rounded-md"
                                        bind:value={
                                            data.node.config[p.ParamJsonField]
                                        }
                                        placeholder={p.ParamType}
                                    />
                                    <p
                                        class="text-sm text-gray-400 italic mb-2"
                                    >
                                        {p.ParamDescription}
                                    </p>
                                {/each}
                            {/if}
                        {/each}
                    {:else}
                        {#each m.Params as p}
                            <p>{p.ParamName}</p>
                            <input
                                class="border w-full focus-visible:outline-0 font-mono p-1 rounded-md"
                                bind:value={data.node.config[p.ParamJsonField]}
                                placeholder={p.ParamType}
                            />
                            <p class="text-sm text-gray-400 italic mb-2">
                                {p.ParamDescription}
                            </p>
                        {/each}
                    {/if}
                {/if}
            {/each}
            <div class="h-20"></div>
        </div>
        <div
            class="p-2 rounded-b-md"
            style="background-color: rgba(0, 0, 0, .6);"
        >
            <fluent-button
                role="none"
                onclick={() => {
                    editOpen = !editOpen;
                }}>Done</fluent-button
            >
        </div>
    </div>
{:else}
    <div
        transition:slide
        class="fpnode text-white p-2 backdrop-blur-md shadow-md rounded-md font-mono"
        style={getStyle()}
    >
        <Handle type="target" position={Position.Top} />
        <div>
            <div>
                <div class="flex items-center">
                    {uppercase(data.node.module)}
                    <fluent-button
                        role="none"
                        onclick={() => {
                            editOpen = !editOpen;
                        }}
                        class="ml-3"
                        size="small"
                        icon-only
                        shape="circular"
                        appearance="outline"
                    >
                        <img
                            height="15"
                            width="15"
                            class="invert"
                            alt="edit"
                            src="/src/assets/edit.svg"
                        />
                    </fluent-button>
                </div>
            </div>
        </div>
        <Handle type="source" position={Position.Bottom} />
    </div>
{/if}
