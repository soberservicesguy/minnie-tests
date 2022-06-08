const { defineConfig } = require('@vue/cli-service')
const utils = './utils'
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: utils.BASE_URL,
  }
})
