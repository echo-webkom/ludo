# Ludo API documentation

## Endpoints

### `GET /items` - All items

Get a list of all items.

Response type:

```json
[
    {
        "id": int
        "title": string
        "description": string
        "connectedBranch": string
        "list": "backlog" | "todo" | "in progress" | "in review"
        "creator": {
            "name": string
        }
        "assignee": {
            "name": string
        }
        "repo": {
            "name": string
            "url": string
        }
    }
]
```

### `POST /item` - Create item

Create new item (todo or backlog).

Expects the following format as the request body, all fields are required:

```json
{
    "title": string
    "description": string // Ignored for backlog items
    "list": "backlog" | "todo"
}
```

Responds with `200` on success.

### `PATCH /item/{id}` - Update item

Todo

### `GET /users` - All users

Get a list of all users.

Response type:

```json
[
	{
        "id": int
		"username": string
	}
]
```

Responds with `200` on success.

### `DELETE /user/{id}` - Delete user

Delete a user by its ID.

Responds with `200` on success.
