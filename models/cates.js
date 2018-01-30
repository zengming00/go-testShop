var types = require('types');
var db = require('./models/db.js');
var dbUtils = require('./models/dbUtils.js');

var db = db.getConn();



exports.getTree = function getTree($rows, $pid, $level) {
  $pid = $pid || 0;
  $level = $level || 0;

  var $tree = [];
  for (var k in $rows) {
    var $row = $rows[k];
    if ($row.parent_id == $pid) {
      $tree.push({  //$row是模型实体对象无法添加level属性
        oid: $row.oid,
        cat_name: $row.cat_name,
        intro: $row.intro,
        parent_id: $row.parent_id,
        level: $level
      });
      //if($row.oid != $row.parent_id){ //原本以为会死循环，但实际上不会出现
      $tree = $tree.concat(getTree($rows, $row.oid, $level + 1));
      //}
    }
  }
  return $tree;
}

exports.getFamily = function getFamily($rows, $catid) {
  var $arr = [], k, row, isFind;
  while ($catid != 0) {
    isFind = false;
    for (k in $rows) {
      row = $rows[k];
      if (row.oid == $catid) {
        $arr.unshift(row);
        $catid = row.parent_id;
        isFind = true;//避免死循环
        break;
      }
    }
    if (!isFind) break;
  }
  return $arr;
}

exports.getChildCates = function getChilds($rows, $catid) {
  var arr = [], k, r;
  for (k in $rows) {
    r = $rows[k];
    if (r.parent_id == $catid) {
      arr.push(r.oid);
      arr = arr.concat(getChilds($rows, r.oid));
    }
  }
  return arr;
}

exports.find = function find() {
  var ret = [];
  var r = db.query('select * from cates');
  if (r.err) {
    throw r.err;
  }
  var rows = r.value;
  while (rows.next()) {
    var id = types.newInt();
    var oid = types.newString();
    var cat_name = types.newString();
    var intro = types.newString();
    var parent_id = types.newString();
    var created_at = types.newString();

    var err = rows.scan(id, oid, cat_name, intro, parent_id, created_at);
    if (err) {
      rows.close();
      throw err;
    }

    ret.push({
      id: types.intValue(id),
      oid: types.stringValue(oid),
      cat_name: types.stringValue(cat_name),
      intro: types.stringValue(intro),
      parent_id: types.stringValue(parent_id),
      created_at: types.stringValue(created_at),
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

exports.add = function add(data) {
  var r = dbUtils.exportKeyValues(data);
  var sql = dbUtils.makeInsertSql('cates', r.keys);
  return dbUtils.insert(db, sql, r.values);
}

exports.delByOid = function delByOid(oid) {
  return dbUtils.delete(db, 'delete from cates where oid=?', [oid]);
}

exports.getByOid = function getByOid(oid) {
  var ret = null;
  var r = db.query('select * from cates where oid = ?', oid);
  if (r.err) {
    throw r.err;
  }
  var rows = r.value;
  while (rows.next()) {
    var id = types.newInt();
    var oid = types.newString();
    var cat_name = types.newString();
    var intro = types.newString();
    var parent_id = types.newString();
    var created_at = types.newString();

    var err = rows.scan(id, oid, cat_name, intro, parent_id, created_at);
    if (err) {
      throw err;
    }

    ret = {
      id: types.intValue(id),
      oid: types.stringValue(oid),
      cat_name: types.stringValue(cat_name),
      intro: types.stringValue(intro),
      parent_id: types.stringValue(parent_id),
      created_at: types.stringValue(created_at),
    };
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

exports.updateByOid = function updateByOid(oid, data) {
  var r = dbUtils.exportKeyValues(data)
  var sql = dbUtils.makeUpdateSql('cates', r.keys, ['oid'])
  // 添加where条件值
  r.values.push(oid)
  return dbUtils.update(db, sql, r.values)
}