## Hands-on exercise #1
1. Create a value and assign it to a variable.
2. Print the address of that value.

**solution:** https://play.golang.org/p/57kW8xd0qT

## Hands-on exercise #2
1. Create a person struct.
2. Create a func called “changeMe” with a *person as a parameter,
    
     - [x] Change the value stored at the *person address.
     
     Important: 
     - to dereference a struct, use (*value).field, for:
         - p1.first 
         - (*p1).first
     -  “As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.” (https://golang.org/ref/spec#Selectors)
3. Create a value of type person and print out the value.
4. In func main call “changeMe” and print out the value.

**solution:** https://play.golang.org/p/AehM662HKS
