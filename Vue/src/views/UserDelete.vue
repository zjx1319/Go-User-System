<template>
  <div class="userDelete">
    <Card>
      <p slot="title">删除用户账号</p>
      <p></p>
      <p>请输入要删除的用户的ID:</p>
      <br />
      <Input v-model="ID" type="text" placeholder="ID" style="width: 300px" />
      <br />
      <br />
      <Button type="error" @click="submit">Delete</Button>
    </Card>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {
      ID: "",
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
    });
  },
  methods: {
    submit() {
      var self = this;
      request({
        method: "delete",
        url: "/" + self.ID,
      })
        .then(function (response) {
          // 处理成功情况
          self.$Message.success("删除成功!");
          self.$router.push("/");
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "删除失败! " +
              error.response.status +
              " " +
              error.response.data.error
          );
        })
        .then(function () {
          // 总是会执行
        });
    },
  },
};
</script>
