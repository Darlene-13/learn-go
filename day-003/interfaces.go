package main

// ============================================================
// GO INTERFACES — NOTES
// ============================================================

// 1. INTERFACES
//
// An interface is a collection of method signatures.
//
// A type satisfies an interface automatically by implementing
// all the methods required by that interface.
//
// Go interfaces are implemented implicitly.
// Unlike Java, we do NOT write:
//
//	implements expense
//
// Example:
//type expense interface {
//cost() float64
//}

//
// Any type with:
//     cost() float64
//
// automatically satisfies expense.

// 2. ONE TYPE CAN IMPLEMENT MULTIPLE INTERFACES
//
// A concrete type can satisfy multiple interfaces as long as
// it has all the required methods.
//
// email has:
//
//     func (e email) cost() float64
//     func (e email) print()
//
// Therefore email satisfies BOTH:
//
//     expense
//     printer
//
// This is why we can do:
//
//     e := email{...}
//     getEmail(e, e)
//
// The SAME email value is passed to both parameters.
//
// getEmail expects:
//
//     func getEmail(e expense, p printer)
//
// The first e must satisfy expense.
// The second e must satisfy printer.
//
// The parameter names e and p do NOT determine the interface.
// Their TYPES do.
//
// We could write:
//
//     func getEmail(exp expense, printer printer)
//
// and it would work exactly the same way.

// 3. := VS =
//
// := declares a NEW variable and assigns a value to it.
//
// Example:
//
//     e := email{...}
//
// This means:
//     "Create a variable called e and store an email in it."
//
// Once e already exists, use = to replace its value:
//
//     e = email{...}
//
// Therefore:
//
//     e := email{...}  // create e
//     e = email{...}   // replace e
//     e = email{...}   // replace e again
//
// You cannot use := again in the same scope just to replace e:
//
//     e := email{...}
//     e := email{...}  // ERROR

// 4. PASSING STRUCT VALUES DIRECTLY
//
// We don't always need to create a variable.
//
// This:
//
//     e := email{...}
//     getEmail(e, e)
//
// can also be written as:
//
//     getEmail(email{...}, email{...})
//
// But these are technically TWO separate email values.
//
// Using:
//
//     e := email{...}
//     getEmail(e, e)
//
// means we create ONE email value and pass the SAME value twice.
//
// We usually create a variable when we want to reuse the value,
// change it later, or make the code easier to read.
//
//
// 5. TYPE ASSERTIONS
//
// A type assertion lets us check/recover the concrete type
// stored inside an interface.
//
// Syntax:
//
//     value, ok := interfaceValue.(ConcreteType)
//
// Example:
//
//     em, ok := e.(email)
//
// This asks:
//
//     "Is the concrete value inside e actually an email?"
//
// If YES:
//
//     ok == true
//     em contains the email value
//
// If NO:
//
//     ok == false
//     em is the zero value of email
//
// Example:
//
//     func getExpenseReport(e expense) (string, float64) {
//
//         em, ok := e.(email)
//         if ok {
//             return em.toAddress, em.cost()
//         }
//
//         sm, ok := e.(sms)
//         if ok {
//             return sm.toPhoneNumber, sm.cost()
//         }
//
//         return "", 0.0
//     }
//
// The important mental model:
//
//     interface value
//          ↓
//     type assertion
//          ↓
//     concrete value
//
//     expense → email
//     expense → sms

// 6. COMMA-OK TYPE ASSERTION
//
// This:
//
//     em, ok := e.(email)
//
// is called the "comma-ok" form.
//
// ok tells us whether the assertion succeeded.
//
//     true  → e contains an email
//     false → e does not contain an email
//
// This is safer than assuming the assertion will succeed.

// 7. TYPE SWITCH
//
// A type switch lets us check multiple possible concrete types
// stored inside an interface.
//
// Example:
//
//     switch e.(type) {
//     case email:
//         ...
//     case sms:
//         ...
//     default:
//         ...
//     }
//
// This is useful when different concrete types need different
// behavior.
//
// Mental model:
//
//     interface
//        ↓
//     "What concrete type are you?"
//        ↓
//     email / sms / something else

// 8. TYPE ASSERTION VS TYPE SWITCH
//
// Type assertion:
//
//     em, ok := e.(email)
//
// asks about ONE specific type.
//
// Type switch:
//
//     switch e.(type) {
//     case email:
//     case sms:
//     }
//
// checks MULTIPLE possible types.
//
// Both are ways of inspecting the concrete type stored inside
// an interface.

// 9. IMPORTANT: INTERFACES ARE ABOUT BEHAVIOR
//
// email does not need to know that it implements expense.
//
// Go simply checks whether email has the required methods.
//
// If email has:
//
//     cost() float64
//
// it satisfies expense.
//
// If email has:
//
//     print()
//
// it satisfies printer.
//
// If another type has cost() float64 but does NOT have print(),
// it can still satisfy expense.
//
// It does NOT need to implement printer.
//
// A type can therefore satisfy one interface, several interfaces,
// or many interfaces depending on the methods it provides.
