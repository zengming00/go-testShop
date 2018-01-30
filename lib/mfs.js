var os = require('os')
var utils = require('utils')
var file = require('file')

// 自己实现nodejs的fs模块
var fs = require_get('fs')
if (!fs) {
  var myExports = {
    readFileSync: function (name) {
      var r = file.read(name);
      if (r.err) {
        throw r.err
      }
      return utils.toString(r.value);
    },
    existsSync: function (name) {
      var r = os.stat(name)
      return !r.err;
    },
    test: function (str) {
      console.log("---------------");
      console.log(str);
      console.log("---------------");
    }
  }
  var myModule = {
    exports: myExports,
  };
  require_set('fs', myModule);
  return;
}
console.log('fs exists!');