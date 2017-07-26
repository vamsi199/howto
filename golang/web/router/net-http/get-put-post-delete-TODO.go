// data exchange format: json
// just using net/http package
// persist data just in memory
// assume below data struct (use slice of this struct to store multiple records)
  // type customer struct{Id int, Name string, Phone string, City string, Gender string}
// below endpoints
  // GET
    // /customer : get all customers
    // /customer/1 : get customer with id=1
    // /customer?name=abc : get all customers with name = abc
    // /customer?city=hyderabad : get all customers with city = hyderabad, likewise search by phonenumber, or gender
    // /customer?city=hyderabad&gender=male : like wise in any combination
  // PUT
    // /customer/1 : modify the customer 1 record with the new data given in the body of the request
  // DELETE
    // /customer/1 : delete customer 1 record
  // POST
    // /customer : add new customer record with the given data in the body of the request
