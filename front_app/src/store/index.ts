import { defineStore } from "pinia";
import { UserInfo } from "@/common/types";

export const useStore = defineStore("main", {
  state: () => {
    return {
      user: <UserInfo>{},
      isLogin: false,
    };
  },
});
