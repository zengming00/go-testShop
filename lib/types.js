
function isString(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object String]');
}

function isNumber(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Number]')
}

function isBoolean(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Boolean]')
}

function isUndefined(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Undefined]')
}

function isNull(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Null]')
}

function isObject(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Object]')
}

function isArray(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Array]')
}

function isFunction(v) {
  var gettype = Object.prototype.toString;
  return (gettype.call(v) === '[object Function]')
}

exports.isString = isString;
exports.isNumber = isNumber;
exports.isBoolean = isBoolean;
exports.isUndefined = isUndefined;
exports.isNull = isNull;
exports.isObject = isObject;
exports.isArray = isArray;
exports.isFunction = isFunction; 
