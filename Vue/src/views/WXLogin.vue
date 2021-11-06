<template>
  <div class="wxlogin">
    <p>正在通过微信登录...</p>
  </div>
</template>

<script>
import Vue from "vue";
import request from "@/api/request";
import { setToken } from "@/api/auth.js";
export default {
  data() {
    return {};
  },
  mounted: function () {
    this.$nextTick(function () {
      //调用需要执行的方法
      var self = this;
      if (typeof self.$route.query.code === "undefined") {
        self.$Message.error("登录失败! 参数不足");
        self.$router.push("/");
        return;
      }
      request
        .get("/WX/token?code=" + self.$route.query.code)
        .then(function (response) {
          // 处理成功情况
          const userMsg = {
            userID: response.data.id,
          };
          Vue.prototype.$userMsg = userMsg;
          setToken(response.data.token);
          self.$Message.success("登录成功!");
          self.$router.push("/userInfo");
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "登录失败! " +
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
