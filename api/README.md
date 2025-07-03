# Ludo API documentation

Ludo is comprised of users, boards, and items. Each board has a list of users, and a list of items.

## Endpoints

### Users

Users are not accounts. You are not _logged in_ as your user. Users are simply used to mark ownership of items. Each user is linked to a GitHub account.

| Method | Path              | Response Type | Description             |
| ------ | ----------------- | ------------- | ----------------------- |
| GET    | `/users`          | `User[]`      | Get all users           |
| POST   | `/users`          | `ID`          | Create new user         |
| GET    | `/users/{userId}` | `User`        | Get user by ID          |
| PATCH  | `/users/{userId}` | -             | Update an existing user |
| DELETE | `/users/{userId}` | -             | Delete a user           |

### Boards

Boards are containers for items and have users assosiated with them. A board may be connected to a GitHub repo.

| Method | Path                                      | Response Type | Description                                |
| ------ | ----------------------------------------- | ------------- | ------------------------------------------ |
| GET    | `/boards`                                 | `Board[]`     | Get all boards                             |
| POST   | `/boards`                                 | `ID`          | Create new board                           |
| GET    | `/boards/{boardId}`                       | `Board`       | Get board by ID                            |
| PATCH  | `/boards/{boardId}`                       | -             | Update an existing board                   |
| DELETE | `/boards/{boardId}`                       | -             | Delete a board                             |
| GET    | `/boards/{boardId}/users`                 | `[]User`      | Get all users in board                     |
| GET    | `/boards/{boardId}/items`                 | `[]Item`      | Get all items in board                     |
| POST   | `/boards/{boardId}/users/{userId}`        | -             | Add user to board                          |
| DELETE | `/boards/{boardId}/users/{userId}`        | -             | Remove user from board                     |
| GET    | `/boards/{boardId}/status/{status}/items` | `[]Item`      | Get all items in board with given `Status` |

### Items

| Method | Path                   | Response Type | Description             |
| ------ | ---------------------- | ------------- | ----------------------- |
| GET    | `/items`               | `[]Item`      | Get all items           |
| POST   | `/items`               | `ID`          | Create an item          |
| GET    | `/items/{itemId}`      | `Item`        | Get item by ID          |
| PATCH  | `/items/{itemId}`      | -             | Update an existing item |
| DELETE | `/items/{itemId}`      | -             | Delete an item          |
| GET    | `/items/{itemId}/data` | `any`         | Get items data field    |
| PATCH  | `/items/{itemId}/data` | -             | Update items data field |

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

| Field       | Type   | Description                                                          |
| ----------- | ------ | -------------------------------------------------------------------- |
| id          | uint   | Unique ID                                                            |
| status      | uint   | Item status. See `Status` enum.                                      |
| title       | string | Items title                                                          |
| description | string | Items description                                                    |
| creator     | uint   | ID of user that created this item                                    |
| assignee    | uint   | ID of user this item is assigned to                                  |
| branch      | string | Name of branch this item is tracking                                 |
| createdAt   | uint   | Unix time of creation                                                |
| updatedAt   | uint   | Unix time of last update                                             |
| data        | string | Additional data field. Usually for board-specific item data as JSON. |

### ID

When creating new objects you get its ID in response.

| Field | Type | Description      |
| ----- | ---- | ---------------- |
| id    | uint | ID of new object |

### Status

Item status is one of the following values. They are automatically updated if the item contains github repo/branch info. Use the integer value as the status in requests.

| Name             | Value | Description                                              |
| ---------------- | ----- | -------------------------------------------------------- |
| StatusBacklog    | 0     | Item has been created with little to no details          |
| StatusReady      | 1     | Item is ready to be worked on and has sufficient details |
| StatusInProgress | 2     | Item is being worked on and tracks a branch              |
| StatusInReview   | 3     | Item has a pull request open                             |
| StatusClosed     | 4     | Items pull request was merged or closed                  |
