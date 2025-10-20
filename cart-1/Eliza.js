// https://github.com/norvig/paip-lisp/blob/main/lisp/eliza.lisp

function EEval(input) {
  return resolve(tokenize(input.toLowerCase()))
}

const t = tokenize
function tokenize(str) {
  return arr2list(str.trim().split(/[\s,.~!?]+/))
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
  let pat  = car(rule)
  let resp = cdr(rule)
  console.log(pat)
  console.log(resp)
  return resp
}


const RULES =
  l(l(t("?x hello ?y"),
      t("How do you do. Please state your problem."),
      t("Hi, my name is Eliza; tell me about your problem.")),
    l(t("?x"),
      t("WHAT?")))

