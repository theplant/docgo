<template>
  <div class="flex flex-row w-full" v-show="keyword">
    <div class="w-full">
      <h1 class="w-full text-center mt-8 mb-10 text-3xl font-extralight text-gray-600 uppercase">
        {{ result.length }} Results Matching "{{ keyword }}"
      </h1>
      <ul class="w-full px-40 mx-0 list-none">
        <li class="mb-6 mt-0" v-for="item in result">
          <h3>
            <a :href="item.URL" class="text-xl text-blue-700">{{ item.Title }}</a>
          </h3>
          <p class="text-base" v-html="item.Body"></p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { eventBus } from './eventBus'

const keyword = ref('')
const indexes = ref([])
const result = ref([])

const escapeHTML = (s) => {
  return s
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;')
    .replace(/'/g, '&#039;')
}

const getURLWithBaseURI = function (path) {
  return document.baseURI.replace(/\/+$/, '') + ('/' + path).replaceAll(/\/{2,}/g, '/')
}

const doSearch = (k) => {
  keyword.value = k
  result.value = []
  const re = new RegExp(k, 'i')
  const rre = new RegExp(`(${k})`, 'ig')
  indexes.value.forEach((index) => {
    const sIdx = index.Body.search(re)
    if (sIdx >= 0) {
      let slices = index.Body.split(' ')
      let startIdx = -1
      let trimLeftLen = 0
      for (let i = 0; i < slices.length; i++) {
        if (trimLeftLen + slices[i].length + 1 + 20 >= sIdx) {
          startIdx = i
          break
        }
        trimLeftLen += slices[i].length + 1
      }
      let endIdx = slices.length - 1
      let txtLen = 0
      for (let i = startIdx; i < slices.length; i++) {
        txtLen += slices[i].length
        if (txtLen >= 200) {
          endIdx = i
          break
        }
      }
      let plainBody = slices.slice(startIdx, endIdx + 1).join(' ')
      plainBody = escapeHTML(plainBody)
      let colorBody = plainBody.replaceAll(rre, "<span class='bg-yellow-300'>$1</span>")
      result.value.push({
        URL: getURLWithBaseURI(`/${index.URL}?h=${k}`),
        Title: index.Title,
        Body: colorBody
      })
    }
  })
}

onMounted(() => {
  eventBus.on('doSearch', doSearch)

  fetch(getURLWithBaseURI('/search_indexes.json'))
    .then((r) => r.json())
    .then((r) => (indexes.value = r || []))
})

onUnmounted(() => {
  eventBus.off('doSearch', doSearch)
})
</script>
