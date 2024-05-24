package data

import "testing"

func TestChecksvalidation(t *testing.T) {
	w := &Website{
        UserID: 123123,
        Name: "Meta",
        URL: "www.meta.com/blogs/engineering",
        ScriptLink: "https://utfs.io/f/fsdfsdfsdfsd-sdf-sdf-sd-----f-34r-fsdfsdf",
    }
	err := w.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
