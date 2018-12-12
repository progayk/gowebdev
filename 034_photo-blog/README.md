#Content Collision

To be able keep only one copy of a data we use sha algorithms.
Put the content into the algorithm and it will return a hashed value
of the passed data which will be a lot smaller as size. We can use this 
value to name the files so that without opening it we can know its content.

```go
h := sha1.New()
io.Copy(h, "selam")
fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + "jpg"
```

