
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");

if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

var allData = cates.find();
var data = {
  tree: cates.getTree(allData),
};
res.render('./views/admin/catelist.ejs', data);