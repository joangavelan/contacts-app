# Todo Application

This is a web-based todo application built using Golang, HTMX, and SQLite. It provides users with a simple yet efficient interface to manage their tasks. The application is developed primarily for learning purposes in backend development. It utilizes JWT authentication for secure access and performs CRUD operations to interact with the SQLite database.

## Features

- **JWT Authentication:** Secure user authentication using JSON Web Tokens (JWT) ensures that only authorized users can access the application's functionalities.
- **CRUD Operations:** Users can create, read, update, and delete tasks, allowing them full control over their todo lists.
- **SQLite Database:** The application utilizes SQLite as the database management system to store and manage todo tasks efficiently.
- **Dynamic User Interface:** HTMX enables dynamic and interactive user interfaces, enhancing user experience without extensive JavaScript coding.

## Roadmap

- [ ] Implement JWT authentication middleware for user authentication.
- [ ] Set up basic CRUD API endpoints for managing todo tasks.
- [ ] Design and implement the database schema for storing todo tasks in SQLite.
- [ ] Create basic HTML templates for rendering todo lists and task forms.
- [ ] Enhance user interface with dynamic updates using HTMX for seamless task management.
- [ ] Add support for task categories and tags for better organization.

## Technologies Used

- **Golang:** Backend logic and server-side operations are implemented using the Go programming language.
- **HTMX:** HTMX simplifies frontend development by allowing HTML to be the primary language for defining dynamic UI interactions.
- **SQLite:** Lightweight and efficient SQLite serve as the database management system for storing todo tasks.
- **JWT:** JSON Web Tokens provide secure authentication and authorization mechanisms for user access control.

## Setup Instructions

1. **Clone the Repository:**
   ```
   git clone https://github.com/joangavelan/todo-app.git
   ```

2. **Navigate to the Project Directory:**
   ```
   cd todo-app
   ```
3. **Install npm dependencies:**
   ```
   npm install
   ```

4. **Run the server:**
   ```
   go run cmd/main.go
   ```

5. **Access the Application:**
   Open your web browser and navigate to `http://localhost:3000` to access the todo application.

## Usage

- **Authentication:** Users need to sign in using their credentials to access the todo application. Upon successful authentication, a JWT token is generated and used for subsequent requests.
- **CRUD Operations:** Users can perform CRUD operations on their todo tasks, including adding, updating, deleting, and marking tasks as completed.
- **Database Interaction:** All user tasks are stored and managed using the SQLite database, ensuring data integrity and persistence.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
