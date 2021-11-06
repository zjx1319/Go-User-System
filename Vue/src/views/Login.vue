<template>
  <div class="login">
    <Card>
      <Input
        v-model="username"
        type="text"
        placeholder="Username"
        style="width: 300px"
      />
      <br />
      <br />
      <Input
        v-model="password"
        type="password"
        placeholder="Password"
        style="width: 300px"
      />
      <br />
      <br />
      <Button type="primary" @click="login">Login</Button>
    </Card>
  </div>
</template>

<script>
import Vue from "vue";
import request from "@/api/request";
import { setToken } from "@/api/auth.js";
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    login() {
      var self = this;
      request
        .get("/token?username=" + self.username + "&password=" + self.password)
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
    },
  },
};
</script>
