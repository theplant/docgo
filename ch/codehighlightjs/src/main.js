import Vue from 'vue'

import hljs from 'highlight.js/lib/core';
import 'highlight.js/styles/stackoverflow-light.css'
import go from 'highlight.js/lib/languages/go';
import javascript from 'highlight.js/lib/languages/javascript';
import hljsVuePlugin from "@highlightjs/vue-plugin";

hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('go', go);

Vue.use(hljsVuePlugin);

(window.__goplaidVueComponentRegisters =
	window.__goplaidVueComponentRegisters || []).push(function (Vue) {
	Vue.component("highlightjs", hljsVuePlugin.component);
});
