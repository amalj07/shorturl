package shorturl

import (
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"
)

func CreateShortURL(params ...interface{}) (string, error)  {

	fmt.Println(params[0])

	//Check if the Long URL is valid
	_, err := url.ParseRequestURI(params[0].(string))
	if err != nil {
		return "", err
	}

	//Check if the Base URL is valid
	_, err = url.ParseRequestURI(params[1].(string))
	if err != nil {
		return "", err
	}

	rand.Seed(time.Now().UTC().UnixNano())

	original := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

	basestring := []rune(original)

	min1 := 10
	max1 := 26

	min2 := 27
	max2 := 54

	r1 := rand.Intn(max1 - min1) + min1
	r2 := rand.Intn(max2 - min2) + min2
	r3 := rand.Intn(max2 - min2) + min2
	s := strconv.Itoa(time.Now().Second())[:1]
	n := strconv.Itoa(time.Now().Nanosecond())[2:4]
	p := strconv.Itoa(os.Getpid())[:2]

	unum := n + p + s

	shortid := ""
	shortid = shortid  + string(basestring[r1]) + string(basestring[r2])

	for _, v := range unum  {
		u, _ := strconv.Atoi(string(v))
		shortid = shortid + string(basestring[u])
	}

	shortid = shortid + string(basestring[r3])

	shorturl := params[1].(string) + "/" + shortid

	return shorturl, nil
}
