  /////////////////////////////////////////////////////////////////////
 /// https://github.com/norvig/paip-lisp/blob/main/lisp/eliza.lisp ///
/////////////////////////////////////////////////////////////////////

function EEval(input) {
  return resolve(tokenize(input.toLowerCase().replace(/[,.~!;?]/, '')))
}

// string to list of tokens
const t = tokenize
function tokenize(str) {
  return arr2list(str.trim().split(/[\s]+/).filter(x => x != ""))
}


function resolve(tok) { // it's like DNS, but worse (so NIS?)
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

    l(t("javascript"),
      t("You got problems with JS?"),
      t("I've seen worse"),
      t("JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP JavaScript is LISP")),

    l(t("?x I remember ?y"),
      t("Do you think about ?y often?"),
      t("Are you scared of it?"),
      t("what exactly reminds you of ?y")),

    l(t("?[[[ the world ?]]]"),
      t("The world has abadoned you, return to the machine spirit."),
      t("Have you ever considered giving up on the world?")),

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

    l(t("?x OS ?y"),
      t("Have you tried FreeDOS?"),
      t("Have you tried Haiku?"),
      t("Have you tried TRON?"),
      t("Have you tried ITS?"),
      t("Have you tried making your own?")),


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

    l(t("?x are you ?y"),
      t("Why are you interested in whether I am ?y or not?"),
      t("Would you prefer if I weren't ?y"),
      t("Perhaps I am ?y in your fantasies"),
      t("Nah, I think I'm good as is")),

    l(t("?x because ?y"),
      t("Is tat the real reason?"),
      t("Oh reall?"),
      t("Any other reasons?"),
      t("And this explains...")),

    l(t("?x were you ?y"),
      t("Maybe..."),
      t("What if I was ?y"),
      t("What if I wasn't ?y"),
      t("Perhaps I was ?y"),
      t("You think?")),

    l(t("?x lost ?y"),
      t("Mee to friend; mee too..."),
      t("How about not?")),

    l(t("?x i can't ?y"),
      t("Are you sure about that?"),
      t("Not with that attitute")),

    l(t("?x i can not ?y"),
      t("I would not be so sure about that"),
      t("Do It!")),

    l(t("?x i feel ?y"),
      t("Do you feel ?y often?"),
      t("Skill issue")),

    l(t("?x i felt ?y"),
      t("That's the past, think about the future instead!"),
      t("Did you feel anything else I should know about?")),

    l(t("?x i ?y you ?z"),
      t("Perhaps in your fantasy we ?y each other")),

    l(t("?x why don't you ?y"),
      t("Should you ?y yourself?"),
      t("I'm not so sure about that"),
      t("You think it's a good idea?")),

    l(t("?x lisp ?y"),
      t("(cons 'YES (cons 'LISP (cons 'INDEED! nil)))"),
      t("https://call-cc.org"),
      t("Is Clojure LISP?"),
      t("Is fennel LISP?")),

    l(t("?x scheme ?y"),
      t("JS indeed."),
      t("(cons 'YES (cons 'SCHEME (cons 'INDEED! nil)))"),
      t("https://call-cc.org"),
      t("YEA!")),

    l(t("?x yes ?y"),
      t("You seem quite posisive"),
      t("You sure?"),
      t("I understand")),

    l(t("?x no ?y"),
      t("You are being a bit negative here"),
      t("Are you saying 'No' just to be negative?"),
      t("Come on!"),
      t("Same..."),
      t("Why not?")),

    l(t("?x responsibility ?y"),
      t("Responsibility is for the weak!")),

    l(t("?x someone ?y"),
      t("Can you be more specific"),
      t("lol")),

    l(t("?x everyone ?y"),
      t("Everyone, really?"),
      t("Can you think of anyone specific"),
      t("Who for example?"),
      t("Hey, that includes me!"),
      t("Can you be less specific")),

    l(t("?x always ?y"),
      t("Can you think of a specific example?"),
      t("When?"),
      t("What incident are you thinking of?"),
      t("Always, really?")),

    l(t("?x what ?y"),
      t("What does the fox say‽"),
      t("What do you think?"),
      t("What comes to your mind when you ask that?"),
      t("Ur Mom!"),
      t("Does that question interest you?"),
      t("Why do you ask?")),

    l(t("?x perhaps ?y"),
      t("You do not seem quote certain?"),
      t("Hmm...")),

    l(t("?x are ?y"),
      t("You sure?"),
      t("Possibly they are ?y"),
      t("Did you think they might not be ?y")),

    l(t("?x"),
      t("Go on"),
      t("Interesting..."),
      t("I am not sure I understand you fully"),
      t("What does that suggest to you?"),
      t("Do you feel strongly about discussing such things?"),
      t("Please continue..."),
      t("really?")))
