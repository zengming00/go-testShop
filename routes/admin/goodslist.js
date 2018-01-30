var res = require('./lib/resp.js');
var cates = require("./models/cates.js");
var goods = require("./models/goods.js");
var dbUtils = require('./models/dbUtils.js');
var utils = require('./lib/utils.js');
var Page = require('./lib/Page.class.js');
var Common = require('./lib/Common.js');


if (!session.get("isAdmin")) {
  return res.redirect('/routes/admin/login.js');
}

if (request.getMethod() === 'GET') {
  var query = utils.getQuery()
  var cat_id = query.cat_id
  var sortTime = query.sortTime || 'D';
  var tplData = {//传递给模板的数据
    cat_id: cat_id,
    sortTime: sortTime,
    toDateStr: Common.toDateStr,
  };

  var total = 0;
  if (cat_id) {
    total = goods.count({ cat_id: cat_id });
  } else { //未传递栏目ID，查找所有
    total = goods.count();
  }
  var req = {
    originalUrl: request.getUri(),
    query: query,
  }
  var page = new Page(req, total, 8);
  var allCates = cates.find();
  tplData.tree = cates.getTree(allCates);
  tplData.page = page.show();

  sortTime = sortTime == 'D' ? 'desc' : 'asc';
  var opt = { sort: { id: sortTime }, skip: page.firstRow, limit: page.listRows };
  if (cat_id) {
    tplData.goods = goods.find({ cat_id: cat_id }, opt);
  } else {
    tplData.goods = goods.find({}, opt);
  }
  if (query.json) {
    res.json(tplData);
    return;
  }
  res.render('./views/admin/goodslist.ejs', tplData);
}

