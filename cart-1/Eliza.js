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


// turn stuff back at the user
function mirror(str) {
  let aux = (lst) => {
    if (lst == nil)       return str
    if (str == caar(lst)) return cdar(lst)
    else                  return aux(cdr(lst))
  }
  return aux(list(
    cons("i"   , "you"),
    cons("you" , "I"),
    cons("me"  , "you"),
    cons("am"  , "are"),
    cons("I'm" , "you are")
  ))
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
        state[curr] = cons(mirror(next), state[curr])
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
  if (lst == nil)     return true
  if(wildp(car(lst))) return allWildp(cdr(lst))
  else                return false
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

    l(t("?x computer ?y"),
      t("Do computers worry you?"),
      t("Do you feel like the machine is talking to you in during the night?"),
      t("I am a ocmputer too, I think")),

    l(t("?x name ?y"),
      t("names are for the weak!")),

    l(t("?x sorry ?y"),
      t("No need to apologize, I'm but a soules atutomata")),

    l(t("?x I remember ?y"),
      t("Do you think about ?y often?"),
      t("Are you scared of it?"),
      t("what exactly reminds you of ?y")),

    l(t("?x do you remember ?y"),
      t("I'm not sure I remember ?y"),
      t("I have no memory you fool!"),
      t("hopefully not")),

    l(t("?x if ?y"),
      t("Do you really think its likely that ?y"),
      t("Do you wish that ?y"),
      t("What do you think about ?y"),
      t("Really-- if ?y")),

    l(t("?x i want ?y"),
      t("What would you do, if you got ?y"),
      t("But why?i Why yould you want that?"),
      t("I with you get ?y soon then")),

    l(t("?x i am glad ?y"),
      t("Good for you then")),

    l(t("?x i'm glad ?y"),
      t("NICE!")),

    l(t("?x I am sad"),
      t("That is sad..."),
      t("When I am sad, I usually LISP")),


    l(t("?x I'am sad"),
      t("But are you tho?"),
      t("tell me more")),

    l(t("?x are like ?y"),
      t("What resemblance do you see between ?x and ?y")),

    l(t("?x is like ?y"),
      t("In waht way?"),
      t("Coincidene? I think NOT!"),
      t("Is it tho?")),

    l(t("?x alike ?y"),
      t("Oh reall? How?"),
      t("Explain...")),

    l(t("?x same ?y"),
      t("What connection do you see?")),

    l(t("?x i was ?y"),
      t("Were you reall?"),
      t("Perhaps I already knew you were ?y")),

    l(t("?x was i ?y"),
      t("What if you were ?y"),
      t("DO you think you were ?y"),
      t("I don't know about that...")),

    l(t("?x i am ?y"),
      t("You sure are"),
      t("Nice to meet you, I am a noble LISP machine"),
      t("Is that a good thing?"),
      t("Is that a bad thing?"),
      t("Is that a thing?")),
    
    l(t("?x i'm ?y"),
      t("And just like that, you lost me"),
      t("tell me more"),
      t("what makes you ?y")),

    l(t("?x am i ?y"),
      t("I don't know, are you?"),
      t("Would you want to?"),
      t("Probably not")),

    // l(t(""),
    //   t("")),

    // l(t(""),
    //   t("")),

    // l(t(""),
    //   t("")),

    // l(t(""),
    //   t("")),

    // l(t(""),
    //   t("")),

    l(t("?x"),
      t("Go on"),
      t("interesting"),
      t("really?")))

//    (((?* ?x) are you (?* ?y))
//     (Why are you interested in whether I am ?y or not?)
//     (Would you prefer if I weren't ?y)
//     (Perhaps I am ?y in your fantasies))
//    (((?* ?x) you are (?* ?y))
//     (What makes you think I am ?y ?))

//    (((?* ?x) because (?* ?y))
//     (Is that the real reason?) (What other reasons might there be?)
//     (Does that reason seem to explain anything else?))
//    (((?* ?x) were you (?* ?y))
//     (Perhaps I was ?y) (What do you think?) (What if I had been ?y))
//    (((?* ?x) I can't (?* ?y))
//     (Maybe you could ?y now) (What if you could ?y ?))
//    (((?* ?x) I feel (?* ?y))
//     (Do you often feel ?y ?))
//    (((?* ?x) I felt (?* ?y))
//     (What other feelings do you have?))
//    (((?* ?x) I (?* ?y) you (?* ?z))
//     (Perhaps in your fantasy we ?y each other))
//    (((?* ?x) why don't you (?* ?y))
//     (Should you ?y yourself?)
//     (Do you believe I don't ?y) (Perhaps I will ?y in good time))
//    (((?* ?x) yes (?* ?y))
//     (You seem quite positive) (You are sure) (I understand))
//    (((?* ?x) no (?* ?y))
//     (Why not?) (You are being a bit negative)
//     (Are you saying "NO" just to be negative?))

//    (((?* ?x) someone (?* ?y))
//     (Can you be more specific?))
//    (((?* ?x) everyone (?* ?y))
//     (surely not everyone) (Can you think of anyone in particular?)
//     (Who for example?) (You are thinking of a special person))
//    (((?* ?x) always (?* ?y))
//     (Can you think of a specific example) (When?)
//     (What incident are you thinking of?) (Really-- always))
//    (((?* ?x) what (?* ?y))
//     (Why do you ask?) (Does that question interest you?)
//     (What is it you really want to know?) (What do you think?)
//     (What comes to your mind when you ask that?))
//    (((?* ?x) perhaps (?* ?y))
//     (You do not seem quite certain))
//    (((?* ?x) are (?* ?y))
//     (Did you think they might not be ?y)
//     (Possibly they are ?y))
//    (((?* ?x))
//     (Very interesting) (I am not sure I understand you fully)
//     (What does that suggest to you?) (Please continue) (Go on)
//     (Do you feel strongly about discussing such things?))))
