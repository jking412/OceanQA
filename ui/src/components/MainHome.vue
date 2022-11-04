<template>
  <div>
    <el-container>
        <el-row type="flex" class="row-bg" justify="start">
          <el-col :span="4"><div class="grid-content answer">Answer</div></el-col>
          <el-col :span="0.5"><div class="grid-content question">问题</div></el-col>
          <el-col :span="0.5"><div class="grid-content tag">标签</div></el-col>
          <el-col :span="0.5"><div class="grid-content star">收藏</div></el-col>
          <el-col :span="4" class="input-content" :push="4"><el-input class="input-search" v-model="keyword" placeholder="请输入内容"
            prefix-icon="el-icon-search"></el-input></el-col>
          <el-col :span="1" :push="8"><el-button type="primary">添加问题</el-button></el-col>

        </el-row>
      <el-main>
        <div  v-for="question in questions" :key="question.id">
          <el-row type="flex" justify="center" class="detail">
            <el-col :span="10" class="title">{{question.title}}</el-col>
            <el-col :span="1" v-for="tag in question.tags" :key="tag.id" :pull="3">
              <el-tag>{{tag.name}}</el-tag>
            </el-col>
          </el-row>
          <el-row type="flex" justify="center">
            <div class="content">
              {{question.content}}
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

      </el-main>
      <el-footer></el-footer>
    </el-container>

  </div>
</template>

<script>

export default {
  name: 'MainHome',
  data: function () {
    return {
      keyword: '',
      pageSize:5,
      currentPage: 1,
      total: 0,
      questions: [],
    }
  },
  created() {
    this.loadAll()
  },
  methods: {
    loadAll(){
      this.request.get('v1/question/all',{
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
  },
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
  margin:auto;
}
.title{
  font-size: 20px;
  font-weight: bold;
  margin-top: 2px;
  margin-bottom: 10px;
}
.el-row {
  margin-bottom: 20px;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  margin: 0 10px;
  border-radius: 4px;
  min-height: 36px;
  color:rgba(255, 255, 255);
  font-size: 20px;
  box-sizing: border-box;
  line-height: 30px;
}
.row-bg {
  margin: 0;
  padding: 10px 0;
  background: linear-gradient(180deg,#03f,#0033fff2);
}

.question,.tag,.star{
  font-size: 15px;
  color: rgba(255, 255, 255, 0.448);
}

.answer{
  text-align: right;
}


.input-search{
  border: none;
  font-size: 15px;
  padding: 0px 5px;
}

.detail{
  text-align: center;
}


</style>