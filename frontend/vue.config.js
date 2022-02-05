// const path = require('path');

if (process.env.NODE_ENV !== undefined) {
  if (!process.env.VUE_APP_API_BASE_URL) {
    throw 'API_BASE_URL に API の URL を指定してください'
  }
}

module.exports = {
  devServer: {
  },

  css: {
    loaderOptions: {
      sass: {
        data: `@import "src/assets/css/common/mixin.scss";`
      }
    }
  }
};