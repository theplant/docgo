<template>
  <input
    v-model="input"
    @input="search"
    @keyup.enter="search"
    class="inline-block w-full h-full px-3 py-3 focus:outline-none bg-gray-50 border-b border-gray-200"
    placeholder="Type to search"
  />
</template>

<script setup>
import { eventBus } from './eventBus'
import { ref } from 'vue'
const input = ref('')
const docContentBox = ref(null)

const search = () => {
  if (!docContentBox.value) {
    docContentBox.value = document.getElementById('docContentBox')
  }
  if (input.value) {
    docContentBox.value.classList.add('hidden')
  } else {
    docContentBox.value.classList.remove('hidden')
  }
  eventBus.emit('doSearch', input.value)
}
</script>
