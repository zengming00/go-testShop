
function verify(yzm) {
    var v = session.get('__verify');
    if (v) {
        if (v === yzm.toUpperCase()) {
            //清空，防止多次使用
            session.set('__verify', null);
            return true;
        }
    }
    return false;
}

exports.verify = verify;
