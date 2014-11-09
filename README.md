# GoSlash

GoSlash is a redirection service accessible via http://go/ that can be use in many ways to increase productivity or ease the day to day web browsing.

For example :

| alias |  details |
|---|---|
| go/wiki | redirect to internal wiki |
| go/wikipedia/Daft Punk | Daft Punk page on wikipedia |
| go/maps  | go to osm.org |
| go/maps/brest | Search Brest city on OSM  |
| go/rfc | go to tools.ietf.org |
| go/rfc/1149 | go to the RFC1149 |
| go/simple | go to http://domain.tld/complex/path/to/simple |

## How it works

To work, you need that http://go/ can be resolved. Some ways to do it : 
  - without proxy : put the IP of the go service in a /etc/hosts
  - with proxy : have a `go.<domain.tld>` that resolves and set `append_domain` on `<domain.tld>` on your proxy

When you request a `go/alias`, the go service performs a redirection based on rules defined for the given alias. Once the rules have been interpreted, the go service sends the redirection through a HTTP 302.

## Alias definition 

Each alias can redirect to a simple URL but it can do more, catching args : 

    go/<alias>/<arg1>/<arg2>/...

These args can be used to compose the target like passing an argument as a query string.

The rules for using parameters are the followings :

```
  ${X:-DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise substitute parameter X
  ${X:!DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise use the empty string ''
  ${X:+DEFAULT}  If parameter X is set, use 'DEFAULT' where $X was replaced by value of X
  ${X} is equivalent to ${X:-}
```

## Alias Storage

Depending on your usage, the goal is to provide different scaling solutions, from file to database. Stores are an interface implementing methods like Get, Insert, Update, Delete.

### CSV Storage

Format : 
    alias,target,creation,update,user,blob

Examples : 

    ddg,https://duckduckgo.com/${1:+?q=$1},s,1392259631,1392259634,blob
    google,http://www.google.fr/${1:+#q=$1},s,1392259631,1392259634,blob
    rfc,http://tools.ietf.org/html/${1:+rfc}${1:+$1},s,1392259631,1392259634,blob
    rfc2,http://tools.ietf.org/html/${1:+rfc$1},s,1392259631,1392259634,blob
    maps,http://osm.org/,s,1392259631,1392259634,blob

## Implementations

Feel free to pull request other implementations.

### Golang

Status : WIP

Todo list : 

  - admin/auth (coming soon)
  - database stores (couchdb, mongodb, redis, couchbase, sql, etc... )
  - admin web UI
  - add metrics reporter

### Python

Status : prototype

This version is just a prototype
