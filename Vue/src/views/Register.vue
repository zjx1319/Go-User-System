<template>
  <div class="register">
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
      <Input
        v-model="email"
        type="email"
        placeholder="Email"
        style="width: 300px"
      />
      <br />
      <br />
      <Button type="primary" :loading="loading" @click="register">Register</Button>
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
      loading: false,
      username: "",
      password: "",
      email: "",
    };
  },
  methods: {
    register() {
      var self = this;
      self.$Message.info("正在注册...有点慢请耐心等待");
      self.loading=true
      request({
        method: "post",
        url: "",
        data: {
          username: self.username,
          password: self.password,
          email: self.email,
        },
      })
        .then(function (response) {
          // 处理成功情况
          const userMsg = {
            userID: response.data.id,
          };
          Vue.prototype.$userMsg = userMsg;
          setToken(response.data.token);
          self.$Message.success("注册成功!请验证邮箱后登录");
          self.$router.push("/");
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "注册失败! " +
              error.response.status +
              " " +
              error.response.data.error
          );
        })
        .then(function () {
          // 总是会执行
        });
      self.loading=false
    },
  },
};
</script>
