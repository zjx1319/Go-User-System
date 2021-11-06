<template>
  <div class="userInfoUpdateA">
    <Card>
      <p slot="title">修改用戶信息</p>
      <p></p>
      <p>请输入要修改的用户的ID:</p>
      <br />
      <Input v-model="ID" type="text" placeholder="ID" style="width: 300px" />
      <br />
      <br />
      <p>请输入要修改的用户的信息:</p>
      <br />
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
      <Input
        v-model="role"
        type="text"
        placeholder="Role"
        style="width: 300px"
      />
      <br />
      <br />
      <p>不修改留空即可</p>
      <p>修改邮箱后需重新验证</p>
      <br />
      <Button type="primary" @click="submit">Submit</Button>
    </Card>
  </div>
</template>

<script>
import request from "@/api/request";
export default {
  data() {
    return {
      ID: "",
      username: "",
      password: "",
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
    });
  },
  methods: {
    submit() {
      var self = this;
      request({
        method: "put",
        url: "/" + self.ID,
        data: {
          username: self.username,
          password: self.password,
          email: self.email,
          role: self.role,
        },
      })
        .then(function (response) {
          // 处理成功情况
          self.$Message.success("修改成功!");
          self.$router.push("/userinfoall");
        })
        .catch(function (error) {
          // 处理错误情况
          self.$Message.error(
            "修改失败! " +
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
