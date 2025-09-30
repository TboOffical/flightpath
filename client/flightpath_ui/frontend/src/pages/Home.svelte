<script>
    import home from "../assets/images/icons8-home.svg";
    import background from "../assets/fp_background.png";
    import settings from "../assets/icons8-settings.svg";
    import { fade, fly } from "svelte/transition";
    import { AddServer, LoadConfig } from "../../wailsjs/go/main/App";
    import {
        currentRoute,
        currentServerAddr,
        currentServerApiKey,
        currentServerName,
    } from "../stores";
    import { onMount } from "svelte";
    import { Button } from "@fluentui/web-components";

    import "@fluentui/web-components/button.js";
    import "@fluentui/web-components/spinner.js";
    import "@fluentui/web-components/tablist.js";
    import "@fluentui/web-components/tab.js";
    import "@fluentui/web-components/divider.js";
    import "@fluentui/web-components/dialog.js";
    import "@fluentui/web-components/dialog-body.js";
    import "@fluentui/web-components/text-input.js";

    let loading = true;
    let appData = {};

    async function mountFunction() {
        let d = await LoadConfig();
        appData = d;
        loading = false;
    }

    onMount(mountFunction);

    let currentPath = "/home";

    let homeItems = [
        {
            path: "/home",
            name: "home",
            icon: home,
        },
        {
            path: "/settings",
            name: "settings",
            icon: settings,
        },
    ];

    let addServer = false;
    let serverName;
    let serverAddress;
    let addServerError = "";
    let serverApiKey;

    let dialog;

    async function addServerAction() {
        if (serverName.length == 0 || serverAddress.length == 0) {
            addServerError = "Please fill out both text boxes";
            return;
        }
        await AddServer(
            serverName.value,
            serverAddress.value,
            serverApiKey.value,
        );
        addServer = false;
        dialog.hide()
        await mountFunction();
    }

    async function openServer(name, addr, apikey) {
        currentServerName.set(name);
        currentServerAddr.set(addr);
        currentServerApiKey.set(apikey);
        currentRoute.set("/editor");
    }
</script>

{#if !loading}
    <div class="home" in:fade>
        <fluent-dialog bind:this={dialog} aria-label="Add Server">
            <fluent-dialog-body
                style="margin: 20px; height: 390px; overflow-y: hidden;"
            >
                <h2 slot="title" class="text-xl">Add Server</h2>

                <div>
                    <br>
                    <fluent-text-input
                        bind:this={serverName}
                        placeholder="Test Server">Server Name</fluent-text-input
                    >
                    <br />
                    <fluent-text-input
                        bind:this={serverAddress}
                        placeholder="http://localhost:3000"
                        >Server Address</fluent-text-input
                    >
                    <br />
                    <fluent-text-input
                        bind:this={serverApiKey}
                        placeholder="83fj9f929jf9.39gmsm#"
                        >API Key</fluent-text-input
                    >
                </div>

                <p class="text-red-400">
                    {addServerError}
                </p>

                <fluent-button
                    role="none"
                    slot="action"
                    on:click={() => {
                        dialog.hide();
                    }}>Close</fluent-button
                >
                <fluent-button
                    appearance="primary"
                    role="none"
                    on:click={addServerAction}
                    slot="action">Add</fluent-button
                >
            </fluent-dialog-body>
        </fluent-dialog>

        <div class="sidebar">
            <fluent-tablist orientation="vertical" appearance="subtle">
                {#each homeItems as item}
                    <fluent-tab
                        role="none"
                        on:click={() => {
                            currentPath = item.path;
                        }}
                        selected={currentPath == item.path ? true : false}
                        style="height: fit-content; width: 60px;"
                    >
                        <div
                            role="none"
                            style="height: 60px; width: 55px; color: white; display: flex; align-items: center; justify-content: center;"
                        >
                            <img
                                style={currentPath == item.path
                                    ? "filter: brightness(0) saturate(100%) invert(70%) sepia(14%) saturate(1635%) hue-rotate(162deg) brightness(102%) contrast(99%);"
                                    : "filter: brightness(0)  invert(70%)"}
                                height="25px"
                                width="25px"
                                src={item.icon}
                                alt="icon"
                            />
                        </div>
                    </fluent-tab>
                {/each}
            </fluent-tablist>
        </div>
        <div
            style="width: 100%; height: 100%; border-radius: 10px; overflow-y: scroll;"
            class="mica"
        >
            {#if currentPath == "/home"}
                <div in:fly={{ y: -10 }}>
                    <div
                        style=" border-radius: 10px; padding: 20px; font-size: 30px; display: flex; justify-content: space-between;"
                        class="mica"
                    >
                        Servers
                        <div style="padding-top: 10px; display: flex;">
                            <fluent-button
                                role="none"
                                on:click={() => {
                                    dialog.show();
                                }}>New Server</fluent-button
                            >
                        </div>
                    </div>
                    <fluent-divider></fluent-divider>
                    <div style="padding: 10px;" class="flex justify-center">
                        {#if appData.Servers == null}
                            <div
                                style="background-color: rgba(0,0,0,0.7);"
                                class=" p-3 rounded-md backdrop-blur-3xl"
                            >
                                No Servers currently saved. Deploy the latest
                                version of flightpath and add a new one with the
                                "New Server"
                            </div>
                        {:else}
                            <div class="flex flex-col w-full space-y-2">
                                {#each appData.Servers as server, i}
                                    {#key i}
                                        <div
                                            class="w-full p-2 rounded-md"
                                            style="background-color: rgba(0,0,0,.3);"
                                            in:fly={{ y: -10, delay: i * 50 }}
                                        >
                                            <div>
                                                <div
                                                    class="flex space-x-2 items-center"
                                                >
                                                    <img
                                                        alt="server icon"
                                                        style="width: 18px; height: 18px;"
                                                        src="/src/assets/icons8-server-48.png"
                                                    />
                                                    {server.Name}
                                                </div>
                                                <div>
                                                    <div
                                                        class="pb-5 flex space-x-2 items-center"
                                                    >
                                                        Path:
                                                        <p
                                                            class="italic font-mono ml-1"
                                                        >
                                                            {server.Addr}
                                                        </p>
                                                    </div>
                                                    <fluent-button
                                                        role="none"
                                                        variant="accent"
                                                        on:click={() => {
                                                            openServer(
                                                                server.Name,
                                                                server.Addr,
                                                                server.ApiKey,
                                                            );
                                                        }}
                                                    >
                                                        Open</fluent-button
                                                    >
                                                </div>
                                            </div>
                                        </div>
                                    {/key}
                                {/each}
                            </div>
                        {/if}
                    </div>
                </div>
            {/if}
            {#if currentPath == "/settings"}
                <div in:fly={{ y: -10 }}>
                    <div style="padding: 10px;">
                        <h1 class="text-2xl mb-2 p-2">Settings</h1>
                        <h1 class="text mb-1 p-2">About</h1>
                        <div>
                            Flightpath Windows Editor<br />
                            <p class="text-sm opacity-75">
                                ©2025 Callum F. All Rights Reserved
                            </p>

                            <fluent-button appearance="transparent"
                                >Star on Github</fluent-button
                            >
                        </div>
                    </div>
                </div>
            {/if}
        </div>
    </div>
{:else}
    <div
        style="width: 100vw; height: 100dvh; display: flex; justify-content: center; align-items: center;"
    >
        <fluent-spinner></fluent-spinner>
    </div>
{/if}
