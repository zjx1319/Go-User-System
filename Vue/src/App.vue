<template>
  <div>
    <Menu mode="horizontal" @on-select="menuSelect">
      <MenuItem name="1">
        <Icon type="ios-home" />
        首页
      </MenuItem>
      <Submenu name="2">
        <template slot="title">
          <Icon type="ios-person" />
          用户中心
        </template>
        <MenuItem name="2-1">注册账号</MenuItem>
        <MenuItem name="2-2">登录账号</MenuItem>
        <MenuItem name="2-3">个人信息</MenuItem>
        <MenuItem name="2-4">修改信息</MenuItem>
        <MenuGroup title="admin">
          <MenuItem name="2-5">所有用户信息</MenuItem>
          <MenuItem name="2-6">修改用戶信息</MenuItem>
          <MenuItem name="2-7">删除用户账号</MenuItem>
        </MenuGroup>
        <MenuGroup title="wx">
          <MenuItem name="2-8">微信登录</MenuItem>
          <MenuItem name="2-9">绑定微信</MenuItem>
          <MenuItem name="2-10">查询绑定</MenuItem>
        </MenuGroup>
      </Submenu>
      <MenuItem name="3">
        <Icon type="ios-more" />
        其他
      </MenuItem>
    </Menu>
    <router-view class="main"></router-view>
  </div>
</template>
<script>
export default {
  methods: {
    menuSelect(menuName) {
      if (menuName == "1") {
        this.$router.push("/");
      } else if (menuName == "2-1") {
        this.$router.push("/register");
      } else if (menuName == "2-2") {
        this.$router.push("/login");
      } else if (menuName == "2-3") {
        this.$router.push("/userinfo");
      } else if (menuName == "2-4") {
        this.$router.push("/userInfoUpdate");
      } else if (menuName == "2-5") {
        this.$router.push("/userinfoall");
      } else if (menuName == "2-6") {
        this.$router.push("/userInfoUpdateA");
      } else if (menuName == "2-7") {
        this.$router.push("/userDelete");
      } else if (menuName == "2-8") {
        this.$Message.error("请先在后台设置!");
        //这里要放微信登录的链接
        //window.location.href =
          "";
      } else if (menuName == "2-9") {
        this.$Message.error("请先在后台设置!");
        //这里要放微信登录的链接
        //window.location.href =
          "";
      } else if (menuName == "2-10") {
        this.$router.push("/userInfoWX");
      } else if (menuName == "3") {
        this.$router.push("/about");
      }
    },
  },
  mounted() {
    document.title = "用户管理系统";
  },
};
</script>
<style>
.main {
  text-align: center;
  background-color: #fff;
  border-radius: 20px;
  width: 350px;
  height: 350px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
</style>
