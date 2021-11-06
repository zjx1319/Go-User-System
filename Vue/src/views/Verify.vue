<template>
  <div class="verify">
    <p>正在验证邮箱...</p>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {};
  },
  mounted: function () {
    this.$nextTick(function () {
      //调用需要执行的方法
      var self = this;
      if (
        typeof self.$route.query.id === "undefined" ||
        typeof self.$route.query.code === "undefined"
      ) {
        self.$Message.error("验证失败! 参数不足");
        self.$router.push("/");
        return;
      }
      request({
        method: "post",
        url: "/email",
        data: {
          id: parseInt(self.$route.query.id),
          code: self.$route.query.code,
        },
      })
        .then(function (response) {
          // 处理成功情况
          self.$Message.success("验证成功!");
          self.$router.push("/");
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "验证失败! " +
              error.response.status +
              " " +
              error.response.data.error
          );
          self.$router.push("/");
        })
        .then(function () {
          // 总是会执行
        });
    });
  },
};
</script>
