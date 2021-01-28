### Register

**Method**: `POST`

**Content-Type**: `multipart/form-data`

**URL**: `/api/register`

**Params**:

| param    | Required | Type   | Description |
| :------- | :------- | :----- | ----------- |
| username | Y        | string | username    |
| email    | Y        | string | email       |
| password | Y        | string | password    |

**Response:**

success:

```json
{
    "error_code": 0,
    "message": "创建成功"
}
```

fail:

```json
{
    "error": "Expected arguments.",
    "error_code": 1,
    "message": "Fail to register."
}
```



