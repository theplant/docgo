import './index.css'
import toc from './toc'
import search from './search'
import searchResult from './searchResult'

(window.__goplaidVueComponentRegisters =
    window.__goplaidVueComponentRegisters || []).push((Vue, vueOptions) => {
    window.searchBus = new Vue()
    window.getURLWithBaseURI = function(path) {
        return document.baseURI.replace(/\/+$/, "") + ("/"+path).replaceAll(/\/{2,}/g, "/")
    }

    Vue.component("toc", toc)
    Vue.component("search", search)
    Vue.component("search-result", searchResult)

    window.storeMenuState = function (activeMenuURL) {
        var state = {
            activeMenuURL: activeMenuURL,
            scrollTop: document.getElementById("menuScroller").scrollTop,
        };
        document.cookie = "menuState=" + JSON.stringify(state) + "; path=/";
    };

    function getCookie(name) {
        var nameEQ = name + "=";
        var ca = document.cookie.split(";");
        for (var i = 0; i < ca.length; i++) {
            var c = ca[i];
            while (c.charAt(0) == " ") c = c.substring(1, c.length);
            if (c.indexOf(nameEQ) == 0)
                return c.substring(nameEQ.length, c.length);
        }
        return null;
    }

    document.addEventListener("DOMContentLoaded", function() {
        // restore menu state
        {
            var path = window.location.pathname.replace(/^\//, '');
            var stateCookieVal = getCookie("menuState");
            if (!stateCookieVal) {
                document.getElementById(path).scrollIntoView({block: "center"});
                return;
            }
            var state = JSON.parse(stateCookieVal);
            if (path != state.activeMenuURL) {
                document.getElementById(path).scrollIntoView({block: "center"});
                return;
            }
            document.getElementById("menuScroller").scrollTop = state.scrollTop;
        }

        // search
        {
            var h = new URLSearchParams(window.location.search).get("h")
            if (h) {
                const re = new RegExp(`(>[^<]*)(${h})([^>]*<)`, 'gi')
                document.getElementById("docMainBox").innerHTML = document.getElementById("docMainBox").innerHTML.replaceAll(re, "$1<span class='search-h bg-yellow-300'>$2</span>$3")
                setTimeout(() => {
                    document.querySelector(".search-h").scrollIntoView()
                }, 100)
            }
        }
    })
});

