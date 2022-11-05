<template>
    <div>
        <el-row style="font-size: 40px;">提交新的问题</el-row>
        <el-row style="font-size: 30px;">
            <label>
                标题
                <span style="margin: 5px;"></span>
                <el-input v-model="title" style="width: 40%;"></el-input>
            </label>
            <div></div>
            <label>
                链接
                <span style="margin: 5px;"></span>
                <el-input v-model="answer_url" style="width: 40%;"></el-input>
            </label>
            <div style="margin: 10px;"></div>
            <el-tag
            v-for="tag in tags"
            :key="tag.id"
            closable
            @close="deleteTag(tag)">
            {{tag.name}}
            </el-tag>
            <el-popover
            placement="right"
            width="200"
            trigger="click">

            <el-autocomplete
            v-model="queryTag"
            placeholder="请输入你想要加入的标签" 
            :fetch-suggestions="querySearch"
            :trigger-on-focus="false"
            @select="handleSelect"
            >
            <template slot-scope="{item}">
                <div v-if="item.id==0" style="display: flex;">
                    <div style="margin: auto;">
                        <el-button type="text" icon="el-icon-plus"></el-button>
                    </div>
                </div>
                <div v-else style="text-align: center;">{{item.name}}</div>
              </template>
            </el-autocomplete>
            <el-button type="primary" icon="el-icon-circle-plus-outline" slot="reference">标签</el-button>
            </el-popover>
        </el-row>
        <markdown-editor v-model="originContent"></markdown-editor>
        <el-row>
            <el-button type="primary" @click="submit">提交</el-button>
        </el-row>

        <el-dialog title="新建标签" :visible.sync="dialogFormVisible">
            <el-form :model="tagForm">
              <el-form-item label="标签名称" label-width="120px">
                <el-input  v-model="tagForm.name"></el-input>
              </el-form-item>
            </el-form>
            <div slot="footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="submitTag">确 定</el-button>
            </div>
          </el-dialog>
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
            title: '',
            answer_url: '',
            originContent: 'hello world',
            tags:[],
            dialogFormVisible: false,
            queryTag: '',
            tagForm: {
                name: '',
            },
            suggestionTags: [],
        }
    },
   methods: {
    querySearch(queryString, callback) {
        this.request.get('/tag/search',{
            params: {
                query: queryString,
            }
        }).then((res) => {
            this.suggestionTags = res.data.msg
            let result = [{
                id: 0,
            }]
            callback(this.suggestionTags.concat(result))
        })
    },
    handleSelect(item){
        if(item.id == 0){
            this.dialogFormVisible = true;
        }else{
            this.tags.push(item);
        }
    },
    submitTag(){
        this.dialogFormVisible = false;
        let data = {
            name: this.tagForm.name,
        }
        this.request.post('/tag/create', data).then(res => {
            this.tags.push(res.data.msg);
        })
    },
    deleteTag(tag){
        for (let i = 0; i < this.tags.length; i++) {
            if (this.tags[i].id === tag.id) {
                this.tags.splice(i, 1);
                break;
            }
        }
    },
    submit(){
        let data = {
            title: this.title,
            answer_url: this.answer_url,
            origin_content: this.originContent,
            parse_content: document.getElementsByClassName('markdown__editor-preview markdown-body')[0].innerHTML,
            tags: this.tags,
        }
        this.request.post('/question/create', data).then(res => {
            if(res.status == 200){
                this.$router.push({path: '/'})
            }
        })
    },
   },
    
}
</script>

<style scoped>
.el-row{
    margin: 30px;
}

.el-tag{
    margin: 5px;
}
</style>