import './output.css'
import App from './App.svelte'
import { mount } from "svelte";
import './custom.css'

import { webDarkTheme, webLightTheme } from '@fluentui/tokens';
import { setTheme } from '@fluentui/web-components';

setTheme(webDarkTheme)

const app = mount(App, {
  target: document.getElementById('app')
})

export default app
