<template>
    <div>
        <vue-markdown>asd**test**
            - defineAsyncComponent
            - defineComponent
        </vue-markdown>

    </div>
</template>

<script>
import  VueMarkdown  from 'vue-markdown';

export default {
    name: 'AddQuestion',
    components: {
        VueMarkdown,
    },
}
</script>

<style scoped>

</style>