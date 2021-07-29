# API 文档

## 公共字段

- 请求体

  | Field         | Type     | Description |
  | ------------- | -------- | ----------- |
  | `accessToken` | `string` | jwt令牌     |
  | `params`      | `object` | 接口参数    |

- 响应体

  | Field  | Type     | Description                       |
  | ------ | -------- | --------------------------------- |
  | `code` | `number` | 0为成功，1为接口错误，2为令牌错误 |
  | `msg`  | `string` | 响应消息，无错误时为空            |
  | `data` | `object` | 响应数据，有错误时为nil           |

- 错误信息
  
  | Error    | Message                      |
  | -------- | ---------------------------- |
  | 令牌错误 | 令牌错误或已过期，请重新登录 |

## 管理后台

### 公有接口（无鉴权）

#### 登录

```http
POST /api/system/auth/login
```

- 请求
  
  | Field      | Type     | Description | Require |
  | :--------- | :------- | :---------- | ------- |
  | `account`  | `string` | 账号        | `true`  |
  | `password` | `string` | raw密码     | `true`  |

- 响应

  | Field          | Type     | Description              |
  | -------------- | -------- | ------------------------ |
  | `access_token` | `string` | jwt令牌                  |
  | `expiry_time`  | `number` | 令牌过期时间，unix时间戳 |
  | `user`         | `object` | 账号信息（待定）         |

- 错误信息
  
  | Error          | Message            |
  | -------------- | ------------------ |
  | 账号或密码为空 | 账号或密码不能为空 |
  | 密码错误       | 密码错误           |

### 私有接口（有鉴权）

#### 获取文章列表

```http
POST /api/system/article/listArticles
```

- 请求

  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应

  | Field        | Type        | Description |
  | ------------ | ----------- | ----------- |
  | `articles`   | `article[]` | 文章列表    |
  | `totalCount` | `number`    | 文章总数    |

- article
  
  | Field            | Type       | Description                               |
  | ---------------- | ---------- | ----------------------------------------- |
  | `articleId`      | `number`   | 文章ID                                    |
  | `title`          | `string`   | 文章标题                                  |
  | `state`          | `number`   | 文章状态（0为草稿，1为已发布，2为已删除） |
  | `views`          | `number`   | 阅读量                                    |
  | `effectiveViews` | `number`   | 有效阅读量（待讨论）                      |
  | `tags`           | `string[]` | 标签列表                                  |
  | `createdTime`    | `number`   | 文章创建时间                              |
  | `updatedTime`    | `number`   | 文章最后更新时间                          |
  | `deletedTime`    | `number`   | 文章删除时间                              |

#### 获取文章详情

```http
POST /api/system/article/getArticleDetails
```

- 请求

  | Field       | Type     | Description | Require |
  | :---------- | :------- | :---------- | ------- |
  | `articleId` | `number` | 文章ID      | `true`  |

- 响应

  | Field            | Type               | Description |
  | ---------------- | ------------------ | ----------- |
  | `articleDetails` | `articleDetails{}` | 文章详情    |

- articleDetails

  | Field            | Type           | Description                               |
  | ---------------- | -------------- | ----------------------------------------- |
  | `articleId`      | `number`       | 文章ID                                    |
  | `title`          | `string`       | 文章标题                                  |
  | `desc`           | `string`       | 文章描述                                  |
  | `content`        | `string`       | 文章内容（markdown）                      |
  | `state`          | `number`       | 文章状态（0为草稿，1为已发布，2为已删除） |
  | `coverImage`     | `coverImage{}` | 文章头图                                  |
  | `views`          | `number`       | 阅读量                                    |
  | `effectiveViews` | `number`       | 有效阅读量（待讨论）                      |
  | `tags`           | `string[]`     | 标签列表                                  |
  | `createdTime`    | `number`       | 文章创建时间                              |
  | `updatedTime`    | `number`       | 文章最后更新时间                          |
  | `deletedTime`    | `number`       | 文章删除时间                              |

- coverImage

  | Field      | Type     | Description          |
  | ---------- | -------- | -------------------- |
  | `url`      | `string` | 头图url              |
  | `mime`     | `string` | image/jpg, image/png |
  | `filename` | `string` | 文件名               |

#### 搜索文章

```http
POST /api/system/article/searchArticle
```

- 请求
  
  | Field     | Type     | Description                      | Require |
  | --------- | -------- | -------------------------------- | ------- |
  | `keyword` | `string` | 搜索关键字（接受标题和文章内容） | `true`  |

- 响应

  | Field      | Type        | Description |
  | ---------- | ----------- | ----------- |
  | `articles` | `article[]` | 文章列表    |

- article
  
  | Field            | Type       | Description                               |
  | ---------------- | ---------- | ----------------------------------------- |
  | `articleId`      | `number`   | 文章ID                                    |
  | `title`          | `string`   | 文章标题                                  |
  | `state`          | `number`   | 文章状态（0为草稿，1为已发布，2为已删除） |
  | `views`          | `number`   | 阅读量                                    |
  | `effectiveViews` | `number`   | 有效阅读量（待讨论）                      |
  | `tags`           | `string[]` | 标签列表                                  |
  | `createdTime`    | `number`   | 文章创建时间                              |
  | `updatedTime`    | `number`   | 文章最后更新时间                          |
  | `deletedTime`    | `number`   | 文章删除时间                              |

#### 保存文章

```http
POST /api/system/article/saveArticle
```

- 请求

  | Field        | Type           | Description | Require |
  | ------------ | -------------- | ----------- | ------- |
  | `title`      | `string`       | 文章标题    | `true`  |
  | `coverImage` | `coverImage{}` | 文章头图    | `false` |
  | `desc`       | `string`       | 文章描述    | `false` |
  | `content`    | `string`       | 文章内容    | `false` |

- 无响应
  
- coverImage

  | Field      | Type     | Description               | Require |
  | ---------- | -------- | ------------------------- | ------- |
  | `image`    | `string` | 图片base64                | `true`  |
  | `mime`     | `string` | only image/jpg, image/png | `true`  |
  | `filename` | `string` | 文件名                    | `true`  |

#### 发布文章

```http
POST /api/system/article/publishArticle
```

- 请求
  
  | Field       | Type     | Description | Require |
  | ----------- | -------- | ----------- | ------- |
  | `articleId` | `number` | 文章ID      | `true`  |

- 无响应

#### 编辑文章

```http
POST /api/system/article/updateArticle
```

- 请求

  | Field       | Type      | Description | Require |
  | ----------- | --------- | ----------- | ------- |
  | `articleId` | `number`  | 文章ID      | `true`  |
  | `article`   | `article` | 文章详情    | `true`  |

- 无响应

#### 删除文章

```http
POST /api/system/article/deleteArticle
```

- 请求
  
  | Field       | Type     | Description | Require |
  | ----------- | -------- | ----------- | ------- |
  | `articleId` | `number` | 文章ID      | `true`  |

- 无响应

#### 恢复文章

```http
POST /api/system/article/recoverArticle
```

- 请求
  
  | Field       | Type     | Description | Require |
  | ----------- | -------- | ----------- | ------- |
  | `articleId` | `number` | 文章ID      | `true`  |

- 无响应

#### 获取标签列表

```http
POST /api/system/article/listArticleTags
```

- 请求
  
  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应

  | Field        | Type     | Description |
  | ------------ | -------- | ----------- |
  | `tags`       | `tag[]`  | 标签列表    |
  | `totalCount` | `number` | 文章总数    |

- tag
  
  | Field               | Type     | Description      |
  | ------------------- | -------- | ---------------- |
  | `tagId`             | `number` | 标签ID           |
  | `name`              | `number` | 标签名称         |
  | `articleTotalCount` | `number` | 标签下的文章数量 |

#### 获取友链列表

```http
POST /api/system/friend/listFriends
```

- 请求
  
  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应
  
  | Field        | Type       | Description |
  | ------------ | ---------- | ----------- |
  | `friends`    | `friend[]` | 友链列表    |
  | `totalCount` | `number`   | 友链总数    |

- friend
  
  | Field       | Type     | Description  |
  | ----------- | -------- | ------------ |
  | `friendId`  | `number` | 友链ID       |
  | `name`      | `string` | 友链博客名称 |
  | `link`      | `string` | 友链url      |
  | `avatarUrl` | `string` | 友链图标url  |
  | `desc`      | `string` | 友链简述     |

#### 删除友链

```http
POST /api/system/friend/deleteFriend
```

- 请求

  | Field      | Type     | Description | Require |
  | ---------- | -------- | ----------- | ------- |
  | `friendId` | `number` | 友链ID      | `true`  |

- 无响应

## 博客主页

### 获取文章列表

```http
POST /api/blog/article/listArticles
```

- 请求

  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应

  | Field        | Type        | Description |
  | ------------ | ----------- | ----------- |
  | `articles`   | `article[]` | 文章列表    |
  | `totalCount` | `number`    | 文章总数    |

- article

  | Field         | Type           | Description  |
  | ------------- | -------------- | ------------ |
  | `articleId`   | `number`       | 文章ID       |
  | `title`       | `string`       | 文章标题     |
  | `desc`        | `string`       | 文章描述     |
  | `coverImage`  | `coverImage{}` | 文章头图     |
  | `tags`        | `string[]`     | 标签列表     |
  | `createdTime` | `number`       | 文章创建时间 |

- coverImage

  | Field      | Type     | Description          |
  | ---------- | -------- | -------------------- |
  | `url`      | `string` | 头图url              |
  | `mime`     | `string` | image/jpg, image/png |
  | `filename` | `string` | 文件名               |

#### 获取文章详情

```http
POST /api/blog/article/getArticleDetails
```

- 请求

  | Field       | Type     | Description | Require |
  | :---------- | :------- | :---------- | ------- |
  | `articleId` | `number` | 文章ID      | `true`  |

- 响应

  | Field            | Type               | Description |
  | ---------------- | ------------------ | ----------- |
  | `articleDetails` | `articleDetails{}` | 文章详情    |

- articleDetails

  | Field         | Type           | Description          |
  | ------------- | -------------- | -------------------- |
  | `articleId`   | `number`       | 文章ID               |
  | `title`       | `string`       | 文章标题             |
  | `desc`        | `string`       | 文章描述             |
  | `content`     | `string`       | 文章内容（markdown） |
  | `coverImage`  | `coverImage{}` | 文章头图             |
  | `tags`        | `string[]`     | 标签列表             |
  | `createdTime` | `number`       | 文章创建时间         |
  | `updatedTime` | `number`       | 文章最后更新时间     |

- coverImage

  | Field      | Type     | Description          |
  | ---------- | -------- | -------------------- |
  | `url`      | `string` | 头图url              |
  | `mime`     | `string` | image/jpg, image/png |
  | `filename` | `string` | 文件名               |

#### 搜索文章

```http
POST /api/blog/article/searchArticle
```

- 请求
  
  | Field     | Type     | Description                      | Require |
  | --------- | -------- | -------------------------------- | ------- |
  | `keyword` | `string` | 搜索关键字（接受标题和文章内容） | `true`  |

- 响应

  | Field      | Type        | Description |
  | ---------- | ----------- | ----------- |
  | `articles` | `article[]` | 文章列表    |

- article

  | Field         | Type           | Description  |
  | ------------- | -------------- | ------------ |
  | `articleId`   | `number`       | 文章ID       |
  | `title`       | `string`       | 文章标题     |
  | `desc`        | `string`       | 文章描述     |
  | `coverImage`  | `coverImage{}` | 文章头图     |
  | `tags`        | `string[]`     | 标签列表     |
  | `createdTime` | `number`       | 文章创建时间 |

#### 获取标签列表

```http
POST /api/blog/article/listArticleTags
```

- 请求
  
  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应

  | Field        | Type     | Description |
  | ------------ | -------- | ----------- |
  | `tags`       | `tag[]`  | 标签列表    |
  | `totalCount` | `number` | 文章总数    |

- tag
  
  | Field               | Type     | Description      |
  | ------------------- | -------- | ---------------- |
  | `tagId`             | `number` | 标签ID           |
  | `name`              | `number` | 标签名称         |
  | `articleTotalCount` | `number` | 标签下的文章数量 |

#### 获取友链列表

```http
POST /api/blog/friend/listFriends
```

- 请求
  
  | Field    | Type     | Description            | Require |
  | :------- | :------- | :--------------------- | ------- |
  | `offset` | `number` | 分页偏移量             | `true`  |
  | `size`   | `number` | 期望分页查询记录的数量 | `true`  |

- 响应
  
  | Field        | Type       | Description |
  | ------------ | ---------- | ----------- |
  | `friends`    | `friend[]` | 友链列表    |
  | `totalCount` | `number`   | 友链总数    |

- friend
  
  | Field       | Type     | Description  |
  | ----------- | -------- | ------------ |
  | `friendId`  | `number` | 友链ID       |
  | `name`      | `string` | 友链博客名称 |
  | `link`      | `string` | 友链url      |
  | `avatarUrl` | `string` | 友链图标url  |
  | `desc`      | `string` | 友链简述     |
