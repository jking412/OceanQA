<template>
    <div>
        <markdown-editor v-model="content"></markdown-editor>
    </div>
</template>

<script>
import { MarkdownEditor } from 'markdown-it-editor'
import 'markdown-it-editor/lib/index.css'

export default {
    name: 'AddQuestion',
    components: {
        MarkdownEditor,
    },
    data () {
        return {
        content: 'hello world',
        }
    },
   methods: {
    test(){
        let e = document.getElementsByClassName('markdown__editor-preview markdown-body')[0];
        console.log(e.innerHTML);
    }
   },
    
}
</script>

<style scoped>

</style>