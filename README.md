# Role-Based Access Control (RBAC)

 Role-Based Access Control (RBAC) is a system designed to manage access to resources within an application based on user roles. It provides a robust framework for implementing secure Authentication and Authorization, ensuring that users can access only the resources and perform actions allowed for their assigned roles.

The project leverages flexible role and permission structures to support granular control over user actions, enabling organizations to define custom roles (e.g., Admin, Moderator, User) and assign permissions tailored to their needs. This ensures that sensitive operations are restricted to authorized users while maintaining scalability and ease of use.

Additionally, the RBAC system is designed to seamlessly integrate with modern application architectures, ensuring that it can be incorporated into APIs, web applications, or microservices with minimal overhead. By implementing industry best practices, such as secure session management (e.g., JWT) and middleware for route protection, the system ensures data integrity and operational security.

This approach not only strengthens the application’s security posture but also simplifies user management, making it a vital component for any system requiring controlled access to resources.
<br><br>

<!--TABLE OF CONTENTS-->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a> 
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a> 
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
  </details>

<!--About the Project-->
  
## About The Project
<!--
### Demo
-->
<br>

### Key Concepts
1. **Authentication and Authorization** : Authentication verifies the identity of a user, ensuring only valid users can access the system. Authorization determines what actions or resources an authenticated user is allowed to access based on their role.
3. **Middleware for Route Protection** : Middleware is used to protect routes by checking user roles and permissions before granting access to a specific endpoint. This ensures unauthorized requests are blocked at an early stage.
4. **Granular Permissions** : Granular permissions allow fine-tuned control over which actions a role can perform. For instance, an Admin might have full control, while a User can only read or create specific resources.
5. **Scalable Architecture** : The RBAC system is designed to integrate seamlessly with APIs, web applications, and microservices, ensuring scalability and maintainability as the application grows.


### Built With
<br><br>

<img height="100px" src="https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg"/>



<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!--GETTING STARTED-->

## Getting Started

To get started with your Golang application, follow these steps:

1. **Install Golang**: Download and install Golang from the [official website](https://golang.org/dl/).

2. **Set Up Your Workspace**: Create a new directory for your project and set your `GOPATH` environment variable to point to this directory.

3. **Initialize Your Project**: Inside your project directory, run the following command to initialize a new Go module:

   ```
   go mod init github.com/your-username/project-name
   ```
   After installing Golang, you can start running your Go project.
4. **Run without Debugging**: In your terminal, navigate to the directory containing your main Go file (usually named `main.go`). Then, run the following command to build and execute your Go application:
   ```
   go run main.go
   ```
   This command will compile and execute your Go program without generating a binary file.



## Installation 

1. Create an image from the docker file:
   
   ```
   docker build -t rbac .
   ```
3. Run this on your terminal (needs docker to be preinstalled):
   
   ```
   docker run -p 3000:3000 -it rbac
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Routes

### Un-Authorized Routes

- **POST "/signup"**
  * Route to create a new User/Moderator/Admin.
  * Request:
    ```
    {
    "first_name":"Test1",
    "last_name":"User1",
    "email":"testuser1@gmail.com",
    "password":"123456",
    "type":"User"/"Admin"/"Moderator"
    }
    ```
  * Response:
    ```
    {
    "data": {
      "InsertedID": "",
      "user": {
        "ID": "",
        "first_name": "",
        "last_name": "",
        "password": "", // this is always null or "" in the response
        "email": "",
        "token": "",
        "refresh_token": "",
        "created_at": "",
        "updated_at": "",
        "type": ""
      }
    },
    "success": true
    }
    ```
      
- **Get "/login"**
  * Request:
    ```
    {
    "email": <email>,
    "password": <password>
    }
    ```
  * Response:
    ```
    {
    "data": {
      "InsertedID": "",
      "user": {
        "ID": "",
        "first_name": "",
        "last_name": "",
        "password": "", // this is always null or "" in the response
        "email": "",
        "token": "",
        "refresh_token": "",
        "created_at": "",
        "updated_at": "",
        "type": ""
      }
    },
    "success": true
    }
    ```

### Authorized Routes
      
- **GET "/users"**
  * This action can be executed by an `ADMIN` or a `Moderator` or a `User` type.
  * Header:
    ```
    {
      "Content-Type":"application/json",
      "Authorization":"Bearer <token>",
    }
    ```
  * Response as:
      - Sucess : True/False
      - Message: Error message if present
      - Data: Respective data
    ```
    {
      "data": {
        "users":[
          {
            "id": <id>,
            "email": "testuser3@gmail.com",
            "first_name": "Test3",
            "last_name": "User3",
            "created_at": "0001-01-01T00:00:00Z",
            "updated_at": "0001-01-01T00:00:00Z",
            "type": "User"
          }
        ],
      },
      "success": true
    }
    ```
- **GET "/users/:id"**
  * This action can be executed by an `ADMIN` or a `Moderator` or a `User` type.
  * Header:
    ```
    {
      "Content-Type":"application/json",
      "Authorization":"Bearer <token>",
    }
    ```
  * Response as:
      - Sucess : True/False
      - Error: Error message if present
      - Data: Respective data
    ```
    {
    "data": {
      "user": {
        "id": "6744734191c2ef7cdec50709",
        "email": "test@gmail.com",
        "first_name": "Test",
        "last_name": "Test",
        "created_at": "0001-01-01T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "type": "Admin"
      }
    },
    "success": true
    }
    ```

- **PATCH "/update/:id"**
  * This action can be executed by an `ADMIN` or a `Moderator` user.
  * Request(add the attributes to be changed):
    ```
    {
    "last_name":"Test2"
    }
    ```
  * Response:
    ```
    {
      "message": "User updated successfully",
      "success": true
    }
    ```
- **DELETE "/delete/:id"**
  * This action can be executed by an `ADMIN` user.

### Frontend Routes
- **"/home"** - Home page of the application.
- **"/register"** - Register page for the application.
- **"/loginPage"** - Login page for the application.
  
## Screenshots:

<br>
<center>
<img width="1000" src="https://github.com/user-attachments/assets/22a6033c-bf70-4e4c-a4cb-0f27d67c58d6"></img>
<br>
<img width="1000" src="https://github.com/user-attachments/assets/03d9c1cb-1f66-42f7-9207-75212ba4ba93"></img>
 <br>
<img width="1000" src="https://github.com/user-attachments/assets/cb366fc7-f99f-43d6-88c0-78fbbae24aec"></img>
</center>
<br>


<!--CONTRIBUTING-->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire ,and create.Any contributions you make are *greatly appreciated*.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Uttkarsh Raj - https://github.com/Uttkarsh-raj <br>

<p align="right">(<a href="#readme-top">back to top</a>)</p>
