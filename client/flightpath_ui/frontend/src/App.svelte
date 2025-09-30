<script>
  import Editor from "./pages/Editor.svelte";
  import Home from "./pages/Home.svelte";
  import { currentRoute } from "./stores";
</script>


<div>
  {#if $currentRoute == "/home"}
    <Home></Home>
  {/if}
  {#if $currentRoute == "/editor"}
    <Editor></Editor>
  {/if}
</div>

