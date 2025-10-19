// https://github.com/norvig/paip-lisp/blob/main/lisp/eliza.lisp

function EEval(input) {
  return resolve(tokenize(input.toLoverCase()))
}

const t = tokenize
function tokenize(str) {
  return arr2list(str.trim().split(/\s+/))
}


function resolve(tok) {
  let aux = (lst) => {
    if (length(lst) == 0) return ""
    else {
      let matched = match(car(lst), RULES)
      if (matched) return lrandom(matched)
      else         return aux(cdr(lst))
    }
  }
  return aux(tok)
}


function match(lst, rules) {
  let stored = {}
}

const RULES = l(
  l(l(t("?x hello ?y"),
      t("How do you do. Please state your problem."),
      t("Hi, my name is Eliza; tell me about your problem.")))
)

