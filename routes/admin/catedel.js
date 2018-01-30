
var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var dbUtils = require('./models/dbUtils.js');


if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

var oid = request.formValue('oid');
cates.delByOid(oid);
res.redirect('./catelist.js');
