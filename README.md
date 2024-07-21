# LruCache backend

## How TO Run and Install

**Clone the repository:**

```sh
  https://github.com/harshsri28/lruCache-Backend
```

**Setup The Environment**

```sh
PORT = "3001"
CACHE_CAPACITY = 10
```

**Install the Dependency**

```sh
go mod tidy
```

**Run the Application**

```sh
go run main.go
```

## Overview

This Project facilates LRU(Least Recently Functionality) and had four api exposed-

### 1. Get All Cache Items

- **Endpoint**: `/cache`
- **Method**: `GET`
- **Response**:
  ```json
  {
    "items": {
      "apple": "109",
      "mango": "109"
    }
  }
  ```

### 2. Get Specific Item using Key

- **Endpoint**: `/cache/mango` // here {key} = mango
- **Method**: `GET`
- **Response**:

  ```json
  {
    "key": "mango",
    "value": "109"
  }
  ```

### 3. Create an Item in Cache

- **Endpoint**: `/cache`
- **Method**: `POST`
- **Request Body**:

  ```json
  {
    "key": "mango",
    "value": "109",
    "duration": 50
  }
  ```

- **Response**:

  ```json
  {
    "status": "success"
  }
  ```

### 4. Delete an Item in Cache

- **Endpoint**: `/cache/{key}`
- **Method**: `POST`

- **Response**:

  ```json
  {
    "status": "success"
  }
  ```

  ## Folder Structure

- controllers
  - cacheControllers.go
- module
  - lruCache.go
- routes
  - cacheRoutes.go
- main.go
