import './index.css'
import toc from './toc'
import search from './search'
import searchResult from './searchResult'

(window.__goplaidVueComponentRegisters =
    window.__goplaidVueComponentRegisters || []).push((Vue, vueOptions) => {
    window.searchBus = new Vue()
    Vue.component("toc", toc)
    Vue.component("search", search)
    Vue.component("search-result", searchResult)

    document.addEventListener("DOMContentLoaded", function() {
        var h = new URLSearchParams(window.location.search).get("h")
        if (h) {
            const re = new RegExp(`(>[^<]*)(${h})([^>]*<)`, 'gi')
            document.getElementById("docContentBox").innerHTML = document.getElementById("docContentBox").innerHTML.replaceAll(re, "$1<span class='search-h bg-yellow-300'>$2</span>$3")
            setTimeout(() => {
                document.querySelector(".search-h").scrollIntoView()
            }, 100)
        }
    })
});

