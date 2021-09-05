const path = require('path');

module.exports = {
  css: {
    loaderOptions: {
      sass: {
        data: `@import "src/assets/css/common/mixin.scss";`
      }
    }
  }
};