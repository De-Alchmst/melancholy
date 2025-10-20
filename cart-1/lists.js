/*

        ~::CONS::~
Who needs objects anyways?

+---+---+   +---+---+
| O | O---->| O | O---->nil
+-|-+---+   +-|-+---+
  |           |
  V           V
  4           2
*/
const nil = null

function cons(h, t) {
  return {
    "contents of the address part of register number":   h,
    "contents of the decrement part of register number": t,
  }
}

function car(cell) {
  return cell["contents of the address part of register number"]
}

function cdr(cell) {
  return cell["contents of the decrement part of register number"]
}


function consp(x) {
  return x &&
         "contents of the address part of register number"   in x &&
         "contents of the decrement part of register number" in x
}


function arr2list(argv) {
  let res = nil
  argv.forEach((s) => {
    res = cons(s, res)
  })
  return lrev(res)
}


const l = list
function list(...argv) {
  return arr2list(argv)
}


function length(l) {
  let aux = (acc, l) => {
    if (!consp(l)) return acc
    else           return aux(acc+1, cdr(l))
  }
  return aux(0, l)
}


function lstring(l) {
  let aux = (str, first, lst) => {
    if (lst == nil) return str + ")"
    else if (first) return aux(str +        String(car(lst)), false, cdr(lst))
    else            return aux(str + ", " + String(car(lst)), false, cdr(lst))
  }
  return aux("list(", true, l)
}


function lmap(l, f) {
  if (l == nil) return nil
  else          return cons(f(car(l)), lmap(cdr(l), f))
}


function lrev(lst) {
  let aux = (acc, l) => {
    if (!consp(l)) return acc
    else return aux(cons(car(l), acc), cdr(l))
  }
  return aux(nil, lst)
}


function ljoin(sep, lst) {
  let aux = (acc, ls) => {
    if (ls == nil) return acc
    else           return aux(acc + sep + car(ls), cdr(ls))
  }
  return aux("", lst)
}


function lrandom(lst) {
  let aux = (n, ls) => {
    if (n <= 0) return ls
    else        return aux(n-1, cdr(ls))
  }
  return car(aux(Math.floor((Math.random() * length(lst)) -0.00000001), lst))
}


function caar  (l) {return car(car(l))}
function cadr  (l) {return car(cdr(l))}
function cdar  (l) {return cdr(car(l))}
function cddr  (l) {return cdr(cdr(l))}
function caaar (l) {return caar(car(l))}
function caadr (l) {return caar(cdr(l))}
function cadar (l) {return cadr(car(l))}
function caddr (l) {return cadr(cdr(l))}
function cdaar (l) {return cdar(car(l))}
function cdadr (l) {return cdar(cdr(l))}
function cddar (l) {return cddr(car(l))}
function cdddr (l) {return cddr(cdr(l))}
function caaaar(l) {return caaar(car(l))}
function caaadr(l) {return caaar(cdr(l))}
function caadar(l) {return caadr(car(l))}
function caaddr(l) {return caadr(cdr(l))}
function cadaar(l) {return cadar(car(l))}
function cadadr(l) {return cadar(cdr(l))}
function caddar(l) {return caddr(car(l))}
function cadddr(l) {return caddr(cdr(l))}
function cdaaar(l) {return cdaar(car(l))}
function cdaadr(l) {return cdaar(cdr(l))}
function cdadar(l) {return cdadr(car(l))}
function cdaddr(l) {return cdadr(cdr(l))}
function cddaar(l) {return cddar(car(l))}
function cddadr(l) {return cddar(cdr(l))}
function cdddar(l) {return cdddr(car(l))}
function cddddr(l) {return cdddr(cdr(l))}
