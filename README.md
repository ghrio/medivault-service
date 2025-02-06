# Medivault

Medivault is a secure, RESTful microservice for storing and managing patient data—such as doctor posts, medical issues, and more. Built with Go's standard `net/http` package and using sqlc for type-safe SQL query generation, Medivault is designed with clean code practices, robust security measures, and compliance in mind.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Overview

Medivault provides a REST API that allows you to perform CRUD (Create, Read, Update, Delete) operations on sensitive patient data. The service emphasizes clear, resource-oriented endpoint naming, proper security practices, and separation of concerns for maintainability. This project is especially designed for healthcare applications where regulatory compliance (e.g. HIPAA, GDPR) is essential.

## Features

- **RESTful API Design:** Intuitive endpoints using plural nouns and versioning (e.g., `/v1/patients`).
- **Type-Safe Database Access:** sqlc generates efficient, type-safe Go code from your SQL queries.
- **Security Focused:** Includes HTTPS (when deployed), input validation, and plans for JWT authentication.
- **Compliance Ready:** Designed with privacy and regulatory guidelines in mind.
- **Separation of Concerns:** Clear separation between HTTP handling, business logic, and data access.

## Tech Stack

- **Language:** Go
- **Web Framework:** Standard `net/http`
- **Database Access:** sqlc (SQL query code generation)
- **Database:** PostgreSQL (or another SQL database of your choice)
- **Other Tools:**
  - Git for version control
  - Go's built-in testing framework for unit and integration tests

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/medivault.git
   cd medivault
   ```
