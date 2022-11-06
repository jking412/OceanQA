<template>
    <div>
        <el-row type="flex" justify="center">
            <span style="font-size: 40px;">{{title}}</span>
            <span class="read">阅读次数 {{readTime}}</span>
        </el-row>
        <el-row type="flex" justify="center" v-if="tags.length != 0">
            <el-tag type="primary" style="margin: 5px;line-height: 30px;" v-for="tag in tags" :key="tag.id">{{tag.name}}</el-tag>
        </el-row>
        <el-row type="flex" justify="center" class="main">
            <el-col v-html="parseContent" class="markdown__editor-preview markdown-body"></el-col>
        </el-row>
        <el-row type="flex" justify="left" style="margin-left:20%">
            <a :href="answer_url" target="_blank">问题博文链接</a>
            <a @click="edit" >edit</a>
            <a @click="del">delete</a>
            <i class="el-icon-star-off" @click="star" v-if="!isStared">收藏</i>
            <i class="el-icon-star-on" @click="star" v-else>取消收藏</i>
        </el-row>
        <el-row type="flex" justify="center">
            <el-col :span="6" class="time">
                创建时间：{{createdAt}}
            </el-col>
            <el-col :span="6" class="time">
                更新时间：{{updatedAt}}
            </el-col>
        </el-row>
        <el-row type="flex" justify="center" style="margin-top: 20px;">
            <el-button type="primary" @click="goBack">返回</el-button>
        </el-row>
    </div>
</template>

<script>
import 'markdown-it-editor/lib/index.css'

export default {
    name: 'DetailQuestion',
    data(){
        return{
            id: 0,
            title: '',
            createdAt: '',
            updatedAt: '',
            answer_url: '',
            originContent: '',
            parseContent: '',
            readTime: 0,
            isStared: false,
            tags:[],
        }
    },
    created(){
        this.id = this.$route.params.id
        this.getQuestion()
    },
    methods:{
        goBack(){
            this.$router.push({path: '/'})
        },
        edit(){

        },
        del(){
        this.$confirm('此操作将永久删除该问题, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
            this.request.get('/question/delete', {
                params: {
                    id: this.id
                }
            }).then(res => {
                if(res.status == 200){
                    this.$message({
                        type: 'success',
                        message: '删除成功!'
                    });
                    this.$router.push({path: '/'})
                }else{
                    this.$message({
                        type: 'error',
                        message: '删除失败!'
                    });
                }
            })
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          });          
        });
        },
        star(){
            if(this.isStared){
                this.request.get('/question/star',{
                    params: {
                        id: this.id,
                        is_stared: false
                    }
                }).then(res => {
                    this.isStared = false
                    if(res.status == 200){
                        this.$message({
                        type: 'info',
                        message: '取消收藏'
                    });
                }
                })
            }else{
                this.request.get('/question/star',{
                    params: {
                        id: this.id,
                        is_stared: true
                    }
                }).then(res => {
                    this.isStared = true
                    if(res.status == 200){
                        this.$message({
                        type: 'success',
                        message: '收藏成功!'
                    });
                }
                })
            }
        },
        getQuestion(){
            this.request.get('/question/show',{
                params: {
                    id: this.id
                }
            }).then((res)=>{
                this.title = res.data.msg.title
                this.createdAt = res.data.msg.created_at
                this.updatedAt = res.data.msg.updated_at
                this.answer_url = res.data.msg.answer_url
                this.originContent = res.data.msg.origin_content
                this.parseContent = res.data.msg.parse_content
                this.readTime = res.data.msg.read_time
                this.isStared = res.data.msg.is_stared
                this.tags = res.data.msg.tags
            })
        },
    },
}
</script>

<style scoped>

.time{
    color: rgba(88, 114, 138, 0.555);
}

span{
    line-height: 40px;
}

.read{
    margin-left: 50px;
    font-size: 15px;
    color: rgba(94, 126, 153, 0.508);
}
.main{
    width: 60%;
    margin: auto;
    background-color: rgba(234, 234, 229, 0.522);
    padding: 70px;
    border-radius: 10px;
}

i{
    margin: 10px;
    margin-left: 30%;
    cursor: pointer;
}

a{
    text-decoration: none;
    color: rgb(162, 187, 203);
    margin: 10px;
}
.el-row{
    margin-top: 20px;
}
</style>