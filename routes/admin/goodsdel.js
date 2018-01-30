
var res = require('./lib/resp.js');
var goods = require("./models/goods.js");
var dbUtils = require('./models/dbUtils.js');



if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

var oid = request.formValue('oid');
goods.delByOid(oid);
res.redirect('./goodslist.js');
