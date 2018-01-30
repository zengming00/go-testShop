
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var dbUtils = require('./models/dbUtils.js');



if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

if (request.getMethod() === 'GET') {
  var oid = request.formValue('oid');
  if (!oid) {
    return res.send('oid is null');
  }
  var allData = cates.find();
  var item = cates.getByOid(oid);
  var data = {
    tree: cates.getTree(allData),
    cate: item,
  };
  res.render('./views/admin/cateedit.ejs', data);
  return;
}

if (request.getMethod() === 'POST') {
  var cat_id = request.formValue('cat_id');
  var data = {
    cat_name: request.formValue('cat_name'),
    intro: request.formValue('cat_desc'),
    parent_id: request.formValue('parent_id'),
  };
  if (cat_id === data.parent_id) {
    return res.send("错误的上级栏目");
  }
  cates.updateByOid(cat_id, data);
  res.redirect('./catelist.js');
}
