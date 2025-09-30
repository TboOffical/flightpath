import { writable } from "svelte/store";

export let currentRoute = writable("/home")

export let currentServerName = writable("")
export let currentServerAddr = writable("")
export let currentServerApiKey = writable("")
export let docs = writable([]);