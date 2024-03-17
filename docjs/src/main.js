import './index.css'
import toc from './toc.vue'
import search from './search.vue'
import searchResult from './searchResult.vue'

window.__goplaidVueComponentRegisters = window.__goplaidVueComponentRegisters || []
window.__goplaidVueComponentRegisters.push((app, vueOptions) => {
  app.component('toc', toc)
  app.component('search', search)
  app.component('search-result', searchResult)

  window.storeMenuState = function (activeMenuURL) {
    const state = {
      activeMenuURL: activeMenuURL,
      scrollTop: document.getElementById('menuScroller').scrollTop
    }
    document.cookie = 'menuState=' + JSON.stringify(state) + '; path=/'
  }

  function getCookie(name) {
    const nameEQ = name + '='
    const ca = document.cookie.split(';')
    for (let i = 0; i < ca.length; i++) {
      let c = ca[i]
      while (c.charAt(0) === ' ') c = c.substring(1, c.length)
      if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length, c.length)
    }
    return null
  }

  // restore menu state
  document.addEventListener('DOMContentLoaded', function () {
    const path = window.location.pathname.replace(/^\//, '')
    const stateCookieVal = getCookie('menuState')
    if (!stateCookieVal) {
      document.getElementById(path).scrollIntoView({ block: 'center' })
      return
    }
    var state = JSON.parse(stateCookieVal)
    if (path !== state.activeMenuURL) {
      document.getElementById(path).scrollIntoView({ block: 'center' })
      return
    }
    document.getElementById('menuScroller').scrollTop = state.scrollTop
  })

  // search highlight
  document.addEventListener('DOMContentLoaded', function () {
    const h = new URLSearchParams(window.location.search).get('h')
    if (h) {
      const re = new RegExp(`(>[^<]*)(${h})([^>]*<)`, 'gi')
      document.getElementById('docMainBox').innerHTML = document
        .getElementById('docMainBox')
        .innerHTML.replaceAll(re, "$1<span class='search-h bg-yellow-300'>$2</span>$3")
      setTimeout(() => {
        document.querySelector('.search-h').scrollIntoView()
      }, 100)
    }
  })
})
