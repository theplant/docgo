import './index.css'
import toc from './toc'

(window.__goplaidVueComponentRegisters =
    window.__goplaidVueComponentRegisters || []).push((Vue, vueOptions) => {
    Vue.component("toc", toc)
});

