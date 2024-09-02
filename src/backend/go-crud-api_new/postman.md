## Postman Examples for API Requests

Here are some Postman examples to interact with the API endpoints of the Go CRUD application.

### 1. Get All Users

- **Method**: `GET`
- **URL**: `http://localhost:8000/users`

#### Postman Request

```http
GET /users HTTP/1.1
Host: localhost:8000
```

- **Response Example**:
    ```json
    [
        {
            "id": 1,
            "name": "John Doe",
            "email": "john@example.com"
        },
        {
            "id": 2,
            "name": "Jane Doe",
            "email": "jane@example.com"
        }
    ]
    ```

### 2. Get a User by ID

- **Method**: `GET`
- **URL**: `http://localhost:8000/users/{id}`

#### Postman Request

```http
GET /users/1 HTTP/1.1
Host: localhost:8000
```

- **Response Example**:
    ```json
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.com"
    }
    ```

### 3. Create a New User

- **Method**: `POST`
- **URL**: `http://localhost:8000/users`
- **Body**: `JSON`

#### Postman Request

```http
POST /users HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "name": "Alice Smith",
    "email": "alice@example.com"
}
```

- **Response Example**:
    ```json
    {
        "id": 3,
        "name": "Alice Smith",
        "email": "alice@example.com"
    }
    ```

### 4. Update an Existing User

- **Method**: `PUT`
- **URL**: `http://localhost:8000/users/{id}`
- **Body**: `JSON`

#### Postman Request

```http
PUT /users/1 HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "name": "Johnathan Doe",
    "email": "johnathan@example.com"
}
```

- **Response Example**:
    ```json
    {
        "id": 1,
        "name": "Johnathan Doe",
        "email": "johnathan@example.com"
    }
    ```

### 5. Delete a User

- **Method**: `DELETE`
- **URL**: `http://localhost:8000/users/{id}`

#### Postman Request

```http
DELETE /users/1 HTTP/1.1
Host: localhost:8000
```

- **Response Example**:
    ```json
    {
        "message": "User deleted"
    }
    ```

### 6. Get All User Roles

- **Method**: `GET`
- **URL**: `http://localhost:8000/roles`

#### Postman Request

```http
GET /roles HTTP/1.1
Host: localhost:8000
```

- **Response Example**:
    ```json
    [
        {
            "id": 1,
            "name": "Admin"
        },
        {
            "id": 2,
            "name": "User"
        }
    ]
    ```

### 7. Create a New User Role

- **Method**: `POST`
- **URL**: `http://localhost:8000/roles`
- **Body**: `JSON`

#### Postman Request

```http
POST /roles HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "name": "Editor"
}
```

- **Response Example**:
    ```json
    {
        "id": 3,
        "name": "Editor"
    }
    ```

### 8. Get All Settings

- **Method**: `GET`
- **URL**: `http://localhost:8000/settings`

#### Postman Request

```http
GET /settings HTTP/1.1
Host: localhost:8000
```

- **Response Example**:
    ```json
    [
        {
            "id": 1,
            "key": "site_name",
            "value": "My Cool Website"
        },
        {
            "id": 2,
            "key": "maintenance_mode",
            "value": "false"
        }
    ]
    ```

### 9. Create a New Setting

- **Method**: `POST`
- **URL**: `http://localhost:8000/settings`
- **Body**: `JSON`

#### Postman Request

```http
POST /settings HTTP/1.1
Host: localhost:8000
Content-Type: application/json

{
    "key": "new_feature_enabled",
    "value": "true"
}
```

- **Response Example**:
    ```json
    {
        "id": 3,
        "key": "new_feature_enabled",
        "value": "true"
    }
    ```

```