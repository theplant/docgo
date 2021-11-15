<template>
  <div ref="toc" class="toc"/>
</template>

<script>
export default {
  name: 'toc',
  mounted() {
    this.$nextTick(() => {
      const toc = this.$refs.toc;
      const matches = document.querySelectorAll(`body h2`);
      const ul = document.createElement('ul');
      toc.appendChild(ul);
      matches.forEach(value => {
        const li = document.createElement('li');
        const a = document.createElement('a');
        value.id = encodeURIComponent(value.textContent)
        a.innerHTML = value.textContent;
        a.href = `${location.pathname}#${value.id}`;
        li.appendChild(a);
        ul.appendChild(li);
      });
    });
  }
};
</script>
