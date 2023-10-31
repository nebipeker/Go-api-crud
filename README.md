
# Simple Go API with Postman Collection

This repository contains a simple Go (Golang) API built using the Echo framework and GORM for interacting with a PostgreSQL database. Additionally, it includes a Postman collection for testing the API endpoints.

## API Features

- **Candidate Management**:
  - Create, retrieve, update, and delete candidate profiles.
  
- **Company Management**:
  - Create, retrieve, update, and delete company profiles.

- **Job Listing Management**:
  - Create, retrieve, update, and delete job listings associated with companies.

## Getting Started

To run this API and test its functionality using Docker and Docker Compose, follow these steps:

1. **Clone the Repository**:
   ```
   git clone https://github.com/nebipeker/Go-api-crud.git
   cd Go-api-crud
   ```

2. **Create a `.env` file**:
   - Create a `.env` file in the project directory to configure environment variables used by Docker Compose. You can use the following example as a starting point:
   ```
   POSTGRES_USER=your_db_user
   POSTGRES_PASSWORD=your_db_password
   POSTGRES_DB=jobkeep
   ```

3. **Build and Start the Docker Containers**:
   ```
   docker-compose up --build
   ```

4. **Import the Postman Collection**:
   - Import the provided Postman collection to test the API endpoints located in the `postman` directory.

5. **Start Testing**:
   - Use Postman to send requests to the API endpoints for candidate, company, and job listing management.

## API Routes

The API exposes the following routes:

- Candidate Management:
  - `/api/candidates`: CRUD operations for candidates.
  
- Company Management:
  - `/api/companies`: CRUD operations for companies.

- Job Listing Management:
  - `/api/jobs`: CRUD operations for job listings.

## Postman Collection

The Postman collection, located in the `postman` directory, contains pre-defined requests for testing the API. You can import this collection into Postman to streamline your testing process.

## Configuration

The API configuration is managed through environment variables set in the `.env` file. Ensure you update this file with your specific configuration.

## Dependencies

- Echo: A fast and minimalist Go web framework.
- GORM: An Object Relational Mapping (ORM) library for Go.
- Postman: A popular tool for testing APIs.
- Postgres: An open-source relational database.

## TODO List

- Parametrization of hard and soft delete.
- Building CI/CD pipeline.
- Adding tests.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to contribute to this project or use it as a reference for your own Go API development. If you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.
