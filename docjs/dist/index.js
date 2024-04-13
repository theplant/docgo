(function(e,u){typeof exports=="object"&&typeof module<"u"?u(require("vue")):typeof define=="function"&&define.amd?define(["vue"],u):(e=typeof globalThis<"u"?globalThis:e||self,u(e.Vue))})(this,function(e){"use strict";const u={__name:"toc",setup(s){const l=e.ref(null);return e.onMounted(()=>{const c=document.querySelectorAll("body h2"),t=document.createElement("ul");l.value.appendChild(t),c.forEach((o,r)=>{const i=document.createElement("li"),n=document.createElement("a"),p=encodeURIComponent(o.textContent)+"-"+r;o.id=p,n.textContent=o.textContent,n.href=`${location.pathname}#${p}`,i.appendChild(n),t.appendChild(i)})}),(c,t)=>(e.openBlock(),e.createElementBlock("div",{ref_key:"toc",ref:l,class:"toc"},null,512))}};function E(s){return{all:s=s||new Map,on:function(l,c){var t=s.get(l);t?t.push(c):s.set(l,[c])},off:function(l,c){var t=s.get(l);t&&(c?t.splice(t.indexOf(c)>>>0,1):s.set(l,[]))},emit:function(l,c){var t=s.get(l);t&&t.slice().map(function(o){o(c)}),(t=s.get("*"))&&t.slice().map(function(o){o(l,c)})}}}const m=E(),L={__name:"search",setup(s){const l=e.ref(""),c=e.ref(null),t=()=>{c.value||(c.value=document.getElementById("docContentBox")),l.value?c.value.classList.add("hidden"):c.value.classList.remove("hidden"),m.emit("doSearch",l.value)};return(o,r)=>e.withDirectives((e.openBlock(),e.createElementBlock("input",{"onUpdate:modelValue":r[0]||(r[0]=i=>l.value=i),onInput:t,onKeyup:e.withKeys(t,["enter"]),class:"inline-block w-full h-full px-3 py-3 focus:outline-none bg-gray-50 border-b border-gray-200",placeholder:"Type to search"},null,544)),[[e.vModelText,l.value]])}},B={class:"flex flex-row w-full"},b={class:"w-full"},O={class:"w-full text-center mt-8 mb-10 text-3xl font-extralight text-gray-600 uppercase"},S={class:"w-full px-40 mx-0 list-none"},I={class:"mb-6 mt-0"},M=["href"],R=["innerHTML"],T={__name:"searchResult",setup(s){const l=e.ref(""),c=e.ref([]),t=e.ref([]),o=n=>n.replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;").replace(/"/g,"&quot;").replace(/'/g,"&#039;"),r=function(n){return document.baseURI.replace(/\/+$/,"")+("/"+n).replaceAll(/\/{2,}/g,"/")},i=n=>{l.value=n,t.value=[];const p=new RegExp(n,"i"),f=new RegExp(`(${n})`,"ig");c.value.forEach(h=>{const y=h.Body.search(p);if(y>=0){let d=h.Body.split(" "),g=-1,w=0;for(let a=0;a<d.length;a++){if(w+d[a].length+1+20>=y){g=a;break}w+=d[a].length+1}let k=d.length-1,x=0;for(let a=g;a<d.length;a++)if(x+=d[a].length,x>=200){k=a;break}let _=d.slice(g,k+1).join(" ");_=o(_);let C=_.replaceAll(f,"<span class='bg-yellow-300'>$1</span>");t.value.push({URL:r(`/${h.URL}?h=${n}`),Title:h.Title,Body:C})}})};return e.onMounted(()=>{m.on("doSearch",i),fetch(r("/search_indexes.json")).then(n=>n.json()).then(n=>c.value=n||[])}),e.onUnmounted(()=>{m.off("doSearch",i)}),(n,p)=>e.withDirectives((e.openBlock(),e.createElementBlock("div",B,[e.createElementVNode("div",b,[e.createElementVNode("h1",O,e.toDisplayString(t.value.length)+' Results Matching "'+e.toDisplayString(l.value)+'" ',1),e.createElementVNode("ul",S,[(e.openBlock(!0),e.createElementBlock(e.Fragment,null,e.renderList(t.value,f=>(e.openBlock(),e.createElementBlock("li",I,[e.createElementVNode("h3",null,[e.createElementVNode("a",{href:f.URL,class:"text-xl text-blue-700"},e.toDisplayString(f.Title),9,M)]),e.createElementVNode("p",{class:"text-base",innerHTML:f.Body},null,8,R)]))),256))])])],512)),[[e.vShow,l.value]])}};window.__goplaidVueComponentRegisters=window.__goplaidVueComponentRegisters||[],window.__goplaidVueComponentRegisters.push((s,l)=>{s.component("toc",u),s.component("search",L),s.component("search-result",T),window.storeMenuState=function(t){const o={activeMenuURL:t,scrollTop:document.getElementById("menuScroller").scrollTop};document.cookie="menuState="+JSON.stringify(o)+"; path=/"};function c(t){const o=t+"=",r=document.cookie.split(";");for(let i=0;i<r.length;i++){let n=r[i];for(;n.charAt(0)===" ";)n=n.substring(1,n.length);if(n.indexOf(o)===0)return n.substring(o.length,n.length)}return null}document.addEventListener("DOMContentLoaded",function(){const t=window.location.pathname.replace(/^\//,""),o=c("menuState");if(!o){document.getElementById(t).scrollIntoView({block:"center"});return}var r=JSON.parse(o);if(t!==r.activeMenuURL){document.getElementById(t).scrollIntoView({block:"center"});return}document.getElementById("menuScroller").scrollTop=r.scrollTop}),document.addEventListener("DOMContentLoaded",function(){const t=new URLSearchParams(window.location.search).get("h");if(t){const o=new RegExp(`(>[^<]*)(${t})([^>]*<)`,"gi");document.getElementById("docMainBox").innerHTML=document.getElementById("docMainBox").innerHTML.replaceAll(o,"$1<span class='search-h bg-yellow-300'>$2</span>$3"),setTimeout(()=>{document.querySelector(".search-h").scrollIntoView()},100)}})})});
