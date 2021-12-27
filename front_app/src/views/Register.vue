<template>
  <el-card class="login">
    <template #header>
      <span style="text-align: center; display: block; font-size: 2rem">注 册</span>
    </template>
    <el-form :model="userinfo" ref="form" :rules="LoginOrRegisterRules" label-width="90px">
      <el-form-item label="用户名" prop="username">
        <el-input v-model="userinfo.username" placeholder="输入用户名" clearable></el-input>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="userinfo.password" placeholder="输入密码" clearable></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onRegister">注册</el-button>
        <router-link to="/">
          <el-button>登录</el-button>
        </router-link>
      </el-form-item>
    </el-form>
  </el-card>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import { LoginOrRegisterRules } from "@/common/rules";
import { UserInfo } from "@/common/types";
import { ElFormItemContext, ElMessage } from "element-plus";
import { register } from "@/service/api";
import { useRouter } from "vue-router";

const router = useRouter();

const userinfo = ref<UserInfo>({ username: "", password: "" });
const formRef = ref<ElFormItemContext>();
const onRegister = async () => {
  try {
    await formRef.value?.validate("blur");
  } catch (_) {
    return;
  }
  register(userinfo.value, (res: any) => {
    if (res.code === 200 && res.token !== "") {
      window.localStorage.setItem("token", res.token);
      ElMessage.success("注册成功");
      router.push("/filemanage");
    }
  });
};
</script>
<style lang="less">
.login {
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  padding: 3rem;
}
</style>
