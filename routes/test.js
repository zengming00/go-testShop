var res = require('./lib/resp.js');
// var os = require('os');
// var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
// var Dirs = require("./lib/Dirs.js");
var utils = require("./lib/utils.js");
var session = require('./lib/session.js');

// var c = cates.find();
// c = cates.getTree(c)

// var r = Dirs.getDateDir()


var oid = '58843e737ce9c30011b39ad2';
// var oid = '58843e737ce9c30011b39ad4';
var num = 1;
var r = goods.decrGoodsNum(oid, num);
console.log('rowsAffected:', r.rowsAffected);
res.json(r);

