# Understanding Struct Tags

If you add **_struct tags_**, in this case to json format, you can define custom titles for the JSON object.

```go
type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

This will result in the json format:

```json
{
	"title": "Lorem ipsum",
	"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In neque lorem, tristique eu urna id, suscipit vehicula augue. Sed eleifend aliquet scelerisque. Aenean id ex sed neque porttitor ornare. In et lobortis sem. Phasellus faucibus fermentum risus id efficitur. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce congue imperdiet ipsum, vel semper eros volutpat eget. In vitae velit sed mi mollis pulvinar. Quisque magna diam, auctor sit amet tristique ac, tristique vitae ligula. Nunc pharetra scelerisque auctor.",
	"created_at": "2024-04-15T21:12:16.8484352-06:00"
}
```

If you don't apply a **_struct tag_**, it will result in the json fields applying the same way as typed in the Struct itself.

---
