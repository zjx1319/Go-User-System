<template>
  <div class="userInfoAll">
    <Card>
      <p slot="title">所有用户信息</p>
      <p></p>
      <Card v-for="user in users">
        <p>用户ID: {{ user.id }}</p>
        <p>用户名: {{ user.username }}</p>
        <p>邮箱地址: {{ user.email }}</p>
        <p>权限组: {{ user.role }}</p>
      </Card>
    </Card>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {
      users: [],
    };
  },
  mounted: function () {
    this.$nextTick(function () {
      //调用需要执行的方法
      var self = this;
      if (typeof self.$userMsg === "undefined") {
        self.$Message.error("未登录");
        self.$router.push("/");
        return;
      }
      request
        .get("")
        .then(function (response) {
          // 处理成功情况
          self.users = response.data;
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "获取信息失败! " +
              error.response.status +
              " " +
              error.response.data.error
          );
        })
        .then(function () {
          // 总是会执行
        });
    });
  },
};
</script>
