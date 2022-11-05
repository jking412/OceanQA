<template>
  <div>
    <el-container>
        <el-row type="flex" class="row-bg" justify="start">
          <el-col :span="4"><div class="grid-content answer" @click="toHome">OceanQA</div></el-col>
          <el-col :span="0.5"><div class="grid-content question" @click="toHome">问题</div></el-col>
          <el-col :span="0.5"><div class="grid-content tag">标签</div></el-col>
          <el-col :span="0.5"><div class="grid-content star">收藏</div></el-col>
          <el-col :span="4" class="input-content" :push="4"><el-input class="input-search" v-model="keyword" placeholder="请输入内容"
            prefix-icon="el-icon-search"></el-input></el-col>
          <el-col :span="1" :push="8"><el-button type="primary" @click="toAdd">添加问题</el-button></el-col>

        </el-row>
      <el-main>
          <router-view></router-view>
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
    }
  },

  methods: {
      toHome(){
        this.$router.push({path: '/'})
      },
      toAdd(){
        this.$router.push({path: '/add'})
      },
  },
}
</script>

<style scoped>
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
  position: sticky;
  top: 0;
  z-index: 1;
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



</style>