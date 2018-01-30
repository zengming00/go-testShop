
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var ordinfos = require("./models/ordinfos.js");
var Common = require('./lib/Common.js');

if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

var allData = cates.find();
var data = {
  tree: cates.getTree(allData),
  ords: ordinfos.find({}, { sort: { id: 'desc' } }),
  toDateStr: Common.toDateStr,
};

res.render('./views/admin/ordlist.ejs', data);