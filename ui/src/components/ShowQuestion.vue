<template>
    <div>
        <div  v-for="question in questions" :key="question.id">
            <el-row type="flex" justify="center" class="detail">
              <el-col :span="10" class="title">
                <div @click="questionDetail(question.id)" style="cursor:pointer">
                  {{question.title}}
                </div>
              </el-col>
              <el-col :span="1" v-for="tag in question.tags" :key="tag.id" :pull="3">
                <el-tag>{{tag.name}}</el-tag>
              </el-col>
            </el-row>
            <el-row type="flex" justify="center">
              <div class="content"  @click="questionDetail(question.id)" style="cursor:pointer">
                {{question.origin_content}}
              </div>
            </el-row>
            <el-row type="flex" justify="center">
              <el-col :span="3" class="time" :push="1">{{question.created_at}}</el-col>
              <el-col :span="3" class="time">{{question.updated_at}}</el-col>
              <el-col :span="4" class="star" :push="3">
                <el-button type="primary" size="mini">{{question.is_stared}}</el-button>
              </el-col>
            </el-row>
            <el-divider></el-divider>
          </div>
          <el-row type="flex" justify="center">
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[2, 5, 10, 20]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total">
            </el-pagination>
          </el-row>
    </div>
</template>

<script>
export default {
    name: 'ShowQuestion',
    data: function () {
        return {
            pageSize:2,
            currentPage: 1,
            total: 0,
            questions: [],
        }
    },
    created() {
    this.loadAll()
    },
    methods:{
    loadAll(){
      this.request.get('/question/all',{
        params:{
          current_num:this.currentPage,
          page_size:this.pageSize,
        },
      }).then(res => {
        this.questions = res.data.msg.question_list
        this.total = res.data.msg.pagination.total
      })
    },
    handleCurrentChange(currentPage){
      this.currentPage = currentPage
      this.loadAll()
    },
    handleSizeChange(pageSize){
      this.pageSize = pageSize
      this.loadAll()
    },
      questionDetail(id){
        this.$router.push({path: '/question/' + id})
      }
    }
}
</script>

<style scoped>
.time{
    font-size: 8px;
    line-height: 30px;
    color:rgba(0, 0, 0, 0.27);
  }
  .content{
    font-size: 14px;
    color: rgba(145, 182, 182, 0.767);
    width: 50%;
  }
  .el-divider {
    width: 50%;
    margin: 30px auto;
  }
  .title{
    font-size: 20px;
    font-weight: bold;
    margin-top: 2px;
    margin-bottom: 10px;
  }
  .question,.tag,.star{
    font-size: 15px;
    color: rgba(255, 255, 255, 0.448);
  }
  
  
  .detail{
    text-align: center;
  }
</style>