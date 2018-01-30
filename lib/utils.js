var url = require('url');
var file = require("file");
var utils = require('utils');

function fileLoader(filePath) {
  var r = file.read(filePath);
  if (r.err) {
    console.log('load:', filePath)
    throw r.err
  }
  return utils.toString(r.value);
}

// 获取与nodejs兼容的query对象
function getQuery() {
  var o = request.getUrl();
  var r = url.parseQuery(o.getRawQuery())
  if (r.err) {
    throw r.err;
  }
  var vs = r.value.getAll();
  var query = {};
  for (var k in vs) {
    if (vs[k].length > 1) {
      query[k] = vs[k];
    } else {
      query[k] = vs[k][0];
    }
  }
  return query;
}

function isEmptyObject(e) {
  if (!e) {
    return true;
  }
  for (var t in e) {
    if (e.hasOwnProperty(t)) {
      return false;
    }
  }
  return true;
}

function isNullOrUndefined(v) {
  if (v === null || v === undefined) {
    return true;
  }
  return false;
}

function isNotNil(v) {
  return !isNullOrUndefined(v);
}

function toInt(v, def) {
  v = parseInt(v);
  return isNaN(v) ? def : v;
}

function encodePassword(pwd, salt) {
  return utils.md5(pwd + salt);
}

exports.encodePassword = encodePassword;
exports.fileLoader = fileLoader;
exports.toInt = toInt;
exports.isEmptyObject = isEmptyObject;
exports.getQuery = getQuery;
exports.isNullOrUndefined = isNullOrUndefined;
exports.isNotNil = isNotNil;
