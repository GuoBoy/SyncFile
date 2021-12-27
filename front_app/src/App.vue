<template>
  <el-container v-if="isLogin">
    <el-container>
      <el-aside width="auto">
        <!-- Aside content -->
        <el-menu
          class="el-menu-vertical-demo"
          :collapse="isCollapse"
          router
          :default-active="currentMenu"
        >
          <el-menu-item v-for="m in menus" :key="m.path" :index="m.path">
            <i :class="['bi', `bi-${m.icon}`]" class="menu-icon"></i>
            <span>{{ m.title }}</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header style="height: 2rem" class="main-header">
          <!-- Header content -->
          <i
            class="bi bi-text-indent-left"
            @click="isCollapse = !isCollapse"
            v-if="isCollapse"
            style="font-size: 2rem; color: cornflowerblue"
          ></i>
          <i
            class="bi bi-text-indent-right"
            @click="isCollapse = !isCollapse"
            v-else
            style="font-size: 2rem; color: cornflowerblue"
          ></i>
          <div class="user">
            <el-avatar
              src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
              class="avater"
            ></el-avatar>
            <el-dropdown>
              <span>用户名</span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item>信息设置</el-dropdown-item>
                  <el-dropdown-item divided>退出登录</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-header>

        <el-main>
          <!-- Main content -->
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
    <el-footer style="background-color: #ff5500; height: 25px">
      <!-- Footer content -->
      底部状态栏
    </el-footer>
  </el-container>
  <div v-else>
    <router-view></router-view>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import { Menu } from "@/common/types";
import { useStore } from "@/store";
import { useRouter } from "vue-router";

const store = useStore();

const isLogin = computed(() => store.isLogin);

const isCollapse = ref(false);
const menus = ref<Menu[]>([
  { path: "filemanage", title: "文件管理", icon: "list-task" },
  { path: "uploadlist", title: "上传列表", icon: "upload" },
  { path: "setting", title: "设置", icon: "sliders" },
]);
const currentMenu = ref(menus.value[0].path);
const router = useRouter();
router.push("/");
</script>

<style lang="less">
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
html,
body,
#app {
  height: 100%;
  margin: 0;
  padding: 0;
}
.el-container,
.el-menu {
  height: 100%;
}
.main-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #ff0055;
}
.menu-icon {
  font-size: 1.5rem;
  margin-right: 0.5rem;
}
.user {
  display: flex;
  align-items: center;
  .avater {
    margin-right: 0.5rem;
  }
}
</style>
