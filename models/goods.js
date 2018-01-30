var types = require('types');
var mtypes = require('./lib/types.js');
var db = require('./models/db.js');
var dbUtils = require('./models/dbUtils.js');
var utils = require('./lib/utils.js');

var db = db.getConn();


exports.count = function count(where) {
  var ret = null;
  if (where) {
    var r = dbUtils.exportKeyValues(where);
    var fields = r.keys;
    var values = r.values;
    var args = [];
    var arr = [];
    for (var i = 0; i < fields.length; i++) {
      var k = fields[i];
      var v = values[i];
      if (mtypes.isObject(v)) {
        if (v.$in) {
          var ins = v.$in;
          arr.push(k + ' in (' + dbUtils.makePlaceStr(ins.length) + ')');
          args = args.concat(ins);
        }
        continue;
      }
      arr.push(k + '=?');
      args.push(v);
    }
    var sql = 'select count(*) from goods where ' + arr.join(' and ');
    args.unshift(sql); // 将sql插入到最前面
    console.log(JSON.stringify(args));
    ret = db.query.apply(db, args);
  } else {
    ret = db.query('select count(*) from goods');
  }
  if (ret.err) {
    throw ret.err;
  }
  var rows = ret.value;
  while (rows.next()) {
    var n = types.newInt();
    var err = rows.scan(n);
    if (err) {
      throw err;
    }
    ret = types.intValue(n);
    break;
  }
  var err = rows.err();
  if (err) {
    throw err;
  }
  err = rows.close();
  if (err) {
    throw err;
  }
  return ret;
}

exports.add = function add(data) {
  var r = dbUtils.exportKeyValues(data);
  var sql = dbUtils.makeInsertSql('goods', r.keys);
  return dbUtils.insert(db, sql, r.values);
}

exports.delByOid = function delByOid(oid) {
  return dbUtils.delete(db, 'delete from goods where oid=?', [oid]);
}

exports.getByOid = function getByOid(oid) {
  var gs = find({ oid: oid }, {});
  if (gs.length > 0) {
    return gs[0];
  }
  return null;
}

function find(where, opt) {
  var sql = 'select * from goods';
  return dbUtils.find(sql, where, opt, query);
}

function query(sql, params) {
  var ret = [];
  params.unshift(sql);
  console.log('%j', params)
  var r = db.query.apply(db, params);
  if (r.err) {
    throw r.err;
  }
  var rows = r.value;
  while (rows.next()) {
    var id = types.newInt();
    var oid = types.newString();
    var goods_name = types.newString();
    var cat_id = types.newString();
    var shop_price = types.newString();
    var goods_img = types.newString();
    var goods_desc = types.newString();
    var goods_number = types.newInt();
    var is_best = types.newInt();
    var is_new = types.newInt();
    var is_hot = types.newInt();
    var is_on_sale = types.newInt();
    var created_at = types.newString();
    var updated_at = types.newString();

    var err = rows.scan(id, oid, goods_name, cat_id, shop_price, goods_img, goods_desc, goods_number, is_best, is_new, is_hot, is_on_sale, created_at, updated_at);
    if (err) {
      rows.close();
      throw err;
    }

    ret.push({
      id: types.intValue(id),
      oid: types.stringValue(oid),
      goods_name: types.stringValue(goods_name),
      cat_id: types.stringValue(cat_id),
      shop_price: types.stringValue(shop_price),
      goods_img: types.stringValue(goods_img),
      goods_desc: types.stringValue(goods_desc),
      goods_number: types.intValue(goods_number),
      is_best: types.intValue(is_best),
      is_new: types.intValue(is_new),
      is_hot: types.intValue(is_hot),
      is_on_sale: types.intValue(is_on_sale),
      created_at: types.stringValue(created_at),
      updated_at: types.stringValue(updated_at),
    });
  }
  var err = rows.err();
  if (err) {
    rows.close();
    throw err;
  }
  err = rows.close();
  if (err) {
    throw err;
  }
  return ret;
}

function updateByOid(oid, data) {
  var r = dbUtils.exportKeyValues(data)
  var sql = dbUtils.makeUpdateSql('goods', r.keys, ['oid'])
  // 添加where条件值
  r.values.push(oid)
  return dbUtils.update(db, sql, r.values)
}

function decrGoodsNum(oid, num) {
  var sql = 'update goods set goods_number=goods_number-? where oid=? and goods_number>=?';
  var args = [num, oid, num];
  return dbUtils.update(db, sql, args);
}

exports.decrGoodsNum = decrGoodsNum;
exports.find = find;
exports.updateByOid = updateByOid;
