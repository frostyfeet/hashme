package main

import (
"fmt"
"encoding/base64"
"net/http"
"crypto/md5"
"crypto/sha1"
"crypto/sha256"
"golang.org/x/crypto/md4"
"crypto/sha512"
"encoding/hex"
"encoding/json"
"net/url"
)

type Hashlist struct { 
  Hex string
  URL string
  Base64 string
  MD4   string
  MD5	string
  SHA1	string
  SHA224 string
  SHA256 string
  SHA384 string
  SHA512 string
}

func Handler(w http.ResponseWriter, r *http.Request) {
    userString := ""
    if len(r.URL.Path[1:]) < 2000 {
        userString = r.URL.Path[1:]
    }
    temp := Hashlist {hash("hex",userString),hash("url",userString), hash("base64",userString), hash("md4",userString),hash("md5",userString),hash("sha1",userString), hash("sha224",userString),hash("sha256",userString), hash("sha384",userString),hash("sha512",userString)}

    data, _ := json.MarshalIndent(temp,"","    ")
	fmt.Fprintf(w, "%s\n", data)
}

func hash(t string, h string) (string) {
    switch t {
    case "url":
        return url.QueryEscape (h)
    case "md5":
        hasher := md5.New()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "sha1":
        hasher := sha1.New()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "sha224":
		hasher := sha256.New224()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
	case "sha256":
		hasher := sha256.New()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "sha384":
		hasher := sha512.New384()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "sha512":
		hasher := sha512.New()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "md4":
		hasher := md4.New()
        hasher.Write([]byte(h))
        return hex.EncodeToString(hasher.Sum(nil))
    case "base64":
        return base64.StdEncoding.EncodeToString([]byte(h))
    case "hex":
        src := []byte(h)
        return hex.EncodeToString(src)
    default: 
        return "" 
    }
}

func main() {

	fmt.Println("Server Starting")
	http.HandleFunc("/", Handler)

	http.ListenAndServe("127.0.0.1:3000", nil)
}
