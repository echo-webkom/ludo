# Ludo API documentation

## Endpoints

### Users

| Method | Path              | Response Type | Description             |
| ------ | ----------------- | ------------- | ----------------------- |
| GET    | `/users`          | `User[]`      | Get all users           |
| POST   | `/users`          | -             | Create new user         |
| GET    | `/users/{userId}` | `User`        | Get user by ID          |
| PATCH  | `/users/{userId}` | -             | Update an existing user |
| DELETE | `/users/{userId}` | -             | Delete a user           |

### Boards

| Method | Path                                     | Response Type | Description                 |
| ------ | ---------------------------------------- | ------------- | --------------------------- |
| GET    | `/boards`                                | `Board[]`     | Get all boards              |
| POST   | `/boards`                                | -             | Create new board            |
| GET    | `/boards/{boardId}`                      | `Board`       | Get board by ID             |
| PATCH  | `/boards/{boardId}`                      | -             | Update an existing board    |
| DELETE | `/boards/{boardId}`                      | -             | Delete a board              |
| GET    | `/boards/{boardId}/users`                | `[]User`      | Get all users in board      |
| GET    | `/boards/{boardId}/items`                | `[]Item`      | Get all items in board      |
| GET    | `/boards/{boardId}/lists/{listId}/items` | `[]Item`      | Get all items in board list |

### Items

| Method | Path                            | Response Type | Description                  |
| ------ | ------------------------------- | ------------- | ---------------------------- |
| GET    | `/items`                        | `[]Item`      | Get all items                |
| POST   | `/items`                        | -             | Create an item               |
| GET    | `/items/{itemId}`               | `Item`        | Get item by ID               |
| PATCH  | `/items/{itemId}`               | -             | Update an existing item      |
| DELETE | `/items/{itemId}`               | -             | Delete an item               |
| PATCH  | `/items/{itemId}/move/{listId}` | -             | Move an item to another list |

## Schemas

All field names are identical to the ones found in the JSON response data.

### User

| Field          | Type   | Description                                |
| -------------- | ------ | ------------------------------------------ |
| id             | uint   | Unique ID                                  |
| displayName    | string | Name to display in boards                  |
| githubUsername | string | Username on GitHub, used to track branches |
| createdAt      | uint   | Unix time of creation                      |
| updatedAt      | uint   | Unix time of last update                   |

### Board

| Field     | Type   | Description               |
| --------- | ------ | ------------------------- |
| id        | uint   | Unique ID                 |
| repoName  | string | Name of remote repository |
| repoUrl   | string | URL to remote repository  |
| createdAt | uint   | Unix time of creation     |
| updatedAt | uint   | Unix time of last update  |

### Item

| Field       | Type   | Description                          |
| ----------- | ------ | ------------------------------------ |
| id          | uint   | Unique ID                            |
| list        | uint   | ID of list this item is in           |
| title       | string | Items title                          |
| description | string | Items description                    |
| creator     | uint   | ID of user that created this item    |
| assignee    | uint   | ID of user this item is assigned to  |
| branch      | string | Name of branch this item is tracking |
| createdAt   | uint   | Unix time of creation                |
| updatedAt   | uint   | Unix time of last update             |
