protocol Greet {
  
  // blueprint of property
  var name: String { get }

  // blueprint of a method 
  func message() 
} 

// conform class to Greet protocol
class Employee: Greet {

  // implementation of property
  var name = "Perry"

  // implementation of method
  func message() {
    print("Good Morning", name)
  }
}

var employee1 = Employee()
employee1.message()