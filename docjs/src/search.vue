<template>
    <input 
    v-model="input"
    @input="search"
    @keyup.enter="search"
    class="inline-block w-full h-full px-3 py-3 focus:outline-none bg-gray-50 border-b border-gray-200"
    placeholder="Type to search"
    />
</template>

<script>
export default {
    name: "search",
    data() {
        return {
            input: "",

            docContentBox: null,
        }
    },
    methods: {
        search() {
            if (!this.docContentBox) {
                this.docContentBox = document.getElementById("docContentBox")
            }
            if (this.input) {
                this.docContentBox.classList.add("hidden")
            } else {
                this.docContentBox.classList.remove("hidden")
            }
            searchBus.$emit("doSearch", this.input)
        }
    },
}
</script>
