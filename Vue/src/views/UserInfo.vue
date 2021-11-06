<template>
  <div class="userInfo">
    <Card>
      <p slot="title">个人信息</p>
      <p></p>
      <p>用户ID: {{ id }}</p>
      <p>用户名: {{ username }}</p>
      <p>邮箱地址: {{ email }}</p>
      <p>权限组: {{ role }}</p>
    </Card>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {
      id: "",
      username: "",
      email: "",
      role: "",
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
        .get("/" + self.$userMsg.userID)
        .then(function (response) {
          // 处理成功情况
          self.id = response.data.id;
          self.username = response.data.username;
          self.email = response.data.email;
          self.role = response.data.role;
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
