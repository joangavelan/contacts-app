# Contacts Application

This is a hypermedia-driven contacts application inspired by the sample application from the
[Hypermedia Systems](https://hypermedia.systems/) book. Built using Golang, HTMX, and SQLite, it provides users with a simple yet efficient interface to manage their contacts. The application is developed primarily for learning purposes in backend development, with a focus on building hypermedia-driven applications.

## Features

- **JWT Authentication:** Secure user authentication using JSON Web Tokens (JWT) ensures that only authorized users can access the application's functionalities.
- **CRUD Operations:** Users can create, read, update, and delete contacts, allowing them full control over their contact lists.
- **Search, Filtering, Pagination, and Ordering**: Users can search for contacts, apply filters, paginate through contact lists, and order contacts based on various criteria for better organization.
- **Upload/Download Contacts:** Users can upload and download their contact lists using CSV or Excel files.

## Roadmap

- [x] Design and implement the database schema for storing contacts in SQLite.
- [x] Implement JWT authentication.
- [ ] Create user interfaces for adding, updating, and viewing contacts.
- [ ] Set up API endpoints for managing contacts.
- [ ] Implement search, filtering, pagination, and ordering functionalities for efficient contact management.
- [ ] Add support for bulk uploading and downloading of contacts using CSV or Excel files .

## Technologies Used

- **Golang:** Backend logic and server-side operations are implemented using the Go programming language.
- **HTMX:** HTMX simplifies frontend development by allowing HTML to be the primary language for defining dynamic UI interactions.
- **SQLite:** Lightweight and efficient SQLite serve as the database management system for storing contacts.
- **JWT:** JSON Web Tokens provide secure authentication and authorization mechanisms for user access control.

## Setup Instructions

1. **Clone the Repository:**

   ```
   git clone https://github.com/joangavelan/contacts-app.git
   ```

2. **Navigate to the Project Directory:**
   ```
   cd contacts-app
   ```
3. **Install npm dependencies:**

   ```
   cd web && npm install
   ```

4. **Run the server:**

   ```
   go run cmd/main.go
   ```

   or

   ```
   air
   ```

5. **Access the Application:**
   Open your web browser and navigate to `http://localhost:3000` to access the application.

## Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
