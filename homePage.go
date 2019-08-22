package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "storage"
    "io/ioutil"
)

// type Storage struct {
//     Key string `json:"bob"`
//     Value string `json:"ross"`
// }

var Storages []storage.Storage
func homePage(w http.ResponseWriter, r *http.Request){

    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllStorage(w http.ResponseWriter, r *http.Request) {
    // storages := make([]Storage, 2)

    fmt.Println("Endpoint Hit: stored")
    json.NewEncoder(w).Encode(Storages)

}

func returnOneKeyValueStore(w http.ResponseWriter, r *http.Request) {
    // variable := mux.Vars(r)
    for k, v := range r.URL.Query() {
        fmt.Printf("%s: %s\n", k, v)
    }
    fmt.Println("Endpoint Hit: get one key")

    reqBody, _ := ioutil.ReadAll(r.Body)
    var storeKeyValue storage.Storage
    json.Unmarshal(reqBody, &storeKeyValue)
    fmt.Println(storeKeyValue)
    key := storeKeyValue.Key
    // fmt.Fprintf(w, "Key: " + key)
    for _, storage := range Storages {
        if storage.Key == key {
            fmt.Println("Found KEY: ", key)

            json.NewEncoder(w).Encode(storage)
        }
    }
    fmt.Fprintf(w, "Key not found", key)
    fmt.Println("Key does not exist: ", key)
}

func createKeyValueStore(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: create")

    reqBody, _ := ioutil.ReadAll(r.Body)
    fmt.Fprintf(w, "%+v", string(reqBody))
}

func createNewKeyValueStore(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: create")

    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.
    reqBody, _ := ioutil.ReadAll(r.Body)
    var storeKeyValue storage.Storage
    json.Unmarshal(reqBody, &storeKeyValue)
    fmt.Println(storeKeyValue)
    // update our global Articles array to include
    // our new Article
    exist := doesKeyExist(Storages, storeKeyValue.Key)
    if exist {
        fmt.Fprintf(w, "That key already exists: %v", storeKeyValue.Key)
    } else {
        Storages = append(Storages, storeKeyValue)
    }

    json.NewEncoder(w).Encode(storeKeyValue)
}

func doesKeyExist(storage []storage.Storage, key string) bool {
    for _, store := range storage {
        if store.Key == key {
            fmt.Println("ERROR: Key already exists", key)
            // fmt.Fprintf(w, "That Key exists already")
            return true
        }
    }
    return false
}

func handleRequestsToStorage(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("This is the request :")
    fmt.Println(r)
    fmt.Printf("This is the method of the request : ")
    fmt.Println(r.Method)

    reqBody, _ := ioutil.ReadAll(r.Body)
    var storeKeyValue storage.Storage
    json.Unmarshal(reqBody, &storeKeyValue)
    fmt.Printf("This is the decoded storage struct of the request : ")
    fmt.Println(storeKeyValue)
    fmt.Printf("This is the request body of the request : ")
    fmt.Println(string(reqBody))

    switch r.Method {
    case "GET":
        // returnOneKeyValueStore(w, r)
        fmt.Println("Endpoint Hit: get one key")
        key := storeKeyValue.Key
        // fmt.Fprintf(w, "Key: " + key)
        for _, storage := range Storages {
            if storage.Key == key {
                fmt.Println("Found KEY: ", key)

                json.NewEncoder(w).Encode(storage)
                break
            }
            fmt.Println("nothing so far....")
        }
        // fmt.Fprintf(w, "Key not found", key)
        // fmt.Println("Key does not exist: ", key)
    case "POST":
        exist := doesKeyExist(Storages, storeKeyValue.Key)
        if exist {
            fmt.Fprintf(w, "That key already exists: ", storeKeyValue.Key)
        } else {
            Storages = append(Storages, storeKeyValue)
        }

        json.NewEncoder(w).Encode(storeKeyValue)
    case "DELETE":
        key := storeKeyValue.Key
        for index, storage := range Storages {
            if storage.Key == key {
                fmt.Println("Found KEY: ", key)
                fmt.Println("Deleting KEY: ", key)
                Storages = append(Storages[:index], Storages[index+1:]...)
                json.NewEncoder(w).Encode(Storages)

            }
            fmt.Println("nothing so far....")
        }
    case "PUT":
        key := storeKeyValue.Key
        for index, storage := range Storages {
            if storage.Key == key {
                fmt.Println("Found KEY: ", key)
                fmt.Println("Updating KEY: ", key)
                Storages[index] = storeKeyValue
                json.NewEncoder(w).Encode(Storages[index])
            }
            fmt.Println("nothing so far....")
        }
    }
}

func handleRequests() {
    //create a mew instance of a mux router
    // myRouter := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/", homePage)
    http.HandleFunc("/storages", returnAllStorage)
    http.HandleFunc("/storage", handleRequestsToStorage)

    // http.HandleFunc("/storage/", returnOneKeyValueStore)

    log.Fatal(http.ListenAndServe(":10000", nil))

}

func main() {
    fmt.Println("Rest API v2.0 ")
    Storages = []storage.Storage{
        storage.Storage{Key: "a1", Value: "Foo"},
        storage.Storage{Key: "a2", Value: "Boo"},
    }

    handleRequests()
}
