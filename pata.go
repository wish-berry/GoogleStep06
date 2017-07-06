package main

import (
    "net/http"
    "html/template"
)

type PataString struct{
    A      string
    B      string
    AB     string
}

func patatokukasi(a string, b string)(string){
    if len(a) == 0 || len(b) == 0 {
        return ""
    }
    
    
    
    runeA := []rune(a)
    runeB := []rune(b)
    
    lenA := len(runeA)
    lenB := len(runeB)
    
    loopCount := 0
    if lenA < lenB {
        loopCount = lenA
    } else {
        loopCount = lenB
    }
    
    var unionAB string
    unionAB = ""
    for i := 0; i < loopCount; i++ {
        unionAB = unionAB + string(runeA[i]) + string(runeB[i])
    }

    // 余っている文字列を付加する
    if lenA > lenB {
        for i := loopCount; i < lenA; i++ {
            unionAB = unionAB + string(runeA[i])
        }
    } else if lenA < lenB {
        for i := loopCount; i < lenB; i++ {
            unionAB = unionAB + string(runeB[i])
        }
    }
    
    return unionAB
}

func pataHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()    // formの解析
    a := r.FormValue("a")
    b := r.FormValue("b")
    ab := patatokukasi(a, b)
    p := &PataString{A: a, B: b, AB: ab}
    t, _ := template.ParseFiles("pata.html")
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/", pataHandler)
    http.ListenAndServe(":8080", nil)
}

