# CRUD-GoLang-Project

# SHOPPING CART MANAGEMENT
This project is a simple REST API written in Go using the gofr framework, facilitating CRUD operations for a shopping cart. The API follows Clean Architecture and SOLID principles.
This is a project for ZopSmart Technologies on GoLang using GoFr.


![CRUD](https://github.com/sriayush02/CRUD-GoLang-Project/assets/83555352/cf36e77f-98a6-4615-9e28-ef780203566d)






# Features
1.Create Cart Item:
> POST /cart/add/{name}/{amount}
> Create a new item in the shopping cart with the provided name and amount.

2.View Cart:
> GET /cart/view
> Retrieve the list of items in the shopping cart.

3.Delete Cart Item:
> DELETE /cart/delete/{id}
> Delete a specific item from the shopping cart using its ID.

4.Update Cart Item:
> PUT /cart/update/{id}
> Update the details of a specific item in the shopping cart.



![uml](https://github.com/sriayush02/CRUD-GoLang-Project/assets/83555352/b7f5a98d-597f-4136-b524-48da2c6394fa)


# API Endpoints

1. Create Cart Item:
   POST /cart/add/{name}/{amount}

2. View Cart:
   GET /cart/view

3. Delete Cart Item:
   DELETE /cart/delete/{id}

4. Update Cart Item:
   PUT /cart/update/{id}

![image](https://github.com/sriayush02/CRUD-GoLang-Project/assets/83555352/c46a3ded-406a-443c-8cd9-71b685c57d26)


# How to RUN 

1. Clone the repository:
   git clone [https://github.com/sriayush02/CRUD-GoLang-Project]

2. Install dependencies:
   go mod tidy

3. Run the application:
   go run main.go


