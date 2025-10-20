// https://github.com/norvig/paip-lisp/blob/main/lisp/eliza.lisp

function EEval(input) {
  return resolve(tokenize(input.toLowerCase().replace(/[,.~!;?]/, '')))
}

// string to list of tokens
const t = tokenize
function tokenize(str) {
  return arr2list(str.trim().split(/[\s]+/).filter(x => x != ""))
}


function resolve(tok) {
  let matched = matchWithRules(tok)
  return lstrJoin(" ", matched)
}


// finds which rule matches given tokens and returns random response from it,
// or nil if no match (won't happen, since last in RULES t("?x"))
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


// matches against specific rule and keep track of saved wildcards
// if matched, returns list of substituted responses, else return nil
function match(tok, rule) {
  let patt  = car(rule)
  let resp  = cdr(rule)
  let state = {}

  let aux = (curr, lst, pat) => {
    if (lst == nil) {
      if (allWildp(pat)) return true // some patterns might be left empty
      else               return false
    }

    let next = car(lst)
    if (wildp(curr)) { // if wildcard, store empty list if first time
      if (!curr in state) state[curr] = list()

      // if token after wildcard matches, wildcard ends, skip the matched, and
      // continue
      if (pat != nil && next == car(pat))
        return aux(cadr(pat), cdr(lst), cddr(pat))
      else { // else prepend the current token to state
        state[curr] = cons(next, state[curr])
        return aux(curr, cdr(lst), pat)
      }
    }

  else { // tamecards, return false if not matching, else just continue forwards
      if (next == curr) return aux(car(pat), cdr(lst), cdr(pat))
      else              return false
    }
  }

  let matched = aux(car(patt), tok, cdr(patt))
  if (matched) return patch(state, resp)
  else         return nil
}


function wildp(str) {
  return str.startsWith('?')
}


function allWildp(lst) {
  if (lst == nil) return true
  else if(wildp(car(lst))) return allWildp(cdr(lst))
  else return false
}


// map wildcard matches to responses
function patch(state, lists) {
  let patchList = (acc, lst) => {
    // all reversed until returned
    if (lst == nil) return lrev(acc)

    let frst = car(lst) // state[frst] might be empty, if so use empty list
    if (wildp(frst)) return patchList(ljoin(state[frst] ?? l(), acc), cdr(lst))
    else             return patchList(cons(frst, acc), cdr(lst))
  }

  return lmap(lists, (lst) => patchList(list(), lst))
}


// list of lists of lists of tokens
// first token list pattern, rest responses
// tokens starting with '?' are wild
const RULES =
  l(l(t("?x hello ?y"),
      t("How do you do. Please state your problem."),
      t("Hi, my name is Eliza; tell me about your problem.")),

    l(t("?x hi ?y"),
      t("Sup!"),
      t("Sup~"),
      t("May the gods be with you")),

    l(t("?x what is ?y"),
      t("I truly do not know what ?y is")),

    l(t("?x"),
      t("WHAT?")))

