<template>
  <div class="wxbind">
    <Card>
      <p slot="title">请输入要绑定的账号和密码</p>
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
      <Button type="primary" @click="bind">Bind</Button>
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
  mounted: function () {
    this.$nextTick(function () {
      //调用需要执行的方法
      var self = this;
      if (typeof self.$route.query.code === "undefined") {
        self.$Message.error("绑定失败! 参数不足");
        self.$router.push("/");
        return;
      }
    });
  },
  methods: {
    bind() {
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

          //绑定微信
          request({
            method: "post",
            url: "/WX",
            data: {
              code: self.$route.query.code,
            },
          })
            .then(function (response) {
              // 处理成功情况
              self.$Message.success("已绑定微信:" + response.data.wx_name);

              self.$router.push("/userInfo");
            })
            .catch(function (error) {
              // 处理错误情况
              self.$Message.error(
                "绑定失败! " +
                  error.response.status +
                  " " +
                  error.response.data.error
              );
            });
          //OK
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
