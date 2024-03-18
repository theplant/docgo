import 'highlight.js/styles/stackoverflow-light.css'
import hljs from 'highlight.js/lib/core'
import go from 'highlight.js/lib/languages/go'
import javascript from 'highlight.js/lib/languages/javascript'
import hljsVuePlugin from '@highlightjs/vue-plugin'

hljs.registerLanguage('javascript', javascript)
hljs.registerLanguage('go', go)

;(window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || []).push(
  function (app) {
    app.use(hljsVuePlugin)
  }
)
