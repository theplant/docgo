<template>
  <div ref="toc" class="toc" />
</template>
<script setup>
import { onMounted, ref } from 'vue'

const toc = ref(null)

onMounted(() => {
  const matches = document.querySelectorAll('body h2')
  const ul = document.createElement('ul')
  toc.value.appendChild(ul)

  matches.forEach((value, index) => {
    const li = document.createElement('li')
    const a = document.createElement('a')
    // Ensuring unique ID by appending an index
    const id = encodeURIComponent(value.textContent) + '-' + index
    value.id = id

    a.textContent = value.textContent
    a.href = `${location.pathname}#${id}`
    li.appendChild(a)
    ul.appendChild(li)
  })
})
</script>
