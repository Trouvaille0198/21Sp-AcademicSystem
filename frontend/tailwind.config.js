module.exports = {
  purge: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}'  //包含了src文件夹下所有的vue,js等等文件
  ],
  content: [],
  theme: {
    extend: {},
  },
  plugins: [],
  shortcuts: {
    // 普通标题
    's-title': 'text-lg font-semibold text-primary',
    // 下划线
    's-underline': 'underline decoration-primary-active decoration-4 underline-offset-3',
    // 分隔线
    's-divide': 'bg-primary-active w-full py-0.5',
    // 卡片样式
    's-card': 'rounded-lg 3xl:rounded-xl shadow-md bg-white dark:bg-gray-800',
  }
}
