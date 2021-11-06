<template>
  <div class="userInfoWX">
    <Card>
      <p slot="title">微信信息</p>
      <p></p>
      <p>已绑定微信昵称: {{ wxname }}</p>
    </Card>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {
      wxname: "",
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
        .get("/WX")
        .then(function (response) {
          // 处理成功情况
          self.wxname = response.data.wx_name;
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
