// const path = require('path');

if (process.env.NODE_ENV !== undefined) {
  if (!process.env.VUE_APP_API_BASE_URL) {
    throw 'API_BASE_URL に API の URL を指定してください'
  }
}

module.exports = {
  devServer: {
    // heroku用
    disableHostCheck: true
  },

  // 型チェックの際のメモリ利用寮の調整
  chainWebpack: (config) => {
    config.plugin("fork-ts-checker").tap((args) => {
      args[0].memoryLimit = 520;
      return args;
    });
  },

  css: {
    loaderOptions: {
      sass: {
        data: `@import "src/assets/css/common/mixin.scss";`,
        implementation: require('sass'),
      }
    }
  }
};