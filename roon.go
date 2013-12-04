package main

import (
    "github.com/hoisie/mustache"
    "fmt"
    "errors"
    "encoding/json"
    "net/http"
    "strings"
    "bytes"
)

type Blog struct {
    ID int
    Name string
    Language string
    URL string
    AccentColor string `json:"accent_color"`
}

type Author struct {
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Bio string `json:"bio_html"`
    Website string
    Twitter string
    AvatarURL string
}

func (a Author) Name() string {
    return fmt.Sprintf("%s %s", a.GivenName, a.FamilyName)
}

type Post struct {
    Title string
    Content string `json:"content_html"`
    Excerpt string `json:"excerpt_html"`
    URL string
    ID int
    Blog Blog
    Author Author `json:"user"`
}

func (p Post) Mustache() map[string]string {
    return map[string]string {
        "post_title": p.Title,
        "post_excerpt_html": p.SanitizedExcerpt(),
        "post_url": p.URL,
        "post_content": p.Content,
        "author_name": p.Author.Name(),
        "author_given_name": p.Author.GivenName,
        "author_family_name": p.Author.FamilyName,
        "author_bio_html": p.Author.Bio,
        "author_website": p.Author.Website,
        "author_twitter": p.Author.Twitter,
    }
}

func (p Post) SanitizedExcerpt() string {
    output := ""
    s := p.Excerpt

    // Shortcut strings with no tags in them
    if !strings.ContainsAny(s, "<>") {
        output = s
    } else {
        // First remove line breaks etc as these have no meaning outside html tags (except pre)
        // this means pre sections will lose formatting... but will result in less uninentional paras.
        s = strings.Replace(s, "\n", "", -1)

        // Then replace line breaks with newlines, to preserve that formatting
        s = strings.Replace(s, "</p>", "\n", -1)
        s = strings.Replace(s, "<br>", "\n", -1)
        s = strings.Replace(s, "</br>", "\n", -1)
        s = strings.Replace(s, "\n", " ", -1)

        // Walk through the string removing all tags
        b := bytes.NewBufferString("")
        inTag := false
        for _, r := range s {
            switch r {
            case '<':
                inTag = true
            case '>':
                inTag = false
            default:
                if !inTag {
                    b.WriteRune(r)
                }
            }
        }
        output = b.String()
    }

    return output
}

func main() {
    post, _ := Get("sam", "onward")
    output := mustache.RenderFile("templates/post.html.mustache", post.Mustache())
    fmt.Println(output)
}

func Get(blogSubdomain string, postSlug string) (*Post, error) {
    url := fmt.Sprintf("https://roon.io/api/v1/blogs/%s/posts/%s", blogSubdomain, postSlug)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(resp.Status)
    }
    r := new(Post)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return nil, err
    }
    return r, nil
}
