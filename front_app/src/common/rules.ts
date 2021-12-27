import { FormItemRule } from "element-plus/es/components/form/src/form.type";

export const LoginOrRegisterRules = {
  username: [
    {
      required: true,
      message: "输入用户名",
      trigger: "blur",
    },
    {
      min: 5,
      max: 20,
      message: "长度在5-20之间",
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      message: "输入密码",
      trigger: "blur",
    },
    {
      min: 5,
      max: 20,
      message: "长度在5-20之间",
      trigger: "blur",
    },
  ],
} as Partial<Record<string, FormItemRule | FormItemRule[]>>;
