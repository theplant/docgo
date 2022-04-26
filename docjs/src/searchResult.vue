<template>
    <div
        class="flex flex-row w-full"
        v-show="keyword"
        >
        <div class="w-full">
            <h1 class="w-full text-center mt-8 mb-10 text-3xl font-extralight text-gray-600 uppercase">{{result.length}} Results Matching "{{keyword}}"</h1>
            <ul class="w-full px-40 mx-0 list-none">
                <li
                    class="mb-6 mt-0"
                    v-for="item in result"
                    >
                    <h3><a 
                            :href="'/' + item.URL + '?h=' + keyword" 
                            class="text-xl text-blue-700"
                            >{{item.Title}}</a></h3>
                    <p class="text-base" v-html="item.Body"></p>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
export default {
    name: "search-result",
    data() {
        return {
            keyword: "",
            // [{URL, Title, Body}]
            indexes: [],
            result: [],
        }
    },
    methods: {
        escapeHTML(s) {
            return s.replace(/&/g, "&amp;")
                .replace(/</g, "&lt;")
                .replace(/>/g, "&gt;")
                .replace(/"/g, "&quot;")
                .replace(/'/g, "&#039;")
        }
    },
    mounted() {
        searchBus.$on("doSearch", k => {
            this.keyword = k
            this.result = []
            var re = new RegExp(k, "i")
            var rre = new RegExp(`(${k})`, "ig")
            this.indexes.forEach(index => {
                var sIdx = index.Body.search(re)
                if (sIdx >= 0) {
                    var slices = index.Body.split(" ")
                    var startIdx = -1
                    var trimLeftLen = 0
                    for (var i = 0; i < slices.length; i++) {
                        if ((trimLeftLen + slices[i].length + 1 + 20) >= sIdx) {
                            startIdx = i
                            break
                        }
                        trimLeftLen += slices[i].length + 1
                    }
                    var endIdx = slices.length - 1
                    var txtLen = 0
                    for (var i = startIdx; i < slices.length; i++) {
                        txtLen += slices[i].length
                        if (txtLen >= 200) {
                            endIdx = i
                            break
                        }
                    }
                    var plainBody = slices.slice(startIdx, endIdx+1).join(" ")
                    plainBody = this.escapeHTML(plainBody)
                    var colorBody = plainBody.replaceAll(rre, "<span class='bg-yellow-300'>$1</span>")
                    this.result.push({
                        URL: index.URL,
                        Title: index.Title,
                        Body: colorBody,
                    })
                }
            })
        })

        this.$nextTick(() => {
            fetch("/search_indexes.json").
                then(r => r.json()).
                then(r => this.indexes = r || [])
        })
    }
}
</script>
