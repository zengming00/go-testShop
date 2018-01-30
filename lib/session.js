var Cart = require('./lib/Cart.class.js');

exports = module.exports = session;

function getCart() {
  var data = session.get('cart');
  data = data ? JSON.parse(data) : [];
  return new Cart(data);
}

function setCart(cart) {
  session.set('cart', JSON.stringify(cart.items()));
}

function getHistory() {
  var history = session.get('history');
  return history ? JSON.parse(history) : [];
}

function setHistory(history) {
  session.set('history', JSON.stringify(history));
}

function getUser() {
  var user = session.get('user');
  return user ? JSON.parse(user) : null;
}

function setUser(user) {
  session.set('user', JSON.stringify(user));
}


exports.getUser = getUser;
exports.setUser = setUser;
exports.getCart = getCart;
exports.setCart = setCart;
exports.getHistory = getHistory;
exports.setHistory = setHistory;