
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var dbUtils = require('./models/dbUtils.js');

if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

if (request.getMethod() === 'GET') {
  var allData = cates.find();
  var data = {
    tree: cates.getTree(allData),
  };
  res.render('./views/admin/cateadd.ejs', data);
  return;
}

if (request.getMethod() === 'POST') {
  var data = {
    cat_name: request.formValue('cat_name'),
    intro: request.formValue('cat_desc'),
    parent_id: request.formValue('parent_id'),
    oid: dbUtils.makeOid(),
  };
  cates.add(data);
  res.redirect('./catelist.js');
}
