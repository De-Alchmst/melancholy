// https://github.com/norvig/paip-lisp/blob/main/lisp/eliza.lisp

function EEval(input) {
  return resolve(tokenize(input.toLowerCase().replace('?', '.')))
}

const t = tokenize
function tokenize(str) {
  return arr2list(str.trim().split(/[\s,.~!;]+/))
}

function resolve(tok) {
  let matched = matchWithRules(tok)
  return ljoin(" ", matched)
}


function matchWithRules(tok) {
  let aux = (rls) => {
    if (rls == nil) return nil
    else {
      let matched = match(tok, car(rls))
      if (matched != nil) return lrandom(matched)
      else                return aux(cdr(rls))
    }
  }
  return aux(RULES)
}


function match(tok, rule) {
  let patt  = car(rule)
  let resp  = cdr(rule)
  let state = {}
  console.log(lstring(patt))

  let aux   = (curr, lst, pat) => {
    if (lst == nil) {
      if (allWild(pat)) return true
      else              return false
    }

    let next = car(lst)
    if (curr.startsWith('?')) {
      if (!curr in state) state[curr] = list()

      if (pat != nil && next == car(pat))
        return aux(cadr(pat), cdr(lst), cddr(pat))
      else {
        state[curr] = cons(next, state[curr])
        return aux(curr, cdr(lst), pat)
      }
    }

    else {
      if (next == curr) return aux(car(pat), cdr(lst), cdr(pat))
      else              return false
    }
  }

  let matched = aux(car(patt), tok, cdr(patt))
  if (matched) return resp
  else         return nil
}


function allWild(lst) {
  if (lst == nil) return true
  else if(car(lst).startsWith('?')) return allWild(cdr(lst))
  else return false
}


const RULES =
  l(l(t("?x hello ?y"),
      t("How do you do. Please state your problem."),
      t("Hi, my name is Eliza; tell me about your problem.")),
    l(t("?x"),
      t("WHAT?")))

